/******************************************************************************

  Copyright (c) 1995-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: read.c,v 1.2 1999/04/06 13:10:04 nick Exp $

******************************************************************************/

#include <assert.h>
#include <errno.h>
#include <string.h>
#include <unistd.h>

#include "goodies.h"

ssize_t
IB_read(int fildes, void *buf, size_t nbytes) {
  ssize_t bytesRead;

  assert(fildes >= 0);
  assert(buf != NULL);
  assert(nbytes > 0);

  do {
    bytesRead = read(fildes, buf, nbytes);
  } while (bytesRead == -1 && errno == EINTR);

  return bytesRead;
}
