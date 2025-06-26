/******************************************************************************

  Copyright (c) 1996-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: log.h,v 1.4 1999/04/23 19:30:29 nick Exp $

******************************************************************************/

#ifndef __IBGAMES_LOG_H
#define __IBGAMES_LOG_H

#include <syslog.h>

#ifdef __cplusplus
extern "C" {
#endif

typedef struct {
  const char *ident;
  int option;
} IB_logOptions_t;

extern const IB_logOptions_t IB_logOptions;

int IB_log(const char *format, ...) __attribute__((format(printf, 1, 2)));
int IB_setLogOptions(const char *ident, int option);

#ifdef __cplusplus
}
#endif

#endif /* __IBGAMES_LOG_H */
