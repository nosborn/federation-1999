/******************************************************************************

  Copyright (c) 1995-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: close.c,v 1.2 1999/04/06 13:10:04 nick Exp $

******************************************************************************/

#include <assert.h>
#include <errno.h>
#include <unistd.h>

#include "goodies.h"

int
IB_close(int fildes) {
  int returnValue;

  assert(fildes >= 0);

  do {
    returnValue = close(fildes);
  } while (returnValue == -1 && errno == EINTR);

  return returnValue;
}
