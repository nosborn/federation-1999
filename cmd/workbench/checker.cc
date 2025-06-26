/******************************************************************************

  Copyright (c) 1987-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: checker.cc,v 1.5 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <algorithm>

#include <ctype.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <ibgames/goodies.h>

#include <dbevent.hh>
#include <dblocation.hh>
#include <dbobject.hh>
#include <workbench.h>

#include "fedwb.hh"

extern bool needsCheck;
extern bool noExchange;

static dbevent_t *eventData = NULL;
static dblocation_t *locationData = NULL;
static dbobject_t *objectData = NULL;

static size_t numEvents = 0;
static size_t numLocations = 0;
static size_t numObjects = 0;

static bool checkEvents();
static bool checkLocations();
static bool checkObjects();
static bool checkRoute(unsigned, unsigned);
static bool _CheckSys(int);
static bool findRoute(unsigned, unsigned, bool);
static bool loadData();
static bool _LocInfo();
static void _PrintBriefDesc(int);

// Checks to see that the various parameters of players' planets are within
// bounds. Returns true if no problems detected, false if a problem is
// detected.
bool
checkPlanet() {
  dbgPrecondition(eventData == NULL && numEvents == 0);
  dbgPrecondition(locationData == NULL && numLocations == 0);
  dbgPrecondition(objectData == NULL && numObjects == 0);

  bool checkedOK = false;

  if (loadData()) {
    if (checkLocations() && checkEvents() && checkObjects()) {
      checkedOK = true;
      needsCheck = false;
    }
  }

  if (eventData != NULL) {
    delete eventData;
    eventData = NULL;
  }

  if (locationData != NULL) {
    delete locationData;
    locationData = NULL;
  }

  if (objectData != NULL) {
    delete objectData;
    objectData = NULL;
  }

  numEvents = numLocations = numObjects = 0;

  if (!checkedOK) {
    puts("Planet check failed!");
  }

  return checkedOK;
}

// Checks the event file for dubious activities. Returns true if no errors
// found, false if errors found.
static bool
checkEvents() {
  Output(mnCheckingEvents);

  bool checkedOK = true;

  for (size_t i = 0; i < numEvents; i++) {
    const dbevent_t &thisEvent = eventData[i];
    const int eventNo = i + 1;

    Output(FormatText(mnCheckingEvent, eventNo));

    if (thisEvent.type != 1 && thisEvent.type != 3 && thisEvent.type != 9) {
      puts("Invalid event type");
      checkedOK = false;
      continue;
    }

    if (isBlank(thisEvent.desc)) {
      puts("Event text is missing.");
      checkedOK = false;
      continue;
    } else if (strchr(thisEvent.desc, '\n') != NULL) {
      puts("Event text may not contain blank lines.");
      checkedOK = false;
      continue;
    }

    if (thisEvent.new_loc != 0) {
      if (thisEvent.new_loc > numLocations) {
        puts("Moves to non-existant location!");
        checkedOK = false;
        continue;
      }

      if (thisEvent.new_loc < 9) {
        puts("Moves to spaceship location!");
        checkedOK = false;
        continue;
      }

      const dblocation_t &newLocation = locationData[thisEvent.new_loc - 1];

      if ((newLocation.map_flag & lfSpace) != 0) {
        puts("Moves to SPACE location!");
        checkedOK = false;
        continue;
      }
    }

    if (thisEvent.type == 9) {
      const int objectNumber = abs(thisEvent.field_8);
      bool objectFound = false;

      for (size_t j = 0; j < numObjects; j++) {
        if (objectData[j].number == objectNumber) {
          objectFound = true;
          break;
        }
      }

      if (!objectFound) {
        puts("Requires non-existent object!");
        checkedOK = false;
        continue;
      }
    }

    puts("OK");
  }

  if (checkedOK) {
    Output(mnCheckEventsOK);
  } else {
    Output(mnCheckEventsError);
  }

  return checkedOK;
}

// Checks to see that the movement tables don't call out of range locations.
// Returns true if all is OK, false if out of range locs.
static bool
checkLocations() {
  Output(mnCheckingLocations);

  bool checkedOK = true;

  unsigned hospitalLocationNo = 0;
  unsigned landingLocationNo = 0;
  unsigned linkLocationNo = 0;
  unsigned orbitLocationNo = 0;
  unsigned tradeLocationNo = 0;

  for (size_t locationIndex = 8; locationIndex < numLocations; locationIndex++) {
    const dblocation_t &thisLocation = locationData[locationIndex];
    const int locationNo = locationIndex + 1;

    Output(FormatText(mnCheckingLocation, locationNo));

    const unsigned deathFlag = (thisLocation.map_flag & lfDeath);
    const unsigned peaceFlag = (thisLocation.map_flag & lfPeace);
    const unsigned spaceFlag = (thisLocation.map_flag & lfSpace);

    // Check desc.

    size_t count;
    bool allSpaces = true;

    for (count = 0; count < 79; count++) {
      if (thisLocation.desc[count] == '\n') {
        break;
      }

      // If there's only one line in the description, there won't be a
      // new-line. I'm not sure if this is the way it should be!

      if (thisLocation.desc[count] == '\0') {
        count = 0;
        break;
      }

      if (isgraph(thisLocation.desc[count]) &&
          static_cast<unsigned char>(thisLocation.desc[count]) != 0xA0U) {
        allSpaces = false;
      }
    }

    if (count == 0 || count == 79 || allSpaces) {
      puts("No short description!");
      checkedOK = false;
      continue;
    }

    // Check events.

    if (thisLocation.events[0] != 0) {
      const unsigned eventNo = thisLocation.events[0];

      if (eventNo > numEvents) {
        puts("Non-existant IN event!");
        checkedOK = false;
        continue;
      }

      if (spaceFlag == lfSpace) {
        puts("Useless IN event in SPACE!");
        checkedOK = false;
        continue;
      }
    }

    if (thisLocation.events[1] != 0) {
      const unsigned eventNo = thisLocation.events[1];

      if (eventNo > numEvents) {
        puts("Non-existant OUT event!");
        checkedOK = false;
        continue;
      }

      if (spaceFlag == lfSpace) {
        puts("Useless OUT event in SPACE!");
        checkedOK = false;
        continue;
      }
    }

    // Check map_flag.

    if ((thisLocation.map_flag & lfHospital) != 0) {
      if (hospitalLocationNo != 0) {
        puts("Can't have multiple Hospitals!");
        checkedOK = false;
        continue;
      }

      hospitalLocationNo = locationNo;
    }

    if ((thisLocation.map_flag & lfLanding) != 0) {
      if (landingLocationNo != 0) {
        puts("Can't have multiple Landing Pads!");
        checkedOK = false;
        continue;
      }

      landingLocationNo = locationNo;
    }

    if ((thisLocation.map_flag & lfLink) != 0) {
      if (linkLocationNo != 0) {
        puts("Can't have multiple Interstellar Links!");
        checkedOK = false;
        continue;
      }

      linkLocationNo = locationNo;
    }

    if ((thisLocation.map_flag & lfOrbit) != 0) {
      if (orbitLocationNo != 0) {
        puts("Can't have multiple Planetary Orbits!");
        checkedOK = false;
        continue;
      }

      orbitLocationNo = locationNo;
    }

    if ((thisLocation.map_flag & lfTrade) != 0) {
      if (tradeLocationNo != 0) {
        puts("Can't have multiple Trading Exchanges!");
        checkedOK = false;
        continue;
      }

      tradeLocationNo = locationNo;
    }

    if ((thisLocation.map_flag & lfCafe) == lfCafe ||
        (thisLocation.map_flag & lfClth) == lfClth ||
        (thisLocation.map_flag & lfCom) == lfCom ||
        (thisLocation.map_flag & lfGen) == lfGen ||
        (thisLocation.map_flag & lfHospital) == lfHospital ||
        (thisLocation.map_flag & lfIns) == lfIns ||
        (thisLocation.map_flag & lfLanding) == lfLanding ||
        (thisLocation.map_flag & lfLink) == lfLink ||
        (thisLocation.map_flag & lfOrbit) == lfOrbit ||
        (thisLocation.map_flag & lfRep) == lfRep ||
        (thisLocation.map_flag & lfTrade) == lfTrade ||
        (thisLocation.map_flag & lfWeap) == lfWeap ||
        (thisLocation.map_flag & lfYard) == lfYard) {
      if (deathFlag == lfDeath) {
        puts("Can't be a death location!");
        checkedOK = false;
        continue;
      }
    }

    if ((thisLocation.map_flag & lfLink) != 0 ||
        (thisLocation.map_flag & lfOrbit) != 0) {
      if (peaceFlag == 0) {
        puts("Must be a peaceful location!");
        checkedOK = false;
        continue;
      }

      if (spaceFlag == 0) {
        puts("Must be a space location!");
        checkedOK = false;
        continue;
      }
    }

    if ((thisLocation.map_flag & lfCafe) == lfCafe ||
        (thisLocation.map_flag & lfClth) == lfClth ||
        (thisLocation.map_flag & lfCom) == lfCom ||
        (thisLocation.map_flag & lfGen) == lfGen ||
        (thisLocation.map_flag & lfHospital) == lfHospital ||
        (thisLocation.map_flag & lfIns) == lfIns ||
        (thisLocation.map_flag & lfLanding) == lfLanding ||
        (thisLocation.map_flag & lfRep) == lfRep ||
        (thisLocation.map_flag & lfTrade) == lfTrade ||
        (thisLocation.map_flag & lfWeap) == lfWeap ||
        (thisLocation.map_flag & lfYard) == lfYard) {
      if (spaceFlag == lfSpace) {
        puts("Can't be a space location!");
        checkedOK = false;
        continue;
      }
    }

    if (spaceFlag == lfSpace) {
      const unsigned allowedFlags = lfDeath |
                                    lfLink |
                                    lfOrbit |
                                    lfPeace |
                                    lfSpace |
                                    lfVacuum;

      if ((thisLocation.map_flag & ~allowedFlags) != 0) {
        puts("One or more silly flags set with SPACE!");
        checkedOK = false;
        continue;
      }
    } else {
      const unsigned disallowedFlags = lfLink |
                                       lfOrbit |
                                       lfSpace;

      if ((thisLocation.map_flag & disallowedFlags) != 0) {
        puts("One or more silly flags set without SPACE!\n");
        checkedOK = false;
        continue;
      }
    }

    // Check mov_tab.

    bool movementOK = true;

    for (size_t moveIndex = 0; moveIndex < 13; moveIndex++) {
      if (thisLocation.mov_tab[moveIndex] == 0) {
        continue;
      }

      if (deathFlag == lfDeath) {
        puts("Useless exit from DEATH location!");
        movementOK = false;
        break;
      }

      const unsigned toLocationNo = thisLocation.mov_tab[moveIndex];

      if (toLocationNo > numLocations) {
        puts("Movement to non-existant location!");
        movementOK = false;
        break;
      }

      if (toLocationNo < 9) {
        puts("Movement to spaceship location!");
        movementOK = false;
        break;
      }

      const dblocation_t &toLocation = locationData[toLocationNo - 1];

      if ((toLocation.map_flag & lfSpace) != spaceFlag) {
        if (spaceFlag) {
          puts("Movement from space to ground!");
        } else {
          puts("Movement from ground to space!");
        }

        movementOK = false;
        break;
      }

      if (locationNo == hospitalLocationNo) {
        if ((toLocation.map_flag & lfDeath) != 0) {
          puts("Movement from HOSPITAL to DEATH location!");
          movementOK = false;
          break;
        }
      }
    }

    if (!movementOK) {
      checkedOK = false;
      continue;
    }

    if (spaceFlag == lfSpace) {
      if (thisLocation.mov_tab[mvIn] != 0 ||
          thisLocation.mov_tab[mvOut] != 0) {
        puts("IN/OUT movement is not valid in SPACE locations!");
        checkedOK = false;
        continue;
      }
    } else {
      if (thisLocation.mov_tab[mvPlanet] != 0) {
        puts("PLANET movement is only valid in SPACE locations!");
        checkedOK = false;
        continue;
      }
    }

    // Check sys_loc.

    if (!_CheckSys(locationIndex)) {
      puts("Out of range system message!");
      checkedOK = false;
      continue;
    }

    puts("OK");
  }

  if (checkedOK) {
    checkedOK = _LocInfo();
  }

  if (checkedOK) {
    if (linkLocationNo != 0 && orbitLocationNo != 0) {
      printf("Checking route from Link to Orbit... ");

      if (!checkRoute(linkLocationNo, orbitLocationNo)) {
        checkedOK = false;
      }

      printf("Checking route from Orbit to Link... ");

      if (!checkRoute(orbitLocationNo, linkLocationNo)) {
        checkedOK = false;
      }
    }

    if (landingLocationNo != 0 && tradeLocationNo != 0 && !noExchange) {
      printf("Checking route from Landing Pad to Exchange... ");

      if (!checkRoute(landingLocationNo, tradeLocationNo)) {
        checkedOK = false;
      }

      printf("Checking route from Exchange to Landing Pad... ");

      if (!checkRoute(tradeLocationNo, landingLocationNo)) {
        checkedOK = false;
      }
    }

    if (hospitalLocationNo != 0 && landingLocationNo != 0) {
      printf("Checking route from Hospital to Landing Pad... ");

      if (!checkRoute(hospitalLocationNo, landingLocationNo)) {
        checkedOK = false;
      }
    }
  }

  if (checkedOK) {
    puts("\nNo errors found in location file!");
  } else {
    puts("\n*** Errors found in location file! ***");
  }

  return checkedOK;
}

// Checks the objects file to make sure that there is nothing untoward. Returns
// true if everything is OK, false otherwise.
static bool
checkObjects() {
  Output(mnCheckingObjects);

  bool flag = true;

  for (size_t objectIndex = 0; objectIndex < numObjects; objectIndex++) {
    dbobject_t *obj = &objectData[objectIndex];

    // Common stuff.
    printf("%-18s", obj->name);

    if (!isalpha(obj->name[0])) {
      puts("The name must begin with a letter.");
      flag = false;
      continue;
    }

    const size_t length = strlen(obj->name);

    if (length < 3) {
      puts("The name must be at least 3 characters long.");
      // We used to allow 2 character names even though we shouldn't have. This
      // is a bodge to avoid breaking planets that have them.
      if (length < 2) {
        flag = false;
        continue;
      }
    }

    for (size_t index = 1; index < length; index++) {
      if (!isalnum(obj->name[index]) && obj->name[index] != '-') {
        puts("Names can only use letters, numbers and hyphens.");
        flag = false;
        continue;
      }
    }

    if (isBlank(obj->desc)) {
      puts("Description text is missing.");
      flag = false;
      continue;
    } else if (strchr(obj->desc, '\n') != NULL) {
      puts("Description text may not contain blank lines.");
      flag = false;
      continue;
    }

    if (isBlank(obj->scan)) {
      puts("Scan text is missing.");
      flag = false;
      continue;
    } else if (strchr(obj->scan, '\n') != NULL) {
      puts("Scan text may not contain blank lines.");
      flag = false;
      continue;
    }

    if (obj->start_loc < 9 || obj->start_loc > numLocations) {
      printf("\n    *** Start location (%d) out of range (1-%zu)! ***\n",
             obj->start_loc, numLocations);
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->get_event > numEvents) {
      puts("\n   *** Get event is out of range ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->give_event > numEvents) {
      puts("\n   *** Give event is out of range ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->consume_event > numEvents) {
      puts("\n   *** Consume event is out of range ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->drop_event > numEvents) {
      puts("\n   *** Drop event is out of range ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->start_loc + obj->offset > numLocations) {
      puts("\n   *** Recycle offset will take object out of location range! ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if ((obj->flags & ofAnimate) == 0) {
      if (obj->value < 0) {
        puts("\n   *** Value of an object can't be negative ***");
        flag = false;
        continue;
      }

      putchar('\n');
      continue;
    }

    /* mobile stuff */
    if (obj->max_loc < 9 || obj->max_loc > numLocations) {
      puts("\n   *** Maximum location for movement is out of range! ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->min_loc < 9 || obj->min_loc > numLocations) {
      puts("\n   *** Minimum location for movement is out of range! ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    if (obj->min_loc > obj->max_loc) {
      puts("\n   *** The lowest location for movement is higher than the highest! ***");
      flag = false;
      continue;
    } else {
      putchar('.');
    }

    unsigned long spaceFlag;

    if ((obj->flags & ofShip) != 0) {
      spaceFlag = lfSpace;
    } else {
      spaceFlag = 0;
    }

    for (unsigned locNo = obj->min_loc; locNo <= obj->max_loc; locNo++) {
      if ((locationData[locNo - 1].map_flag & lfSpace) == spaceFlag) {
        continue;
      }

      if (spaceFlag == lfSpace) {
        puts("\n   *** Moves through non-SPACE location! ***");
      } else {
        puts("\n   *** Moves through SPACE location! ***");
      }

      flag = false;
      break;
    }

    if (!flag) {
      continue;
    }

    if (obj->pref_object > 0) {
      bool obj_flag = false;

      for (size_t index = 0; index < numObjects; index++) {
        if (obj->pref_object == objectData[index].number) {
          obj_flag = true;
          break;
        }
      }

      if (!obj_flag) {
        printf("\n   *** Preferred object (#%d) doesn't exist! ***\n",
               obj->pref_object);
        flag = false;
      }
    } else {
      putchar('.');
    }

    putchar('\n');
  }

  if (flag) {
    puts("\nNo errors found in objects file!\n");
  } else {
    puts("\n*** Errors found in objects file! ***\n");
  }

  return flag;
}

