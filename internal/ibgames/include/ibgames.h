/******************************************************************************

  Copyright (c) 1997-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: ibgames.h,v 1.3 1999/03/17 12:16:21 nick Exp $

******************************************************************************/

#ifndef __IBGAMES_H
#define __IBGAMES_H

#include <limits.h>

/*
 * Unique account identifier.
 */
typedef unsigned int account_id_t;

/*
 * The lowest and highest real account IDs. Values below 100000 were used for
 * Federation personas moved from AOL and don't have corresponding account
 * details.
 */
#define MIN_ACCOUNT_ID 100000
#define MAX_ACCOUNT_ID INT_MAX

#endif /* __IBGAMES_H */
