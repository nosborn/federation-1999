/******************************************************************************

  Copyright (c) 1995-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: trimString.c,v 1.2 1999/04/06 13:10:04 nick Exp $

******************************************************************************/

#include <assert.h>
#include <ctype.h>
#include <string.h>

#include "goodies.h"

char *
IB_trimString(char *s) {
  assert(s != NULL);

  while (!isgraph(*s)) {
    s++;
  }

  if (s[0] != '\0') {
    char *end = s + strlen(s);

    while (!isgraph(*--end)) {
      *end = '\0';
    }
  }

  return s;
}