static bool
checkRoute(unsigned fromLocationNo, unsigned toLocationNo) {
  // This will mung the location and event data, so save a copy first.

  const size_t locationSize = sizeof(dblocation_t) * numLocations;
  void *locationCopy = alloca(locationSize);
  dbgCheck(locationCopy != NULL);
  memcpy(locationCopy, locationData, locationSize);

  const size_t eventSize = sizeof(dbevent_t) * numEvents;
  void *eventCopy = (eventSize > 0) ? alloca(eventSize) : NULL;

  if (eventSize > 0) {
    dbgCheck(eventCopy != NULL);
    memcpy(eventCopy, eventData, eventSize);
  }

  //

  const bool routeOK = findRoute(fromLocationNo, toLocationNo, false);

  if (routeOK) {
    puts("OK");
  } else {
    puts("Failed!");
  }

  // Restore the original data.

  memcpy(locationData, locationCopy, locationSize);

  if (eventSize > 0) {
    memcpy(eventData, eventCopy, eventSize);
  }

  //

  return routeOK;
}

// Check that the system messages are the one that are allowed.
// Params: Index to loc array
// Returns true if in range, false if out of range.
static bool
_CheckSys(int index) {
  static const unsigned allowable[] = {
    0, 1, 2, 7, 12, 23, 26, 30, 32, 33, 351, 352
  };

  for (size_t i = 0; i < DIM_OF(allowable); i++) {
    if (locationData[index].sys_loc == allowable[i]) {
      return true;
    }
  }

  return false;
}

