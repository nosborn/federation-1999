/******************************************************************************

  Copyright © 1998-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: rules.c,v 1.2 1999/03/17 12:16:22 nick Exp $

******************************************************************************/


#include <sys/types.h>
#include <assert.h>
#include <limits.h>
#include <stdio.h>
#include <unistd.h>

#include <ibgames.h>

#include "goodies.h"
#include "rules.h"


#define FALSE 0
#define TRUE  1


/*-----------------------------------------------------------------------------

  i s L o c k e d O u t

-----------------------------------------------------------------------------*/

bool
isLockedOut( const account_id_t uid )
{
   assert(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);

   if (access(rulesLockFile(uid), F_OK) == 0) {
      return TRUE;
   }

   return FALSE;
}


/*-----------------------------------------------------------------------------

  r u l e s L o c k F i l e

-----------------------------------------------------------------------------*/

const char*
rulesLockFile( const account_id_t uid )
{
   static char lockFile[PATH_MAX];

   assert(uid >= MIN_ACCOUNT_ID && uid <= MAX_ACCOUNT_ID);
   sprintf(lockFile, "%s/lock/%u", IB_homeDir(), uid);

   return lockFile;
}
