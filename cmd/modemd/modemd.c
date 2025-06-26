#include <arpa/inet.h>
#include <arpa/telnet.h>
#include <netinet/in.h>
#include <netinet/ip.h>
#include <netinet/tcp.h>
#include <sys/socket.h>
#include <errno.h>
#include <fcntl.h>
#include <pthread.h>
#include <signal.h>
#include <stdatomic.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "modem.h"

extern char **environ;

static atomic_bool urgent_received = false;

static bool read_proxy_line(unsigned char *line, size_t len);
static void set_environment(void);
static void sniff_client_tspeed(unsigned char c);
static void start_telnetd(int *stdinfd, int *stdoutfd);
static void *send_thread(void *arg);
static void *receive_thread(void *arg);
static void urgent_handler(int sig);

#define TS_DATA 0 // base state
#define TS_IAC 1  // look for double IAC's
#define TS_SB 3   // throw away begin's...
#define TS_SE 4   // ...end's (suboption negotiation)

int
main(int argc, char *argv[]) {
  const int on = 1;
  if (setsockopt(STDIN_FILENO, SOL_SOCKET, SO_KEEPALIVE, &on, sizeof(on)) < 0) {
    perror("setsockopt(SO_KEEPALIVE)");
  }
  if (setsockopt(STDIN_FILENO, IPPROTO_TCP, TCP_NODELAY, &on, sizeof(on)) < 0) {
    perror("setsockopt(TCP_NODELAY)");
  }

  const int tos = IPTOS_LOWDELAY;
  if (setsockopt(STDOUT_FILENO, IPPROTO_IP, IP_TOS, &tos, sizeof(tos)) < 0) {
    perror("setsockopt(IP_TOS)");
  }

  modem_init();

  sigset_t mask;
  sigemptyset(&mask);
  sigaddset(&mask, SIGURG);
  if (sigprocmask(SIG_BLOCK, &mask, NULL) < 0) {
    perror("sigprocmask(SIG_BLOCK)");
    exit(EXIT_FAILURE);
  }

  if (getenv("FLY_APP_NAME") != NULL) {
    set_environment();
  }

  int telnetd_stdinfd, telnetd_stdoutfd;
  start_telnetd(&telnetd_stdinfd, &telnetd_stdoutfd);

  pthread_t send_tid, receive_tid;
  if (pthread_create(&send_tid, NULL, send_thread, &telnetd_stdoutfd) != 0) {
    perror("pthread_create(send_thread)");
    exit(EXIT_FAILURE);
  }
  if (pthread_create(&receive_tid, NULL, receive_thread, &telnetd_stdinfd) != 0) {
    perror("pthread_create(receive_thread)");
    exit(EXIT_FAILURE);
  }

  pause();

  exit(EXIT_SUCCESS);
}

static void
set_environment(void) {
  unsigned char line[110]; // worst case (107 chars) + CR/LF + NUL
  if (!read_proxy_line(line, sizeof(line))) {
    fprintf(stderr, "modemd[%d]: failed to read PROXY protocol header\n", getpid());
    exit(EXIT_FAILURE);
  }

  char proto[8];
  char src_ip[INET6_ADDRSTRLEN], dst_ip[INET6_ADDRSTRLEN];
  unsigned short src_port, dst_port;
  if (sscanf((char *) line, "PROXY %7s %s %s %hu %hu\r\n", proto, src_ip, dst_ip, &src_port, &dst_port) != 5) {
    fprintf(stderr, "modemd[%d]: failed to parse PROXY protocol header\n", getpid());
    exit(EXIT_FAILURE);
  }
  if (strcmp(proto, "UNKNOWN") == 0) {
    return;
  }

  char port_buf[6];
  sprintf(port_buf, "%hu", src_port);
  setenv("TCPREMOTEIP", src_ip, 1);
  setenv("TCPREMOTEPORT", port_buf, 1);
  sprintf(port_buf, "%hu", dst_port);
  setenv("TCPLOCALIP", dst_ip, 1);
  setenv("TCPLOCALPORT", port_buf, 1);
}

static bool
read_proxy_line(unsigned char *line, size_t len) {
  int i = 0;

  while (i < len) {
    unsigned char c;

    if (recv(STDIN_FILENO, &c, 1, 0) != 1) {
      return false;
    }
    line[i++] = c;
    if (c == '\r') {
      if (i == len || recv(STDIN_FILENO, &c, 1, 0) != 1 || c != '\n') {
        return false;
      }
      line[i++] = c;
      break;
    }
  }
  line[i] = '\0';
  return true;
}