// FIXME: Well actually, not a FIXME, but review this function when making
// changes to the main event handling code. In particular, this assumes that IN
// events and DEATH flags are ignored when moved to a location by an event.
static bool
findRoute(unsigned fromLocationNo, unsigned toLocationNo, bool checkEntry) {
  dblocation_t *thisLocation = &locationData[fromLocationNo - 1];

  if (checkEntry && thisLocation->events[0] != 0) {
    dbevent_t *thisEvent = &eventData[thisLocation->events[0] - 1];

    if (thisEvent->field_2 < 0) {
      return false; // IN event removes stamina.
    }

    if (thisEvent->new_loc > 0) {
      if (thisEvent->new_loc == USHRT_MAX) {
        return false; // Give up, we're looping
      }

      const int nextLocationNo = thisEvent->new_loc;
      thisEvent->new_loc = USHRT_MAX;

      if (findRoute(nextLocationNo, toLocationNo, false)) {
        return true; // We got there alive!
      }
    }
  }

  if (checkEntry && (thisLocation->map_flag & lfDeath) != 0) {
    return false; // Plain DEATH location.
  }

  if (fromLocationNo == toLocationNo) {
    return true; // We got there alive!
  }

  if (thisLocation->events[1] != 0) {
    dbevent_t *thisEvent = &eventData[thisLocation->events[1] - 1];

    if (thisEvent->field_2 < 0) {
      return false; // OUT event removes stamina.
    }

    if (thisEvent->new_loc > 0) {
      bool mustFire = true;

      for (size_t i = 0; i < DIM_OF(thisLocation->mov_tab); i++) {
        if (thisLocation->mov_tab[i] != 0) {
          mustFire = false;
          break;
        }
      }

      if (mustFire) {
        if (thisEvent->new_loc == USHRT_MAX) {
          return false; // Give up, we're looping
        }

        const int nextLocationNo = thisEvent->new_loc;
        thisEvent->new_loc = USHRT_MAX;

        if (findRoute(nextLocationNo, toLocationNo, false)) {
          return true; // We got there alive!
        }
      }
    }
  }

  for (size_t i = 0; i < DIM_OF(thisLocation->mov_tab); i++) {
    if (thisLocation->mov_tab[i] > 0) {
      if (thisLocation->mov_tab[i] != USHRT_MAX) {
        const int nextLocationNo = thisLocation->mov_tab[i];
        thisLocation->mov_tab[i] = USHRT_MAX;

        if (findRoute(nextLocationNo, toLocationNo, true)) {
          return true; // We got there alive!
        }
      }
    }
  }

  return false; // We ran out of routes to try!
}

