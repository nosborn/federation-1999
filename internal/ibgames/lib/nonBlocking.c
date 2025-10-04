/******************************************************************************

  Copyright (c) 1995-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: nonBlocking.c,v 1.2 1999/04/06 13:10:04 nick Exp $

******************************************************************************/

#include <assert.h>
#include <fcntl.h>
#include <unistd.h>

#include "goodies.h"

int
IB_nonBlocking(int fildes) {
  int ioval;

  assert(fildes >= 0);

  ioval = fcntl(fildes, F_GETFL, 0);

  if (ioval < 0) {
    return -1;
  }

  return fcntl(fildes, F_SETFL, ioval | O_NONBLOCK);
}
