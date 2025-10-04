/******************************************************************************

  Copyright (c) 1995-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: homeDir.c,v 1.2 1999/04/06 13:10:04 nick Exp $

******************************************************************************/

#include <assert.h>
#include <limits.h>
#include <pwd.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include "goodies.h"

const char *
IB_homeDir(void) {
  static char homeDir[PATH_MAX] = "";

  if (homeDir[0] == '\0') {
    const struct passwd *pwd = getpwuid(geteuid());
    assert(pwd != NULL);
    assert(strlen(pwd->pw_dir) < sizeof(homeDir));
    strcpy(homeDir, pwd->pw_dir);
  }

  assert(memchr(homeDir, '\0', sizeof(homeDir)) != NULL);
  assert(strlen(homeDir) > 0);

  return homeDir;
}
