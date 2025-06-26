/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: dbevent.hh,v 1.3 1999/04/11 16:56:11 nick Exp $

******************************************************************************/

#ifndef _DBEVENT_HH
#define _DBEVENT_HH

#include <fed.hh>

struct dbevent {
  unsigned short type;        // Type of event:
                              // 1 = change persona attr,
                              // 3 = persona attr test for min val,
                              // 5 = persona attr test for max val,
                              // 7 = toggle object description,
                              // 8 = misc special changes,
                              // 9 = check for object when entering location
  char desc[EVENT_DESC_SIZE]; // Text output to user
  short field_1;              // hull/strength change
  short field_2;              // shield/stamina change
  short field_3;              // engine/intelligence change
  short field_4;              // dexterity change
  short field_7;              // change amount for type 3 - 6 events
  short field_8;              // object number, +ve carried/-ve not carried
  unsigned short new_loc;     // new location to move player/ship to
};

typedef struct dbevent dbevent_t;

#endif /* _DBEVENT_HH */
