/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: fed.hh,v 1.3 1999/04/06 14:28:39 nick Exp $

******************************************************************************/

#ifndef _FED_HH
#define _FED_HH

#include <limits.h>
#include <time.h>

#include <ibgames/log.h>
#include <ibgames.h>

#define FE_WILL_FORMAT
#undef SEND_MAIL
#undef USER_LOG
#undef VERBOSE_LOGON

#define GENERAL_JOB_AGING 5
#define JOB_DELIVERY_AGE 20
#define MAX_XT_CHANNEL 26
#define MIN_IB_RANK SENATOR

#define MOOD_SIZE 36
#define NAME_SIZE 16

#define COMPANY_NAME_SIZE 32
#define EVENT_DESC_SIZE 160
#define OBJECT_DESC_SIZE 82
#define OBJECT_SCAN_SIZE 201
#define SHIP_DESC_SIZE 160

#define MAX_GUNS 8  // Max number of guns a ship can carry
#define MAX_LOAD 15 // Max cargo loads a ship can carry
#define MAX_STORES 10

#define MAX_BALANCE (LONG_MAX - 1L)
#define MIN_BALANCE (LONG_MIN + 1L)

#define MAX_INPUTS 6 // Maximum inputs to make a commodity

/* Game constants */
#define DESC_SIZE 1024     /* size of location text */
#define EVENT_SIZE 2       /* max number of events per location */
#define FIRST_COMMOD 10000 /* vocab number of first commodity */
#define LAST_COMMOD 10051  /* vocab number of last commodity */
#define MAX_FACTORY 12     /* Number of factories in persona file record */
#define MAX_HOARDING 120   /* Maximum object hoarding time */
#define MAX_PER_DESC 152   // Max size of player desription field
#define MIN_HOARDING 90    /* Minimum object hoarding time (minutes) */
#define SHIP_START 426     /* Sol start location for new space ships */

#define SPYBEAM_COST 10000000   // 10,000,000 IG
#define SPYBEAM_RESALE 5000000  //  5,000,000 IG
#define SPYSCREEN_COST 50000000 // 50,000,000 IG

#define SPYBEAM_WEIGHT 50 //         50 tons

// The number of seconds in a 24 hour period.
#define SECS_IN_A_DAY (24 * 60 * 60)

/* # defines for get_cache() & flush_cache() */
#define NEW_REC -1L

#define DIM_OF(array) (sizeof(array) / sizeof(array[0]))
#define LOG(fmt, args...) IB_log(fmt, ##args)

// Debugging support.

#ifdef NDEBUG

# define dbgPrecondition(expr) ((void) 0)
# define dbgCheck(expr) ((void) 0)
# define dbgTrace(fmt, args...) ((void) 0)

#else /* NDEBUG */

# include <assert.h>
# include <stdlib.h>
# include <string.h>

enum {
  debugCheck = 0x01,
  debugPrecondition = 0x02,
  debugTrace = 0x04
};

extern unsigned debugEnableFlags;

inline bool
isDebugEnabled(unsigned theFlag) {
  return (debugEnableFlags & theFlag);
}

# define dbgCheck(expr)            \
   if (isDebugEnabled(debugCheck)) \
   assert(expr)

# define dbgPrecondition(expr)            \
   if (isDebugEnabled(debugPrecondition)) \
   assert(expr)

# define dbgIsValidString(buf) \
   (memchr(buf, '\0', sizeof(buf)) != NULL)

# define dbgTrace(fmt, args...)    \
   if (isDebugEnabled(debugTrace)) \
   IB_log(fmt, ##args)

#endif /* NDEBUG */
#endif /* _FED_HH */
