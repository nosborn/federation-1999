/******************************************************************************

  Copyright (c) 1997-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: perivale.cc,v 1.8 1999/04/24 21:03:39 nick Exp $

******************************************************************************/

#include <algorithm>
#include <string>

#include <arpa/inet.h>
#include <netinet/in.h>
#include <sys/ioctl.h>
#include <sys/socket.h>
#include <sys/time.h>
#include <sys/un.h>
#include <netdb.h>
#include <assert.h>
#include <ctype.h>
#include <errno.h>
#include <signal.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <syslog.h>
#include <termios.h>
#include <unistd.h>

#include <ibgames/goodies.h>
#include <ibgames/rules.h>
#include <ibgames.h>

#include <link.hh>

static void error(const char *);
static void fatal(const char *) __attribute__((noreturn));
static void flushAllOutput();
static void flushOutput();
static void lockout() __attribute__((noreturn));
static void output(const unsigned char *, size_t);
static void readInput(int);
static void sigWinch(int);
static void unavailable() __attribute__((noreturn));
static void warning(const char *);
static void writeOutput(const char *, size_t);

typedef enum {
  inReady,
  inAckWait,
  inThrottle
} in_state_t;

static sig_atomic_t winch = true; // Fake it the first time through

static unsigned rows = 0;
static unsigned cols = 0;

static in_state_t inState = inReady;
static struct timeval unthrottle = { 0, 0 };

typedef struct iblock {
  struct iblock *next;
  struct iblock *prev;
  size_t len;
  char data[0];
} iblock_t;

typedef struct oblock {
  struct oblock *next;
  struct oblock *prev;
  size_t sent;
  size_t len;
  char data[0];
} oblock_t;

static iblock_t iqHead = { NULL, NULL, 0, {} };
static iblock_t iqTail = { NULL, NULL, 0, {} };
static oblock_t oqHead = { NULL, NULL, 0, 0, {} };
static oblock_t oqTail = { NULL, NULL, 0, 0, {} };

