/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: location.hh,v 1.2 1999/04/11 16:56:12 nick Exp $

******************************************************************************/

#ifndef _COMMON_LOCATION_HH
#define _COMMON_LOCATION_HH

// Location flags.

enum {
  lfDark = 0x00000001,   // location is unlit [L]
  lfSpace = 0x00000002,  // location is in space [S]
  lfDeath = 0x00000004,  // location kills player [D]
  lfVacuum = 0x00000008, // vacuum location [V]
  lfTrade = 0x00000010,  // trading exchange [T]
  lfYard = 0x00000020,   // ship yard [Y]
  //
  lfLink = 0x00000080, // interstellar link [I]
  //
  lfGen = 0x00000200,      // general store [G]
  lfWeap = 0x00000400,     // weapon shop [W]
  lfCafe = 0x00000800,     // cafe/bar [R]
  lfRep = 0x00001000,      // ship repairer [B]
  lfCom = 0x00002000,      // comms shop [E]
  lfClth = 0x00004000,     // clothing shop [F]
  lfPeace = 0x00008000,    // no fighting [P]
  lfHospital = 0x00010000, // hospital location [H]
  lfIns = 0x00020000,      // Insurance broker
  lfLock = 0x00040000,     // Lockable - dropped objects recycled
  //
  lfShield = 0x00100000,   // Location is teleport-shielded
  lfLanding = 0x00200000,  // Landing pad
  lfOrbit = 0x00400000,    // Planetary orbit
  lfIndoors = 0x00800000,  // Location is indoors
  lfOutdoors = 0x01000000, // Location is outdoors
  lfHidden = 0x02000000    // Pretend location doesn't exist
};

// Allowable flags for a player planet SPACE location.

#define lfSpaceAllowed (lfSpace + lfDeath + lfVacuum + lfLink + lfPeace + lfOrbit)

// Disallowed flags for a player planet non-SPACE location.

#define lfGroundDenied (lfSpace + lfLink + lfCommand + lfOrbit)

// Symbolic names for mov_tab entries.

enum {
  mvNorth = 0,
  mvNE,
  mvEast,
  mvSE,
  mvSouth,
  mvSW,
  mvWest,
  mvNW,
  mvUp,
  mvDown,
  mvIn,
  mvOut,
  mvPlanet
};

#endif /* _COMMON_LOCATION_HH */
