/******************************************************************************

  Copyright (c) 1987-1997 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: genesis.cc,v 1.2 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <ctype.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <ibgames/goodies.h>
#include <dbevent.hh>
#include <dblocation.hh>
#include <dbobject.hh>
#include <fed.hh>
#include <workbench.h>

#include "fedwb.hh"

static const char *chomp(const char *);
static bool dumpEvents();
static bool dumpLocations();
static bool dumpObjects();

static const struct {
  unsigned flag;
  const char *name;
} loc_flags[] = {
  { lfCafe, "CAFE" },
  { lfClth, "CLTH" },
  { lfCom, "COM" },
  { lfDark, "DARK" },
  { lfDeath, "DEATH" },
  { lfGen, "GEN" },
  { lfHospital, "HOSPITAL" },
  { lfIns, "INS" },
  { lfLanding, "LANDING" },
  { lfLink, "LINK" },
  { lfLock, "LOCK" },
  { lfOrbit, "ORBIT" },
  { lfPeace, "PEACE" },
  { lfRep, "REP" },
  { lfShield, "SHIELD" },
  { lfSpace, "SPACE" },
  { lfTrade, "TRADE" },
  { lfVacuum, "VACUUM" },
  { lfWeap, "WEAP" },
  { lfYard, "YARD" }
};

static const struct {
  unsigned flag;
  const char *name;
} ob_flags[] = {
  { ofAnimate, "ANIMATE" },
  { ofCleaner, "CLEANER" },
  { ofEdible, "EDIBLE" },
  { ofLight, "LIGHT" },
  { ofLiquid, "LIQUID" },
  { ofMusic, "MUSIC" },
  { ofNoThe, "NO_THE" },
  { ofShip, "SHIP" }
};

void
download() {
  printf("LAYOUT 1\n");

  if (dumpEvents() && dumpLocations() && dumpObjects()) {
    printf("OK\n");
  }
}

static bool
dumpEvents() {
  const int fd = OpenEvFile();

  if (fd == -1) {
    printf("FAIL Can't open event file\n");
    return false;
  }

  for (int eventNo = 1; eventNo <= EVENT_LIMIT; eventNo++) {
    dbevent_t event;
    size_t index;

    const ssize_t nbytes = IB_read(fd, &event, sizeof(event));

    if (nbytes != sizeof(event)) {
      if (nbytes == 0) {
        break;
      }

      printf("FAIL Read error on event file\n");
      IB_close(fd);
      return false;
    }

    printf("EVENT %hu\n", eventNo);

    printf("TYPE %hd\n", event.type);

    for (index = 0; index < strlen(event.desc); index++) {
      if (isspace(event.desc[index])) {
        ;
      }
    }

    const char *desc = IB_trimString(event.desc);

    if (strlen(desc) > 0) {
      printf("DESC \"%s\"\n", desc);
    }

    if (event.field_1 != 0) {
      printf("FIELD1 %hd\n", event.field_1);
    }

    if (event.field_2 != 0) {
      printf("FIELD2 %hd\n", event.field_2);
    }

    if (event.field_3 != 0) {
      printf("FIELD3 %hd\n", event.field_3);
    }

    if (event.field_4 != 0) {
      printf("FIELD4 %hd\n", event.field_4);
    }

    if (event.field_7 != 0) {
      printf("FIELD7 %d\n", event.field_7);
    }

    if (event.field_8 != 0) {
      printf("FIELD8 %d\n", event.field_8);
    }

    if (event.new_loc != 0) {
      printf("NEWLOC %hd\n", event.new_loc);
    }

    printf("END %hu\n", eventNo);
  }

  IB_close(fd);
  return true;
}

static bool
dumpLocations() {
  const int fd = OpenLocFile();

  if (fd == -1) {
    printf("FAIL Can't open location file\n");
    return false;
  }

  for (int locationNo = 9; locationNo <= LOCATION_LIMIT + 8; locationNo++) {
    dblocation_t location;
    size_t index;

    const ssize_t nbytes = IB_read(fd, &location, sizeof(location));

    if (nbytes < sizeof(location)) {
      if (nbytes == 0) {
        break;
      }
      printf("FAIL Read error on location file\n");
      IB_close(fd);
      return false;
    }

    printf("LOCATION %hu\n", locationNo);

    char *desc = IB_trimString(strtok(location.desc, "\n"));

    while (desc != NULL) {
      printf("DESC \"%s\"\n", chomp(desc));
      desc = strtok(NULL, "\n");
    }

    for (index = 0; index < DIM_OF(location.events); index++) {
      if (location.events[index] != 0) {
        printf("EVENT%zu %hd\n", index, location.events[index]);
      }
    }

    for (index = 0; index < DIM_OF(loc_flags); index++) {
      if ((location.map_flag & loc_flags[index].flag) != 0) {
        printf("FLAG %s\n", loc_flags[index].name);
      }
    }

    for (index = 0; index < DIM_OF(location.mov_tab); index++) {
      if (location.mov_tab[index] != 0) {
        printf("MOVTAB%zu %hd\n", index, location.mov_tab[index]);
      }
    }

    if (location.sys_loc != 0) {
      printf("SYSLOC %hd\n", location.sys_loc);
    }

    printf("END %hu\n", locationNo);
  }

  IB_close(fd);
  return true;
}

static bool
dumpObjects() {
  const int fd = OpenObjFile();

  if (fd == -1) {
    printf("FAIL Can't open object file\n");
    return false;
  }

  for (int objectNo = 1; objectNo <= OBJECT_LIMIT; objectNo++) {
    dbobject_t object;
    size_t index;

    const ssize_t nbytes = IB_read(fd, &object, sizeof(object));

    if (nbytes != sizeof(object)) {
      if (nbytes == 0) {
        break;
      }
      printf("FAIL Read error on event file\n");
      IB_close(fd);
      return false;
    }

    printf("OBJECT %hu\n", objectNo);

    printf("NUMBER %d\n", object.number);

    for (index = 0; index < DIM_OF(ob_flags); index++) {
      if ((object.flags & ob_flags[index].flag) != 0) {
        printf("FLAG %s\n", ob_flags[index].name);
      }
    }

    //

    const char *str = IB_trimString(object.name);

    if (strlen(str) > 0) {
      printf("NAME \"%s\"\n", str);
    }

    str = IB_trimString(object.desc);

    if (strlen(str) > 0) {
      printf("DESC \"%s\"\n", str);
    }

    str = IB_trimString(object.scan);

    if (strlen(str) > 0) {
      printf("SCAN \"%s\"\n", str);
    }

    printf("SEX %c\n", object.sex);

    if (object.start_loc != 0) {
      printf("STARTLOC %hd\n", object.start_loc);
    }

    if (object.weight != 0) {
      printf("WEIGHT %hd\n", object.weight);
    }

    if (object.get_event != 0) {
      printf("GET %hd\n", object.get_event);
    }

    if (object.give_event != 0) {
      printf("GIVE %hd\n", object.give_event);
    }

    if (object.consume_event != 0) {
      printf("CONSUME %hd\n", object.consume_event);
    }

    if (object.value != 0) {
      printf("VALUE %d\n", object.value);
    }

    if (object.drop_event != 0) {
      printf("DROP %d\n", object.drop_event);
    }

    if (object.offset != 0) {
      printf("OFFSET %d\n", object.offset);
    }

    if (object.max_loc != 0) {
      printf("MAXLOC %hd\n", object.max_loc);
    }

    if (object.min_loc != 0) {
      printf("MINLOC %hd\n", object.min_loc);
    }

    if (object.attack_percent != 0) {
      printf("ATTACK %hd\n", object.attack_percent);
    }

    if ((object.flags & ofShip) != 0) {
      for (index = 0; index < DIM_OF(object.ship_guns); index++) {
        if (object.ship_guns[index].type == 0) {
          continue;
        }
        printf("GUNS%zu %hd\n", index, object.ship_guns[index].type);
      }

      if (object.hull != 0) {
        printf("HULL %hd\n", object.hull);
      }

      if (object.shield != 0) {
        printf("SHIELD %hd\n", object.shield);
      }

      if (object.engine != 0) {
        printf("ENGINE %hd\n", object.engine);
      }

      if (object.computer != 0) {
        printf("COMPUTER %hd\n", object.computer);
      }

      if (object.fuel != 0) {
        printf("FUEL %hd\n", object.fuel);
      }

      if (object.hold != 0) {
        printf("HOLD %hd\n", object.hold);
      }

      if (object.tonnage != 0) {
        printf("TONNAGE %hd\n", object.tonnage);
      }
    }

    if (object.pref_object != 0) {
      printf("OBJECT %hd\n", object.pref_object);
    }

    if (object.move_counter != 0) {
      printf("MOVE %hd\n", object.move_counter);
    }

    printf("END %hu\n", objectNo);
  }

  IB_close(fd);
  return true;
}

void
upload() {
}

static const char *
chomp(const char *string) {
  static char buffer[DESC_SIZE];

  strcpy(buffer, string);
  size_t len = strlen(buffer);

  while (len-- > 0 && isspace(buffer[len])) {
    buffer[len] = '\0';
  }

  return buffer;
}