static void
start_telnetd(int *stdinfd, int *stdoutfd) {
  if (signal(SIGCHLD, SIG_IGN) == SIG_ERR) {
    perror("signal");
    exit(EXIT_FAILURE);
  }

  int sv[2];
  if (socketpair(AF_UNIX, SOCK_STREAM, 0, sv) == -1) {
    perror("socketpair");
    exit(EXIT_FAILURE);
  }

  const pid_t pid = fork();
  if (pid < 0) {
    perror("fork");
    exit(EXIT_FAILURE);
  }

  if (pid == 0) {
    close(sv[1]); // close parent's end

    if (dup2(sv[0], STDIN_FILENO) != STDIN_FILENO) {
      perror("dup2");
      _exit(EXIT_FAILURE);
    }
    if (dup2(sv[0], STDOUT_FILENO) != STDOUT_FILENO) {
      perror("dup2");
      _exit(EXIT_FAILURE);
    }

    for (int fd = STDERR_FILENO + 1; fd < 1024; fd++) {
      if (close(fd) == -1 && errno != EBADF) {
        perror("close");
      }
    }

    char *argv[] = { "/usr/sbin/telnetd", "--authmode=off", "--exec-login=/opt/fed/bin/login", "--no-hostinfo", NULL };
    execve(argv[0], argv, environ);
    _exit(EXIT_FAILURE); // exec never returns
  }

  close(sv[0]); // close child's end

  *stdinfd = sv[1];
  *stdoutfd = sv[1];
}

static void *
send_thread(void *arg) {
  const int telnetd_stdoutfd = *(int *) arg;
  unsigned char buf[BUFSIZ];

  for (;;) {
    ssize_t n = recv(telnetd_stdoutfd, buf, sizeof(buf), 0);
    if (n < 0) {
      perror("send_thread: recv");
      exit(EXIT_FAILURE);
    }
    if (n == 0) {
      exit(EXIT_SUCCESS);
    }

    if (modem_send(STDOUT_FILENO, buf, n, 0) < 0) {
      perror("send_thread: modem_send");
      exit(EXIT_FAILURE);
    }
  }

  return NULL;
}

static void *
receive_thread(void *arg) {
  int telnetd_stdinfd = *(int *) arg;

  struct sigaction sa = {
    .sa_handler = urgent_handler,
    .sa_flags = 0
  };
  sigemptyset(&sa.sa_mask);
  if (sigaction(SIGURG, &sa, NULL) < 0) {
    perror("sigaction(SIGURG)");
    exit(EXIT_FAILURE);
  }

  sigset_t mask;
  sigemptyset(&mask);
  sigaddset(&mask, SIGURG);
  if (pthread_sigmask(SIG_UNBLOCK, &mask, NULL) != 0) {
    perror("pthread_sigmask(SIG_UNBLOCK)");
    exit(EXIT_FAILURE);
  }

  for (;;) {
    struct timeval *timeout = modem_read_timeout();
    if (timeout != NULL) {
      if (usleep(timeout->tv_usec) != 0) {
        if (errno == EINTR) {
          continue;
        }
        perror("receive_thread: usleep");
        exit(EXIT_FAILURE);
      }
    }

    unsigned char c;
    ssize_t n = recv(STDIN_FILENO, &c, 1, 0);
    if (n < 0) {
      if (errno == EINTR) {
        if (atomic_exchange(&urgent_received, false)) {
          unsigned char oob[2];
          ssize_t oob_n = recv(STDIN_FILENO, oob, sizeof(oob), MSG_OOB);
          if (oob_n < 0) {
            perror("receive_thread: recv OOB");
            exit(EXIT_FAILURE);
          }
          if (oob_n > 0) {
            ssize_t sent = 0;
            while (sent < oob_n) {
              ssize_t s = send(telnetd_stdinfd, oob + sent, oob_n - sent, MSG_OOB);
              if (s < 0) {
                perror("receive_thread: send OOB");
                exit(EXIT_FAILURE);
              }
              sent += s;
            }
          }
        }
        continue;
      }
      perror("receive_thread: recv");
      exit(EXIT_FAILURE);
    }
    if (n == 0) {
      exit(EXIT_SUCCESS);
    }

    if (send(telnetd_stdinfd, &c, 1, 0) < 0) {
      perror("receive_thread: send");
      exit(EXIT_FAILURE);
    }

    sniff_client_tspeed(c);
    modem_on_read();
  }

  return NULL;
}

static void
urgent_handler(int sig) {
  atomic_store(&urgent_received, true);
}

static void
sniff_client_tspeed(unsigned char c) {
  static unsigned char buf[512];
  static unsigned char *bufp = buf;
  static unsigned char *endp = buf;
  static int state = TS_DATA;

  switch (state) {
    case TS_DATA:
      if (c == IAC) {
        state = TS_IAC;
      }
      break;
    case TS_IAC:
      if (c == SB) {
        state = TS_SB;
        bufp = buf;
      } else {
        state = TS_DATA;
      }
      break;
    case TS_SB:
      if (c == IAC) {
        state = TS_SE;
      } else {
        if (bufp < (buf + sizeof(buf))) {
          *bufp++ = c;
        }
      }
      break;
    case TS_SE:
      if (c == SE) {
        endp = bufp;
        bufp = buf;

        if (*bufp++ == TELOPT_TSPEED) {
          if (bufp >= endp || *bufp++ != TELQUAL_IS) {
            break;
          }
          const int xspeed = atoi((char *) bufp);
          while (*bufp++ != ',' && bufp < endp) {
            ;
          }
          if (bufp >= endp) {
            break;
          }
          const int rspeed = atoi((char *) bufp);
          set_modem_speeds(xspeed, rspeed);
        }
        state = TS_DATA;
      }
      break;
  }
}
