/******************************************************************************

  Copyright (c) 1996-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: logOptions.c,v 1.2 1999/04/23 19:30:29 nick Exp $

******************************************************************************/

#include <stddef.h>

#include "log.h"

/* Default options structure. */
const IB_logOptions_t IB_logOptions __attribute__((weak)) = {
  NULL, /* ident */
  0     /* option */
};