int
main(int argc, char *argv[]) {
  fflush(stdin);
  fflush(stdout);

  if (argc != 4) {
    fprintf(stderr, "usage: %s id address robobod\n", argv[0]);
    exit(EXIT_FAILURE);
  }

  //

  iqHead.next = &iqTail;
  iqTail.prev = &iqHead;
  oqHead.next = &oqTail;
  oqTail.prev = &oqHead;

  //

  const char *term = getenv("TERM");

  if (term != NULL) {
    ;
  }

  // Check for a valid UID.

  const account_id_t uid = atoi(argv[1]);

  if (uid < MIN_ACCOUNT_ID || uid >= MAX_ACCOUNT_ID) {
    fprintf(stderr, "%s: Bad UID\n", argv[0]);
    exit(EXIT_FAILURE);
  }

  // Check for lock out!

  if (isLockedOut(uid)) {
    lockout();
  }

  // Set up the login string in advance; the TP task is going to block between
  // our connect and send, so keep the window as small as possible.

  struct in_addr remote_addr;
  remote_addr.s_addr = inet_addr(argv[2]);

  const struct hostent *remote_host =
    gethostbyaddr(reinterpret_cast<char *>(&remote_addr), sizeof(remote_addr), AF_INET);

  const char *remote_hostname;

  if (remote_host == NULL) {
    remote_hostname = argv[2]; // Use the dotted quad instead
  } else {
    remote_hostname = remote_host->h_name;

    while (strlen(remote_hostname) > LINK_HOSTNAME_SIZE && *remote_hostname != '.') {
      remote_hostname++;
    }
  }

  char login[1024];
  const int len = snprintf(login, sizeof(login), LINK_FORMAT, uid, argv[2], remote_hostname, atoi(argv[3]));
  assert(len <= static_cast<int>(sizeof(login)));

  //

  struct sockaddr_un addr;

  memset(&addr, '\0', sizeof(addr));
  addr.sun_family = AF_UNIX;
  sprintf(addr.sun_path, LINK_SOCKET, IB_homeDir());

  const int sd = socket(addr.sun_family, SOCK_STREAM, 0);

  if (sd < 0) {
    fatal("socket");
  }

  if (connect(sd, reinterpret_cast<struct sockaddr *>(&addr), sizeof(addr)) < 0) {
    unavailable();
  }

  if (send(sd, login, len, 0) != len) {
    unavailable();
  }

  //

  if (IB_nonBlocking(STDIN_FILENO) == -1) {
    fatal("IB_nonBlocking");
  }

  if (IB_nonBlocking(STDOUT_FILENO) == -1) {
    fatal("IB_nonBlocking");
  }

  // Let's try this without input echo and see if there are any complaints.

  struct termios termios;

  if (tcgetattr(STDIN_FILENO, &termios) == -1) {
    warning("tcgetattr");
  } else {
    termios.c_iflag |= ICRNL;
    termios.c_iflag |= IGNBRK;
    termios.c_iflag |= ISTRIP;

    termios.c_lflag |= ICANON;
    termios.c_lflag &= ~ECHO;
    termios.c_lflag &= ~ISIG;

    termios.c_cc[VINTR] = _POSIX_VDISABLE;
    termios.c_cc[VQUIT] = _POSIX_VDISABLE;
    termios.c_cc[VKILL] = _POSIX_VDISABLE;

    if (tcsetattr(STDIN_FILENO, TCSANOW, &termios) == -1) {
      warning("tcsetattr");
    }
  }

  //

  signal(SIGINT, SIG_IGN);
  signal(SIGPIPE, SIG_IGN);
  signal(SIGQUIT, SIG_IGN);
  signal(SIGTSTP, SIG_IGN);
  signal(SIGWINCH, sigWinch);

  //

  for (;;) {
    fd_set readfds;
    FD_ZERO(&readfds);

    FD_SET(STDIN_FILENO, &readfds);
    FD_SET(sd, &readfds);

    fd_set writefds;
    FD_ZERO(&writefds);

    if (oqHead.next != &oqTail) {
      FD_SET(STDOUT_FILENO, &writefds);
    }

    struct timeval *timeout = NULL;
    struct timeval wakeup;

    if (inState == inThrottle) {
      struct timeval now;
      gettimeofday(&now, NULL);

      if (timercmp(&now, &unthrottle, >)) {
        inState = inReady;
      } else {
        wakeup.tv_sec = unthrottle.tv_sec - now.tv_sec;
        wakeup.tv_usec = unthrottle.tv_usec - now.tv_usec;
        if (wakeup.tv_usec < 0) {
          wakeup.tv_sec -= 1;
          wakeup.tv_usec += 1000000;
        }
        timeout = &wakeup;
      }
    }

    const int nfds = select(sd + 1, &readfds, &writefds, NULL, timeout);

    if (nfds == -1 && errno != EINTR) {
      fatal("select");
    }

    if (winch) {
      struct winsize size;

      if (ioctl(STDIN_FILENO, TIOCGWINSZ, &size) < 0) {
        fatal("TIOCGWINSZ");
      }

      rows = size.ws_row;
      cols = size.ws_col;

      winch = false;
    }

    if (nfds == -1) {
      continue;
    }

    if (nfds == 0 && inState == inThrottle) {
      inState = inReady;
    }

    if (nfds > 0) {
      if (FD_ISSET(sd, &readfds)) {
        unsigned char buf[32768];
        const ssize_t nbytes = read(sd, buf, sizeof(buf));

        if (nbytes <= 0) {
          flushAllOutput();
          if (nbytes == -1) {
            fatal("read");
          }
          sleep(1);
          exit(EXIT_SUCCESS);
        }

        output(buf, nbytes);
      }

      if (FD_ISSET(STDOUT_FILENO, &writefds)) {
        flushOutput();
      }

      if (FD_ISSET(STDIN_FILENO, &readfds)) {
        readInput(sd);
      }
    }

    if (iqHead.next != &iqTail && inState == inReady) {
      iblock_t *b = iqHead.next;

      if (write(sd, b->data, b->len) != static_cast<ssize_t>(b->len)) {
        fatal("write");
      }

      iqHead.next = b->next;
      b->next->prev = &iqHead;
      free(b);

      inState = inAckWait;
    }
  }
}

static void
error(const char *s) {
  const int errnum = errno;

  perror(s);
  syslog(LOG_ERR, "%s: %s", s, strerror(errnum));
}

static void
fatal(const char *s) {
  error(s);
  exit(EXIT_FAILURE);
}

static void
flushAllOutput() {
  while (oqHead.next != &oqTail) {
    flushOutput();
    sleep(1);
  }
}

static void
flushOutput() {
  assert(oqHead.next != &oqTail);

  do {
    oblock_t *b = oqHead.next;
    ssize_t sent = write(STDOUT_FILENO, b->data + b->sent, b->len);

    if (sent == -1) {
      if (errno != EAGAIN) {
        fatal("write");
      }
      sent = 0;
    }

    b->sent += sent;
    b->len -= sent;

    if (b->len > 0) {
      break;
    }

    oqHead.next = b->next;
    b->next->prev = &oqHead;
    free(b);
  } while (oqHead.next != &oqTail);
}

static void
lockout() {
  puts("");
  puts("Your account has been locked out of the Federation game. If you");
  puts("are unsure about why you have been locked out of the game, please");
  puts("read the rules at <URL:http://www.ibgames.net/ibinfo/t&c.html> or");
  printf("send e-mail to %s.\n\n", RULES_INFO_EMAIL);
  fflush(stdout);

  sleep(5);

  exit(EXIT_SUCCESS);
}

