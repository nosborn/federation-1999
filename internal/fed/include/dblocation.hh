/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: dblocation.hh,v 1.2 1999/04/11 16:56:11 nick Exp $

******************************************************************************/

#ifndef _DBLOCATION_HH
#define _DBLOCATION_HH

#include <common/location.hh>
#include <fed.hh>

struct dblocation {           // Map location record (in database)
  char desc[DESC_SIZE];       // Description array
  unsigned short events[2];   // Array giving ref number's of event
                              // for this location
                              // [0] = enter event
                              // [1] = move in non-exit direction event
  unsigned map_flag;          // Location flags - 32 bits
  unsigned short mov_tab[13]; // Movement table
  unsigned short sys_loc;     // Message for non-movement
};

typedef struct dblocation dblocation_t;

#endif /* _DBLOCATION_HH */
