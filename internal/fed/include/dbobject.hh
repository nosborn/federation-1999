/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: dbobject.hh,v 1.3 1999/04/11 16:56:11 nick Exp $

******************************************************************************/

#ifndef _DBOBJECT_HH
#define _DBOBJECT_HH

#include <common/object.hh>
#include <typedefs.hh>

struct dbobject { // Object record
  // Common properties

  int number;                   /* object's vocab number */
  unsigned flags;               /* flags class to which the object belongs */
  char name[NAME_SIZE];         /* mobile name */
  char desc[82];                /* descriptions of object */
  char scan[201];               /* desription for examine */
  sex_t sex;                    /* m or f or n */
  unsigned short cur_loc;       /* current locations */
  unsigned short start_loc;     /* start locations */
  unsigned short weight;        /* wt of object in strength units */
  unsigned short get_event;     /* event when player tries to get object */
  unsigned short give_event;    /* event when player gives object */
  unsigned short consume_event; /* event when drinking/eating */
  int value;                    /* value/price on obj
                                   -ve goes on killer's reward! */

  // Object specific properties

  unsigned short drop_event; /* event when player drops object */
  unsigned short offset;     /* max value of offset for object recycling */

  // Mobile specific properties

  unsigned short max_loc; /* high/low loc for mobile */
  unsigned short min_loc;
  old_s_guns_t ship_guns[4];
  unsigned short hull;             // Hull strength
  unsigned short shield;           // Shield strength
  unsigned short engine;           // Engine capacity
  unsigned short computer;         // Computer size
  unsigned short fuel;             // Fuel
  unsigned short hold;             // Cargo capacity
  unsigned short tonnage;          // Overall size of ship
  unsigned short attack_percent;   /* % chance of mobile attacking */
  unsigned short kill_event;       /* event when player tries to kill mobile */
  short pref_object;               /* item which mobile will pay double for! */
  short move_counter, max_counter; /* counter for movement & counter
                                 reset level. -ve move_counter = immobile */
};

typedef struct dbobject dbobject_t;

#endif // _DBOBJECT_HH