// codeql[c-cpp/poorly-documented-function] fedtpd output processing state machine - complexity is inherent
static void
output(const unsigned char *buf, size_t len) {
  typedef enum {
    tsData,
    tsSpyDepth,
    tsTrace,
    tsEscape
  } state_t;

  static state_t state = tsData;
  static bool trace = false;
  static bool blankPending = true;
  static size_t column = 1;
  static bool allWS = false;
  static bool eatBlanks = false;
  static bool lastBlank = false;
  static int lastSpyDepth = 0;
  static std::string prefix = "";
  static int spyDepth = 0;
  static bool continuation = false;

  assert(buf != NULL);

  std::string output;
  std::string line;
  char tbuf[BUFSIZ];
  bool wasTrace;

  if (trace) {
    sprintf(tbuf, "[LEN=%zu]", len);
    output += tbuf;
  }

  size_t i = 0;

  while (i < len) {
    unsigned char ch = buf[i++];

    // codeql[c-cpp/long-switch] fedtpd output state machine requires detailed case handling
    switch (state) {
      case tsData:
        // codeql[c-cpp/long-switch] DLE escape sequence handling
        switch (ch) {
          case DLE:
            state = tsEscape;
            break;

          case '\a':
            line += ch;
            break;

          case '\n':
            if (allWS) {
              if (!lastBlank) {
                lastBlank = true;
              }
              line.erase();
              blankPending = true;
            } else {
              lastBlank = false;
            }

            if (!line.empty() || continuation) {
              size_t rpos = line.find_last_not_of(' ');
              if (rpos != std::string::npos) {
                size_t chop = (line.size() - rpos) - 1;
                if (chop > 0) {
                  line.erase(rpos + 1, chop);
                }
              }
              if (blankPending) {
                if (trace) {
                  output += "[PEND]";
                }
                if (spyDepth == lastSpyDepth) {
                  output += prefix;
                }
                output += '\n';
                blankPending = false;
              }
              if (spyDepth > 0) {
                output += prefix;
                lastSpyDepth = spyDepth;
              }
              output += line;
              if (trace) {
                output += "[PARA]";
              }
              output += '\n';
              line.erase();
              continuation = false;
            }

            column = 1;
            allWS = true;
            eatBlanks = false;

            break;

          case '\r': // There shouldn't be any!
            continue;

          case ' ':
            if (eatBlanks) {
              continue;
            }
            // FALL THROUGH

          default:
            allWS &= (ch == ' ');

            if (eatBlanks) {
              if (ch == '/' || ch == '>') {
                continue;
              }
              eatBlanks = false;
            }

            if (!isascii(ch) || !isprint(ch)) {
              line += '^';
              line += ch ^ 0100;
            } else {
              line += ch;
            }

            if (column++ == cols - prefix.length()) {
              size_t rpos = line.find_last_of(' ');
              if (rpos != std::string::npos) {
                size_t chop = (line.size() - rpos) - 1;
                if (chop > 0) {
                  i -= chop;
                  line.erase(rpos + 1, chop);
                }
              }

              rpos = line.find_last_not_of(' ');

              if (rpos != std::string::npos) {
                size_t chop = (line.size() - rpos) - 1;
                if (chop > 0) {
                  line.erase(rpos + 1, chop);
                }
              }

              if (blankPending) {
                if (trace) {
                  output += "[PEND]";
                }
                if (spyDepth == lastSpyDepth) {
                  output += prefix;
                }
                output += '\n';
                blankPending = false;
              }

              if (spyDepth > 0) {
                output += prefix;
                lastSpyDepth = spyDepth;
              }

              output += line;
              if (trace) {
                output += "[WRAP]";
              }
              output += "\n";
              line.erase();
              continuation = false;

              column = 1;
              allWS = (line.find_first_not_of(' ') == std::string::npos);
              eatBlanks = true;
            }
        }
        break;

      case tsEscape:
        switch (ch) {
          case leAck:
            if (spyDepth == 0) {
              if (trace) {
                output += "[ACK]";
              }
              if (inState == inAckWait) {
                gettimeofday(&unthrottle, NULL);
                if (unthrottle.tv_usec < 500000) {
                  unthrottle.tv_usec += 500000;
                } else {
                  unthrottle.tv_sec += 1;
                  unthrottle.tv_usec -= 500000;
                }
                inState = inThrottle;
              }
            }
            state = tsData;
            break;

          case leSpy:
            state = tsSpyDepth;
            break;

          case leTrace:
            state = tsTrace;
            break;

          default:
            state = tsData;
        }
        break;

      case tsSpyDepth:
        if (trace) {
          sprintf(tbuf, "[S=%u]", ch);
          output += tbuf;
        }

        prefix = "";
        lastSpyDepth = spyDepth;
        spyDepth = ch;

        if (spyDepth > 0) {
          for (int j = 0; j < spyDepth; j++) {
            prefix += '/';
          }
          prefix += ' ';
        }

        state = tsData;
        break;

      case tsTrace:
        if (spyDepth == 0) {
          wasTrace = trace;
          trace = (ch != '-');

          if (trace || wasTrace) {
            sprintf(tbuf, "[T=%d]", trace);
            output += tbuf;
          }
        }
        state = tsData;
        break;
    }
  }

  if (!line.empty()) {
    if (blankPending) {
      if (spyDepth == lastSpyDepth) {
        output += prefix;
      }
      output += '\n';
      blankPending = false;
    }
    if (spyDepth > 0) {
      output += prefix;
      lastSpyDepth = spyDepth;
    }
    output += line;
    continuation = true;
  }

  if (!output.empty()) {
    writeOutput(output.data(), output.size());
  }
}

