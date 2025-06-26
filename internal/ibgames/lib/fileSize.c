/******************************************************************************

  Copyright (c) 1995-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: fileSize.c,v 1.2 1999/04/06 13:10:04 nick Exp $

******************************************************************************/

#include <sys/stat.h>
#include <sys/types.h>
#include <assert.h>

#include "goodies.h"

off_t
IB_fileSize(int fildes) {
  struct stat buf;

  assert(fildes >= 0);

  if (fstat(fildes, &buf) == -1) {
    return -1;
  }

  return buf.st_size;
}