bool
isBlank(const char *line) {
  while (*line) {
    unsigned char ch = *line;

    if (ch == 0xA0U) { // nonbreakspace
      ch = ' ';
    }

    if (isgraph(ch)) {
      return false;
    }

    line++;
  }

  return true;
}

// Loads in the location, object and event  files and sets up the max record
// numbers. Returns true if there is no problem, false if unable to open files
// or out of memory.
static bool
loadData() {
  dbgPrecondition(eventData == NULL && numEvents == 0);
  dbgPrecondition(locationData == NULL && numLocations == 0);
  dbgPrecondition(objectData == NULL && numObjects == 0);

  // Load locations.

  int fd = OpenLocFile();

  if (fd == -1) {
    return false;
  }

  off_t fileSize = IB_fileSize(fd);

  if (fileSize == -1) {
    return false;
  }

  if (fileSize > 0) {
    numLocations = std::min(static_cast<size_t>(fileSize) / sizeof(dblocation_t), 8 + LOCATION_LIMIT);
    locationData = new (std::nothrow) dblocation_t[numLocations];

    if (locationData == NULL) {
      return false;
    }

    const size_t readSize = sizeof(dblocation_t) * numLocations;

    if (read(fd, locationData, readSize) != readSize) {
      return false;
    }
  }

  close(fd);
  Output(FormatText(mnCheckLocationCount, numLocations));

  // Load events.

  fd = OpenEvFile();

  if (fd == -1) {
    return false;
  }

  fileSize = IB_fileSize(fd);

  if (fileSize == -1) {
    return false;
  }

  if (fileSize > 0) {
    numEvents = std::min(static_cast<size_t>(fileSize) / sizeof(dbevent_t), EVENT_LIMIT);
    eventData = new (std::nothrow) dbevent_t[numEvents];

    if (eventData == NULL) {
      return false;
    }

    const size_t readSize = sizeof(dbevent_t) * numEvents;

    if (read(fd, eventData, readSize) != readSize) {
      return false;
    }
  }

  close(fd);
  Output(FormatText(mnCheckEventCount, numEvents));

  /* load objects */

  fd = OpenObjFile();

  if (fd == -1) {
    return false;
  }

  fileSize = IB_fileSize(fd);

  if (fileSize == -1) {
    return false;
  }

  if (fileSize > 0) {
    numObjects = std::min(static_cast<size_t>(fileSize) / sizeof(dbobject_t), OBJECT_LIMIT);
    objectData = new (std::nothrow) dbobject_t[numObjects];

    if (objectData == NULL) {
      return false;
    }

    const size_t readSize = sizeof(dbobject_t) * numObjects;

    if (read(fd, objectData, readSize) != readSize) {
      return false;
    }
  }

  close(fd);
  Output(FormatText(mnCheckObjectCount, numObjects));

  return true;
}