static void
readInput(int pipe) {
  (void) pipe;

  static unsigned char line[256];
  static size_t len = 0;

  unsigned char buf[BUFSIZ];
  const ssize_t nbytes = read(STDIN_FILENO, buf, sizeof(buf));

  if (nbytes <= 0) {
    if (nbytes == -1) {
      fatal("read");
    }
    exit(EXIT_SUCCESS);
  }

  iblock_t *b;

  for (ssize_t i = 0; i < nbytes; i++) {
    const unsigned char ch = buf[i];

    switch (ch) {
      case 0x00:
      case 0x01: // Ctrl-A
      case 0x02: // Ctrl-B
      case 0x03: // Ctrl-C
      case 0x04: // Ctrl-D
      case 0x05: // Ctrl-E
      case 0x06: // Ctrl-F
      case 0x07: // Ctrl-G
      case 0x08: // Ctrl-H
      case 0x09: // Ctrl-I
        break;

      case 0x0A: // Ctrl-J
        while (len > 0 && line[len - 1] == ' ') {
          len--;
        }
        line[len++] = '\n';

        b = static_cast<iblock_t *>(malloc(sizeof(iblock_t) + len));

        if (b == NULL) {
          fatal("malloc");
        }

        b->next = &iqTail;
        b->prev = iqTail.prev;
        b->len = len;
        memcpy(b->data, line, len);

        b->prev->next = b;
        iqTail.prev = b;

        len = 0;
        break;

      case 0x0B: // Ctrl-K
      case 0x0C: // Ctrl-L
      case 0x0D: // Ctrl-M
      case 0x0E: // Ctrl-N
      case 0x0F: // Ctrl-O
      case 0x10: // Ctrl-P
      case 0x11: // Ctrl-Q
        break;

      case 0x12: // Ctrl-R
        if (len > 0) {
          output(line, len);
        }
        break;

      case 0x13: // Ctrl-S
      case 0x14: // Ctrl-T
      case 0x15: // Ctrl-U
      case 0x16: // Ctrl-V
      case 0x17: // Ctrl-W
      case 0x18: // Ctrl-X
      case 0x19: // Ctrl-Y
      case 0x1A: // Ctrl-Z
      case 0x1B:
      case 0x1C:
      case 0x1D:
      case 0x1E:
      case 0x1F:
        break;

      case 0x7F:
        break;

      default:
        line[len] = ch;

        if (len < sizeof(line)) {
          len++;
        }
    }
  }
}

static void
sigWinch(int) {
  signal(SIGWINCH, sigWinch);
  winch = true;
}

static void
unavailable() {
  static const char message[] =
    "\n\nFederation is temporarily unavailable. Please try again later.\n\n";

  sleep(2);

  fputs(message, stdout);
  fflush(stdout);

  sleep(3);

  exit(EXIT_SUCCESS);
}

static void
warning(const char *s) {
  const int errnum = errno;

  perror(s);
  syslog(LOG_ERR, "%s: %s", s, strerror(errnum));
}

static void
writeOutput(const char *buf, size_t count) {
  if (oqHead.next == &oqTail) {
    ssize_t sent = write(STDOUT_FILENO, buf, count);

    if (sent == -1) {
      if (errno != EAGAIN) {
        fatal("write");
      }
      sent = 0;
    }

    buf += sent;
    count -= sent;
  }

  if (count == 0) {
    return;
  }

  oblock_t *b = static_cast<oblock_t *>(malloc(sizeof(oblock_t) + count));

  if (b == NULL) {
    fatal("malloc");
  }

  b->next = &oqTail;
  b->prev = oqTail.prev;
  b->sent = 0;
  b->len = count;
  memcpy(b->data, buf, count);

  b->prev->next = b;
  oqTail.prev = b;
}
