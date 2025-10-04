/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: typedefs.hh,v 1.6 1999/04/23 11:21:29 nick Exp $

******************************************************************************/

#ifndef _TYPEDEFS_HH
#define _TYPEDEFS_HH

#include <sys/types.h>

typedef enum {
  commodityGAsChips,
  commodityBioChips,
  commodityMasers,
  commodityWeapons,
  commodityVidicasters,
  commodityElectros,
  commodityTools,
  commoditySynths,
  commodityDroids,
  commodityAntiMatter,
  commodityPowerPacks,
  commodityControllers,
  commodityGenerators,
  commodityPolymers,
  commodityLubOils,
  commodityPharmaceuticals,
  commodityPetrochemicals,
  commodityRNA,
  commodityPropellants,
  commodityExplosives,
  commodityLanzariK,
  commodityNitros,
  commodityMunitions,
  commodityMechParts,
  commodityCereals,
  commodityWoods,
  commodityHides,
  commodityTextiles,
  commodityMeat,
  commoditySpices,
  commodityFruit,
  commoditySoya,
  commodityLivestock,
  commodityFurs,
  commodityRadioactives,
  commodityNickel,
  commodityXmetals,
  commodityCrystals,
  commodityAlloys,
  commodityGold,
  commodityMonopoles,
  commodityHypnotapes,
  commodityStudios,
  commoditySensAmps,
  commodityGames,
  commodityArtifacts,
  commodityKatydidics,
  commodityMusiks,
  commodityLibraries,
  commodityHolos,
  commodityUnivators,
  commoditySimulations
} commodity_t;

typedef enum {
  groupNone = -1,
  groupAgricultural = 1,
  groupMining,
  groupIndustrial,
  groupTechnological,
  groupLeisure
} commodity_group_t;

typedef enum { // Factory delivery points
  deliverExchange,
  deliverWarehouse,
  deliverFactory
} delivery_t;

typedef enum {
  directionNorth,
  directionNE,
  directionEast,
  directionSE,
  directionSouth,
  directionSW,
  directionWest,
  directionNW,
  directionUp,
  directionDown,
  directionIn,
  directionOut,
  directionPlanet
} direction_t;

typedef enum {
  eContinue,
  eStop
} event_result_t;

typedef enum {
  hContinue,
  hStop
} hook_result_t;

typedef enum {
  levelNoProduction,
  levelAgricultural,
  levelMining,
  levelIndustrial,
  levelTechnological,
  levelLeisure,
  levelCapital
} level_t;

typedef enum {
  // Explorer builds
  projectLink = 1,

  // Planetary investment builds
  projectEducation,
  projectEnergy,
  projectHealth,
  projectInfra,
  projectSecurity,

  // Duke puzzle builds
  projectUpside,
  projectDownside,
  projectAccel,
  projectMatProc,
  projectC3,
  projectMatTrans,
  projectTimeMach
} project_t;

typedef enum {
  rankGroundHog,
  rankCommander,
  rankCaptain,
  rankAdventurer,
  rankTrader,
  rankMerchant,
  rankJP,
  rankGM,
  rankExplorer,
  rankSquire,
  rankThane,
  rankIndustrialist,
  rankTechnocrat,
  rankBaron,
  rankDuke,
  rankSenator,
  rankTreasurer,
  rankHostess,
  rankSubManager,
  rankManager,
  rankDeity,
  rankEmperor
} rank_t;

typedef enum {
  sexFemale = 'f',
  sexMale = 'm',
  sexNeuter = 'n'
} sex_t;

typedef struct {
  bool expand;
  const char *text;
} message_t;

typedef unsigned msg_id_t;

typedef enum {
  traceNone,
  tracePerivale
} trace_t;

#endif // _TYPEDEFS_HH
