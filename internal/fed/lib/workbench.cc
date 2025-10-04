/******************************************************************************

  Copyright (c) 1987-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: workbench.cc,v 1.2 1998/12/24 09:59:15 nick Exp $

******************************************************************************/

#include <sys/stat.h>
#include <errno.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <ibgames/goodies.h>
#include <ibgames.h>

#include <dblocation.hh>
#include <fed.hh>

#include "workbench.h"

static const char *workbenchBasename(const account_id_t);

/*-----------------------------------------------------------------------------

  c r e a t e W o r k b e n c h F i l e s

-----------------------------------------------------------------------------*/

bool
createWorkbenchFiles(account_id_t uid, unsigned which, const char *pszMini) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);

  char pathname[PATH_MAX];

  if ((which & WB_CREATE_EVT) == WB_CREATE_EVT) {
    getEventPathname(uid, pathname, sizeof(pathname));
    dbgCheck(dbgIsValidString(pathname));

    const int fd = creat(pathname, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP);

    if (fd < 0) {
      return false;
    }

    close(fd);
  }

  if ((which & WB_CREATE_LOC) == WB_CREATE_LOC) {
    size_t cbInit;
    dblocation_t *plocInit;

    if (pszMini == NULL) {
      cbInit = sizeof(dblocation_t) * 8;
      plocInit = static_cast<dblocation_t *>(alloca(cbInit));

      if (plocInit == NULL) {
        return false;
      }

      memset(plocInit, '\0', cbInit);
    } else {
      sprintf(pathname, "%s/%s.l", DATADIR, pszMini);
      dbgCheck(dbgIsValidString(pathname));

      const int fd = open(pathname, O_RDONLY);

      if (fd < 0) {
        return false;
      }

      cbInit = IB_fileSize(fd);

      if (cbInit % sizeof(dblocation_t) != 0) {
        close(fd);
        return false;
      }

      plocInit = static_cast<dblocation_t *>(alloca(cbInit));

      if (plocInit == NULL) {
        close(fd);
        return false;
      }

      if (read(fd, plocInit, cbInit) != static_cast<ssize_t>(cbInit)) {
        close(fd);
        return false;
      }

      close(fd);
    }

    getLocationPathname(uid, pathname, sizeof(pathname));
    dbgCheck(dbgIsValidString(pathname));

    const int fd = creat(pathname, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP);

    if (fd < 0) {
      return false;
    }

    if (write(fd, plocInit, cbInit) != static_cast<ssize_t>(cbInit)) {
      close(fd);
      return false;
    }

    close(fd);
  }

  if ((which & WB_CREATE_OBJ) == WB_CREATE_OBJ) {
    getObjectPathname(uid, pathname, sizeof(pathname));
    dbgCheck(dbgIsValidString(pathname));

    const int fd = creat(pathname, S_IRUSR | S_IWUSR | S_IRGRP | S_IWGRP);

    if (fd < 0) {
      return false;
    }

    close(fd);
  }

  return true;
}

/*-----------------------------------------------------------------------------

  d e l e t e W o r k b e n c h F i l e s

-----------------------------------------------------------------------------*/

void
deleteWorkbenchFiles(account_id_t uid) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);

  char pathname[PATH_MAX];

  getEventPathname(uid, pathname, sizeof(pathname));
  dbgCheck(dbgIsValidString(pathname));
  unlink(pathname);

  getLocationPathname(uid, pathname, sizeof(pathname));
  dbgCheck(dbgIsValidString(pathname));
  unlink(pathname);

  getObjectPathname(uid, pathname, sizeof(pathname));
  dbgCheck(dbgIsValidString(pathname));
  unlink(pathname);
}

/*-----------------------------------------------------------------------------

  g e t E v e n t P a t h n a m e

-----------------------------------------------------------------------------*/

void
getEventPathname(account_id_t uid, char *pathname, size_t cbPathname) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);
  dbgPrecondition(pathname != NULL);
  dbgPrecondition(cbPathname > 0);

  sprintf(pathname, "%s.e", workbenchBasename(uid));
  dbgCheck(memchr(pathname, '\0', cbPathname) != NULL);
}

/*-----------------------------------------------------------------------------

  g e t L o c a t i o n P a t h n a m e

-----------------------------------------------------------------------------*/

void
getLocationPathname(account_id_t uid, char *pathname, size_t cbPathname) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);
  dbgPrecondition(pathname != NULL);
  dbgPrecondition(cbPathname > 0);

  sprintf(pathname, "%s.l", workbenchBasename(uid));
  dbgCheck(memchr(pathname, '\0', cbPathname) != NULL);
}

/*-----------------------------------------------------------------------------

  g e t O b j e c t P a t h n a m e

-----------------------------------------------------------------------------*/

void
getObjectPathname(account_id_t uid, char *pathname, size_t cbPathname) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);
  dbgPrecondition(pathname != NULL);
  dbgPrecondition(cbPathname > 0);

  sprintf(pathname, "%s.o", workbenchBasename(uid));
  dbgCheck(memchr(pathname, '\0', cbPathname) != NULL);
}

/*-----------------------------------------------------------------------------

  w o r k b e n c h A c c e s s

-----------------------------------------------------------------------------*/

int
workbenchAccess(account_id_t uid) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);

  char eventPathname[PATH_MAX];
  char locationPathname[PATH_MAX];
  char objectPathname[PATH_MAX];

  getEventPathname(uid, eventPathname, sizeof(eventPathname));
  getLocationPathname(uid, locationPathname, sizeof(locationPathname));
  getObjectPathname(uid, objectPathname, sizeof(objectPathname));

  // See if the files exist.

  if (access(eventPathname, F_OK) != 0 ||
      access(locationPathname, F_OK) != 0 ||
      access(objectPathname, F_OK) != 0) {
    return WB_NO_FILES;
  }

  // OK, now see if we'll be able to write to them.

  if (access(eventPathname, W_OK) != 0 ||
      access(locationPathname, W_OK) != 0 ||
      access(objectPathname, W_OK) != 0) {
    return WB_CANT_WRITE;
  }

  // All OK.

  return WB_ACCESS_OK;
}

/*-----------------------------------------------------------------------------

  w o r k b e n c h B a s e n a m e

-----------------------------------------------------------------------------*/

static const char *
workbenchBasename(account_id_t uid) {
  dbgPrecondition(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);

  static char pathname[PATH_MAX];

  sprintf(pathname, "%s/data/workbench%d/%u",
          IB_homeDir(),
          static_cast<int>(uid % 10L),
          uid);

  return pathname;
}
