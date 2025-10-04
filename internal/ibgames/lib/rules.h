/******************************************************************************

  Copyright © 1998-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: rules.h,v 1.1 1999/01/22 15:53:49 nick Exp $

******************************************************************************/


#ifndef RULES_H
#define RULES_H


#include <ibgames.h>


// Default values if not already defined. This is brain-dead (the
// library doesn't even use them) and should be done with per-game
// configuration files and suitable functions in the library instead.

#define DEFAULT_RULES_MANAGER "Hazed"
#define DEFAULT_RULES_MANAGER_EMAIL "fi@ibgames.com"
#define DEFAULT_RULES_INFO_EMAIL "rules@ibgames.net"
#define DEFAULT_RULES_INFO_URL "http://www.ibgames.net/t&c/rules.html"

// Perona name of the House Rules manager.

#ifndef RULES_MANAGER
# define RULES_MANAGER DEFAULT_RULES_MANAGER
#endif

// Email address of the House Rules manager.

#ifndef RULES_MANAGER_EMAIL
# define RULES_MANAGER_EMAIL DEFAULT_RULES_MANAGER_EMAIL
#endif

// Email address to which queries about House Rules should be directed.

#ifndef RULES_INFO_EMAIL
# define RULES_INFO_EMAIL DEFAULT_RULES_INFO_EMAIL
#endif

// URL where information about House Rules can be found.

#ifndef RULES_INFO_URL
# define RULES_INFO_URL DEFAULT_RULES_INFO_URL
#endif


#ifdef __cplusplus
extern "C" {
#else
# define bool int
#endif
bool isLockedOut( const account_id_t );
const char* rulesLockFile( const account_id_t );
#ifdef __cplusplus
}
#endif


#endif /* RULES_H */