/*-----------------------------------------------------------------------------

  _ L o c I n f o

  Gives the explorer a list of the locs that have different flags set and
  check for unallowed multiples.

  Returns true if no unallowed multiples, false if unallowed multiples.

-----------------------------------------------------------------------------*/

static bool
_LocInfo() {
  static const char one[] = " *** There must be one, and only one";

  int flag = 0;
  bool return_flag = true;

  for (size_t i = 0; i < 20; i++) {
    int number = 0;

    switch (i) {
      case 0:
        puts("\nInterstellar link");
        flag = lfLink;
        break;

      case 1:
        puts("\nSpace locations");
        flag = lfSpace;
        break;

      case 2:
        puts("\nPlanetary orbit");
        flag = lfOrbit;
        break;

      case 3:
        puts("\nLanding pad");
        flag = lfLanding;
        break;

      case 4:
        puts("\nLockable locations");
        flag = lfLock;
        break;

      case 5:
        puts("\nUnlit locations");
        flag = lfDark;
        break;

      case 6:
        puts("\nDeath locs");
        flag = lfDeath;
        break;

      case 7:
        puts("\nVacuum locs");
        flag = lfVacuum;
        break;

      case 8:
        puts("\nTrading exchange");
        flag = lfTrade;
        break;

      case 9:
        puts("\nGeneral stores");
        flag = lfGen;
        break;

      case 10:
        puts("\nWeapon shops");
        flag = lfWeap;
        break;

      case 11:
        puts("\nCafe/Bars");
        flag = lfCafe;
        break;

      case 12:
        puts("\nClothes shops");
        flag = lfClth;
        break;

      case 13:
        puts("\nPeace locations");
        flag = lfPeace;
        break;

      case 14:
        puts("\nHospital");
        flag = lfHospital;
        break;

      case 15:
        puts("\nElectronics shops");
        flag = lfCom;
        break;

      case 16:
        puts("\nRepair shops");
        flag = lfRep;
        break;

      case 17:
        puts("\nShipyards");
        flag = lfYard;
        break;

      case 18:
        puts("\nInsurance offices");
        flag = lfIns;
        break;

      case 19:
        puts("\nTeleport-shielded areas");
        flag = lfShield;
        break;
    }

    for (size_t count = 9; count <= numLocations; count++) {
      if (locationData[count - 1].map_flag & flag) {
        number++;
        printf("%3zu ", count);
        _PrintBriefDesc(count);
      }
    }

    if (number == 0) {
      puts("     None!");
    } else {
      printf("     %d location(s)\n", number);
    }

    if (i == 0 && number != 1) {
      printf("%s Interstellar link! ***\n", one);
      return_flag = false;
    } else {
      if (i == 2 && number != 1) {
        printf("%s Planetary orbit! ***\n", one);
        return_flag = false;
      } else if (i == 3 && number != 1) {
        printf("%s Landing pad ! ***\n", one);
        return_flag = false;
      } else if (i == 8 && number != 1 && !noExchange) {
        printf("%s Trading exchange! ***\n", one);
        return_flag = false;
      }
    }
  }

  puts("\nLocations calling events");

  for (size_t count = 1; count <= numLocations; count++) {
    if (locationData[count - 1].events[0]) {
      printf("Event %d called on entry to loc %2zu - ",
             locationData[count - 1].events[0], count);
      _PrintBriefDesc(count);
    }

    if (locationData[count - 1].events[1]) {
      printf("Event %d called on wrong move in loc %2zu - ",
             locationData[count - 1].events[1], count);
      _PrintBriefDesc(count);
    }
  }

  return return_flag;
}

/*-----------------------------------------------------------------------------

  _ P r i n t B r i e f D e s c

  Prints out the brief desc of the specified location.

  Params: The loc number

-----------------------------------------------------------------------------*/

static void
_PrintBriefDesc(int loc_no) {
  char buffer[80];

  memcpy(buffer, locationData[loc_no - 1].desc, 79);
  buffer[79] = '\0';

  for (size_t i = 0; i < 79; i++) {
    if (buffer[i] == '\n') {
      buffer[i] = '\0';
      break;
    }
  }

  puts(buffer);
}
