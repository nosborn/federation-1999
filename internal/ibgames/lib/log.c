/******************************************************************************

  Copyright (c) 1996-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: log.c,v 1.4 1999/04/23 19:30:29 nick Exp $

******************************************************************************/

#include <errno.h>
#include <fcntl.h>
#include <limits.h>
#include <stdarg.h>
#include <stdbool.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <syslog.h>
#include <time.h>
#include <unistd.h>

#include "goodies.h"
#include "log.h"

static struct {
  char *ident;
  bool addPid;
} options;

static int optionsSet = 0;

static int setOptions(const char *, int);

/*-----------------------------------------------------------------------------

  I B _ l o g

-----------------------------------------------------------------------------*/

int
IB_log(const char *format, ...) {
  int len;
  va_list arg;
  char logLine[PIPE_BUF];

  if (optionsSet == 0) {
    extern const IB_logOptions_t IB_logOptions;
    if (setOptions(IB_logOptions.ident, IB_logOptions.option) == -1) {
      return -1;
    }
  }

  if (options.ident == NULL) {
    len = 0;
  } else {
    if (options.addPid) {
      len = sprintf(logLine, "%s[%d] ", options.ident, getpid());
    } else {
      len = sprintf(logLine, "%s ", options.ident);
    }
  }

  va_start(arg, format);
  len += vsprintf(logLine + len, format, arg);
  logLine[len++] = '\n';
  va_end(arg);

  for (;;) {
    if (write(STDERR_FILENO, logLine, len) == len) {
      return 0;
    }
    /* Hmmm... what about partial writes? */
    if (errno != EPIPE) {
      const int saved_errno = errno;
      syslog(LOG_ERR, "IB_log: write: %m");
      errno = saved_errno;
      return -1;
    }
    sleep(1);
  }
}

/*-----------------------------------------------------------------------------

  I B _ s e t L o g O p t i o n s

-----------------------------------------------------------------------------*/

int
IB_setLogOptions(const char *ident, int option) {
  return setOptions(ident, option);
}

/*-----------------------------------------------------------------------------

  s e t O p t i o n s

-----------------------------------------------------------------------------*/

static int
setOptions(const char *ident, int option) {
  if (options.ident != NULL) {
    free(options.ident);
    options.ident = NULL;
  }
  if (ident != NULL && strlen(ident) > 0) {
    options.ident = strdup(ident);
    if (options.ident == NULL) {
      errno = ENOMEM;
      return -1;
    }
  }
  options.addPid = ((option & LOG_PID) == LOG_PID);
  return 0;
}
