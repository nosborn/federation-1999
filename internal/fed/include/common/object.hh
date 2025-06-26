/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: object.hh,v 1.4 1999/04/11 16:56:12 nick Exp $

******************************************************************************/

#ifndef _COMMON_OBJECT_HH
#define _COMMON_OBJECT_HH

#include <fed.hh>

// Container class definitions
enum {
  ofLight = 0x00000001,   // Object is a light source (T)
  ofLiquid = 0x00000040,  // Object is a liquid (H)
  ofEdible = 0x00000080,  // Object is edible (E)
  ofAnimate = 0x00000200, // Object is a mobile (M)
  ofShip = 0x00000400,    // Mobile is spaceship (S)

  ofNoThe = 0x00002000,   // Don't insert a 'the' if set
  ofMusic = 0x00004000,   // Object is a music instrument
  ofCleaner = 0x00008000, // Mobile is a cleaning droid

  ofIndoors = 0x01000000,  // Must stay in INDOOR locations
  ofOutdoors = 0x02000000, // Must stay in OUTDOOR locations

  ofStoic = 0x10000000, // No interactions
  ofDuke = 0x20000000,  // Object is set up for the Duke puzzle
  ofHidden = 0x40000000 // Object is out of the game
};

typedef struct { // Old style ship mounted weapons - for mobiles
  short type;    // Weapon type
  char name[20]; // Weapon name
  short damage;  // Base damage
  short power;   // Power consumption
} old_s_guns_t;

#endif /* _COMMON_OBJECT_HH */
