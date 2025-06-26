/******************************************************************************

  Copyright (c) 1987-1997 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: objects.cc,v 1.2 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <ctype.h>
#include <fcntl.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <dbobject.hh>
#include <workbench.h>

#include "fedwb.hh"

extern bool needsCheck;

static void _GetMobileDetails(dbobject_t *);
static void _ObDesEdit(dbobject_t *);
static void _ObEvEdit(dbobject_t *);
static void _ObFlagsEdit(dbobject_t *);
static void _ObjDisplay(dbobject_t *);
static void _ObjEdit(int);
static void _ObjList(int);
static void _ObjWrite(int);
static void _ObLocEdit(dbobject_t *);
static void _ObMobEdit(dbobject_t *);
static void _ObShipEdit(dbobject_t *);
static bool validateName(char *name);

/*-----------------------------------------------------------------------------

  O b j e c t s

  Finds out which of the object tools the player wishes to use.

-----------------------------------------------------------------------------*/

void
Objects() {
  const int fd = OpenObjFile();

  if (fd < 0) {
    return;
  }

  for (;;) {
    switch (doMenu(mnObjectMenu, 3)) {
      case 1:
        _ObjWrite(fd);
        needsCheck = true;
        break;

      case 2:
        _ObjEdit(fd);
        needsCheck = true;
        break;

      case 3:
        _ObjList(fd);
        break;

      case 0:
        close(fd);
        return;
    }
  }
}

/*-----------------------------------------------------------------------------

  O p e n O b j F i l e

  Opens the player's object file.

-----------------------------------------------------------------------------*/

int
OpenObjFile() {
  char pathname[PATH_MAX];

  getObjectPathname(GetUserID(), pathname, sizeof(pathname));
  dbgCheck(dbgIsValidString(pathname));

  return open(pathname, O_RDWR);
}

/*-----------------------------------------------------------------------------

  _ G e t M o b i l e D e t a i l s

-----------------------------------------------------------------------------*/

static void
_GetMobileDetails(dbobject_t *obj) {
  int count;
  char buffer[30];

  obj->offset = 0;
  Output("\nHighest permitted location? ");
  GetInput(buffer, 15, true);
  obj->max_loc = atoi(buffer);
  Output("\nLowest permitted location? ");
  GetInput(buffer, 15, true);
  obj->min_loc = atoi(buffer);

  if ((obj->flags & ofShip) == 0) {
    // FIX ME -- stats kept for Genesis
    Output("\nStrength? ");
    GetInput(buffer, 15, true);
    Output("\nStamina? ");
    GetInput(buffer, 15, true);
    Output("\nIntelligence? ");
    GetInput(buffer, 15, true);
    Output("\nDexterity? ");
    GetInput(buffer, 15, true);

    Output("\nPreferred object? ");
    GetInput(buffer, 15, true);
    if ((obj->pref_object = atoi(buffer)) < -1) {
      obj->pref_object = 0;
    }
  } else {
#if 0
      obj->max_str = obj->cur_str = obj->max_sta = obj->cur_sta = 99;
      obj->max_int = obj->cur_int = obj->max_dex = obj->cur_dex = 99;
#endif
    obj->pref_object = 0;

    // for (count = 0; count < 5; obj->load[count++].goods = -1);
    Output("\nAny guns? ");
    GetInput(buffer, 15, true);
    for (count = 0; count < 4; obj->ship_guns[count++].type = 0)
      ;
    if (tolower(buffer[0]) == 'y') {
      for (count = 0; count < 4; count++) {
        Output("\n1= mag gun/2 = missile rack\n3 = laser/4 = twin-laser\n");
        Output("Weapon type number? [0 = no more weapons] ");
        GetInput(buffer, 15, true);
        switch (atoi(buffer)) {
          case 1:
            obj->ship_guns[count].type = 1;
            strcpy(obj->ship_guns[count].name, "Mag Gun");
            obj->ship_guns[count].damage = 2;
            obj->ship_guns[count].power = 2;
            break;

          case 2:
            obj->ship_guns[count].type = 2;
            strcpy(obj->ship_guns[count].name, "Missile");
            obj->ship_guns[count].damage = 4;
            obj->ship_guns[count].power = 0;
            break;

          case 3:
            obj->ship_guns[count].type = 3;
            strcpy(obj->ship_guns[count].name, "Laser");
            obj->ship_guns[count].damage = 5;
            obj->ship_guns[count].power = 15;
            break;

          case 4:
            obj->ship_guns[count].type = 4;
            strcpy(obj->ship_guns[count].name, "Twin Laser");
            obj->ship_guns[count].damage = 10;
            obj->ship_guns[count].power = 30;
            break;

          default:
            for (; count < 4; count++) {
              obj->ship_guns[count].type = 0;
            }
            break;
        }
      }
    }

    obj->tonnage = 1000;
    Output("\nHull strength? ");
    GetInput(buffer, 15, true);
    if ((obj->hull = atoi(buffer)) < 10) {
      obj->hull = 10;
    }
    Output("\nShield strength? ");
    GetInput(buffer, 15, true);
    if ((obj->shield = atoi(buffer)) < 0) {
      obj->shield = 0;
    }
    obj->engine = obj->tonnage / 10;
    Output("\nComputer level? ");
    GetInput(buffer, 15, true);
    if ((obj->computer = atoi(buffer)) < 1) {
      obj->computer = 1;
    }
    if (obj->computer > 6) {
      obj->computer = 6;
    }
    obj->fuel = 999;
    obj->hold = 999;
  }

  Output("\nPercentage attack factor? ");
  GetInput(buffer, 15, true);
  if ((obj->attack_percent = atoi(buffer)) > 100) {
    obj->attack_percent = 100;
  }

  obj->kill_event = 0;

  Output("\nCan this mobile move? ");
  GetInput(buffer, 15, true);
  if (buffer[0] == 'n') {
    obj->move_counter = obj->max_counter = -1;
  } else {
    Output("\nMovement interval? ");
    GetInput(buffer, 15, true);
    if ((obj->move_counter = obj->max_counter = atoi(buffer)) < -1) {
      obj->move_counter = obj->max_counter = -1;
    }
  }
}

/*-----------------------------------------------------------------------------

  _ O b D e s E d i t

  Allows player to edit object descs.

-----------------------------------------------------------------------------*/

static void
_ObDesEdit(dbobject_t *obj) {
  char buffer[205];

  Output("\nEdit [L]ook message or [E]xamine message? ");
  GetInput(buffer, 5, true);
  switch (toupper(buffer[0])) {
    case 'L':
      strcpy(buffer, obj->desc);
      Output("\nYou have 80 characters for the description\n");
      if (editor(buffer, 80)) {
        strcpy(obj->desc, buffer);
      } else {
        Output("\nNo change made!\n");
      }
      break;

    case 'E':
      strcpy(buffer, obj->scan);
      Output("\nYou have 200 characters for the scan description\n");
      if (editor(buffer, 200)) {
        strcpy(obj->scan, buffer);
      } else {
        Output("\nNo change made!\n");
      }
      break;

    default:
      Output("\nNo changes made\n");
      break;
  }
}

/*-----------------------------------------------------------------------------

  _ O b E v E d i t

  Edit events associated with an object.

-----------------------------------------------------------------------------*/

static void
_ObEvEdit(dbobject_t *obj) {
  char buffer[60];

  for (;;) {
    Output("\nWhich event do you want to change?\n");
    if ((obj->flags & ofAnimate) == 0) {
      Output("  [D]rop the object\n");
      Output("  [E]at/drink the substance\n");
    }
    Output("  [G]ive away the object\n");
    Output("  [P]ick up (get) the object\n");
    Output("  [R]e-display object\n");
    Output("  [Q]uit editing object events\n");
    GetInput(buffer, 5, true);
    switch (toupper(buffer[0])) {
      case 'D':
        if (obj->flags & ofAnimate) {
          break;
        }
        sprintf(buffer, "Drop - current value is %d\n", obj->drop_event);
        Output(buffer);
        Output("New value? ");
        GetInput(buffer, 20, true);
        obj->drop_event = atoi(buffer);
        break;

      case 'E':
        if (obj->flags & ofAnimate) {
          break;
        }
        sprintf(buffer, "Eat/Drink - current value is %d\n", obj->consume_event);
        Output(buffer);
        Output("New value? ");
        GetInput(buffer, 20, true);
        obj->consume_event = atoi(buffer);
        break;

      case 'G':
        sprintf(buffer, "Give - current value is %d\n", obj->give_event);
        Output(buffer);
        Output("New value? ");
        GetInput(buffer, 20, true);
        obj->give_event = atoi(buffer);
        break;

#if 0
	 case 'K':
	    if (!(obj->flags & ofAnimate))
	       break;
	    sprintf(buffer, "Kill - current value is %d\n", obj->kill_event);
	    Output(buffer);
	    Output("New value? ");
	    GetInput(buffer, 20, true);
	    obj->kill_event = atoi(buffer);
	    break;
#endif

      case 'P':
        sprintf(buffer, "Get - current value is %d\n", obj->get_event);
        Output(buffer);
        Output("New value? ");
        GetInput(buffer, 20, true);
        obj->get_event = atoi(buffer);
        break;

      case 'Q':
        return;

      case 'R':
        _ObjDisplay(obj);
        break;
    }
  }
}

/*-----------------------------------------------------------------------------

  _ O b F l a g s E d i t

  Edit flags associated with an object.

-----------------------------------------------------------------------------*/

static void
_ObFlagsEdit(dbobject_t *obj) {
  char buffer[60];

  for (;;) {
    Output("Object flags are as follows:\n");

    sprintf(buffer, "  [C] - Cleaner:            ");
    if (obj->flags & ofCleaner) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    sprintf(buffer, "  [E] - Edible:             ");
    if (obj->flags & ofEdible) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    sprintf(buffer, "  [I] - Musical instrument: ");
    if (obj->flags & ofMusic) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    sprintf(buffer, "  [L] - Drinkable:          ");
    if (obj->flags & ofLiquid) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    sprintf(buffer, "  [M] - Mobile:             ");
    if (obj->flags & ofAnimate) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    sprintf(buffer, "  [S] - Spaceship:          ");
    if (obj->flags & ofShip) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    sprintf(buffer, "  [T] - Light source:       ");
    if (obj->flags & ofLight) {
      strcat(buffer, "set\n");
    } else {
      strcat(buffer, "clear\n");
    }
    Output(buffer);

    Output("Which one do you want to change?\n");
    Output("[Q] to quit editing flags. ");
    GetInput(buffer, 5, true);

    switch (toupper(buffer[0])) {
      case 'C':
        obj->flags ^= ofCleaner;
        if (obj->flags & ofCleaner) {
          obj->flags |= ofAnimate;
          obj->flags &= ~ofShip;
        }
        break;

      case 'E':
        obj->flags ^= ofEdible;
        if ((obj->flags & ofEdible) == ofEdible) {
          obj->flags &= ~ofLiquid;
        }
        break;

      case 'I':
        obj->flags ^= ofMusic;
        break;

      case 'L':
        obj->flags ^= ofLiquid;
        if ((obj->flags & ofLiquid) == ofLiquid) {
          obj->flags &= ~ofEdible;
        }
        break;

      case 'M':
        obj->flags ^= ofAnimate;
        if ((obj->flags & ofAnimate) != ofAnimate) {
          obj->flags &= ~ofShip;
          obj->flags &= ~ofCleaner;
        }
        break;

      case 'P':
        obj->flags ^= ofNoThe;
        break;

      case 'Q':
        return;

      case 'S':
        obj->flags ^= ofShip;
        if (obj->flags & ofShip) {
          obj->flags |= ofAnimate;
          obj->flags &= ~ofCleaner;
        }
        break;

      case 'T':
        obj->flags ^= ofLight;
        break;
    }
  }
}

/*-----------------------------------------------------------------------------

  _ O b j D i s p l a y

  Displays an object's details.

-----------------------------------------------------------------------------*/

static void
_ObjDisplay(dbobject_t *obj) {
  static const char line[] = "\n------------------------------------------------\n";

  int count;
  int flag;
  char buffer[80];

  Output(line);
  sprintf(buffer, "Name: %s   Sex: %c   ID number: %d\n",
          obj->name, obj->sex, obj->number);
  Output(buffer);

  Output("Desc:\n");
  Output(obj->desc);
  Output("\nScan:\n");
  Output(obj->scan);

  Output("\nObject flags:\n");
  if (obj->flags & ofLight) {
    Output("  Light source\n");
  }
  if (obj->flags & ofMusic) {
    Output("  Musical instrument\n");
  }
  if (obj->flags & ofLiquid) {
    Output("  Drinkable\n");
  }
  if (obj->flags & ofEdible) {
    Output("  Edible\n");
  }
  if (obj->flags & ofAnimate) {
    Output("  Mobile\n");
  }
  if (obj->flags & ofShip) {
    Output("  Spaceship\n");
  }
  if (obj->flags & ofNoThe) {
    Output("  No 'the' before name\n");
  }
  if (obj->flags & ofCleaner) {
    Output("  Cleaning droid\n");
  }

  if (!(obj->flags & ofAnimate)) {
    sprintf(buffer, "\nStart loc: %d  Recycling offset: %d\n", obj->start_loc, obj->offset);
    Output(buffer);
    sprintf(buffer, "Weight: %d  Base value: %d\n", obj->weight, obj->value);
    Output(buffer);
  } else {
    sprintf(buffer, "\nReward: %4d\n", obj->value);
    Output(buffer);
  }

  Output("\nEvents:      Get    Give   Drop   Eat/Drink\n");
  sprintf(buffer, "%16d%7d%7d%10d\n", obj->get_event, obj->give_event,
          obj->drop_event, obj->consume_event);
  Output(buffer);
#if 0
   if (obj->flags & ofAnimate)
   {
      sprintf(buffer, "Kill event: %d\n", obj->kill_event);
      Output(buffer);
   }
#endif

  if (obj->flags & ofAnimate) {
    Output("Locations:  Start   High   Low\n");
    sprintf(buffer, "%16d%7d%7d\n", obj->cur_loc, obj->max_loc, obj->min_loc);
    Output(buffer);
    if (obj->flags & ofShip) {
      flag = true;
      Output("\nObject is a spaceship\n");
      sprintf(buffer, "Weaponry:  ");
      for (count = 0; count < 4; count++) {
        if (obj->ship_guns[count].type > 0) {
          if (count != 0) {
            strcat(buffer, "/");
          }
          strcat(buffer, obj->ship_guns[count].name);
          flag = false;
        }
      }
      if (flag) {
        strcat(buffer, " None");
      }
      strcat(buffer, "\n");
      Output(buffer);

      sprintf(buffer, "Hull: %d   Shields: %d   Computer: %d\n",
              obj->hull, obj->shield, obj->computer);
      Output(buffer);
    } else {
      // FIX ME -- stats kept for Genesis
      sprintf(buffer, "Str: 1   Sta: 1   Int: 1   Dex: 1\n");
      Output(buffer);
      sprintf(buffer, "Preferred object: %d\n", obj->pref_object);
      Output(buffer);
    }
    sprintf(buffer, "\nAttack likelihood %d%%\n", obj->attack_percent);
    Output(buffer);
    if (obj->move_counter < 0) {
      Output("Mobile does not move\n");
    } else {
      sprintf(buffer, "Moves every %d x 5 seconds\n", obj->max_counter);
      Output(buffer);
    }
  }
  Output(line);
}

/*-----------------------------------------------------------------------------

  _ O b j E d i t

  Allows a player to edit the attributes of objects in the object file.

-----------------------------------------------------------------------------*/

static void
_ObjEdit(int fd) {
  dbobject_t obj;
  int count, index;
  unsigned obj_size;
  int name_size, number;
  char buffer[50], temp[50];

  Output("\nObject editor\n");
  obj_size = sizeof(dbobject_t);

  for (;;) {
    Output("Object to edit? [QUIT to exit editor] ");
    GetInput(buffer, 30, true);
    name_size = strlen(buffer);
    for (count = 0; count < name_size; count++) {
      buffer[count] = tolower(buffer[count]);
    }
    if (strcmp(buffer, "quit") == 0) {
      return;
    }
    number = atoi(buffer);
    lseek(fd, 0, SEEK_SET);
    for (index = 1; index < 9999; index++) {
      if (read(fd, &obj, obj_size) == 0) {
        Output("\nUnable to find an object with that name or number!\n");
        return;
      }

      if (index == number) {
        break;
      }

      strcpy(temp, obj.name);
      name_size = strlen(temp);
      for (count = 0; count < name_size; count++) {
        temp[count] = tolower(temp[count]);
      }
      if (strcmp(temp, buffer) == 0) {
        break;
      }
    }

    _ObjDisplay(&obj);
    Output("\nDo you want to edit this object? ");
    GetInput(buffer, 30, true);
    if (toupper(buffer[0]) == 'N') {
      continue;
    }
    for (;;) {
      Output("\nWhich items do you want to alter?\n");
      Output("  [D]escriptions  [E]vents        [F]lags\n");
      Output("  [L]ocations     [M]obile items  [N]ame\n");
      Output("  [S]ex           [V]alue/Reward  [W]eight\n");
      Output("  [R]e-display this object\n");
      Output("  [I]d number     e[X]it - abandon changes\n");
      GetInput(buffer, 30, true);
      if (toupper(buffer[0]) == 'X') {
        Output("\nChanges abandoned\n");
        break;
      }
      switch (toupper(buffer[0])) {
        case 'D':
          _ObDesEdit(&obj);
          break;
        case 'E':
          _ObEvEdit(&obj);
          break;
        case 'F':
          _ObFlagsEdit(&obj);
          break;
        case 'I':
          obj.number = PromptForInteger("\nID number to use for this object/mobile?");
          if ((obj.number < 0) || (obj.number > 49)) {
            Output("\nMust be between 0 and 49...");
            obj.number = 0;
          }
          break;

        case 'L':
          _ObLocEdit(&obj);
          break;
        case 'M':
          _ObMobEdit(&obj);
          break;
        case 'N':
          Output("\nNew name? [15 characters max] ");
          GetInput(buffer, 20, true);
          buffer[15] = '\0';

          if (validateName(buffer)) {
            strcpy(obj.name, buffer);
          }

          break;

        case 'R':
          _ObjDisplay(&obj);
          break;

        case 'S':
          Output("\nSex of object? [M,F, or N] ");
          GetInput(buffer, 30, true);
          buffer[0] = tolower(buffer[0]);
          if ((buffer[0] != 'm') || (buffer[0] != 'f')) {
            obj.sex = static_cast<sex_t>('n');
          }
          obj.sex = static_cast<sex_t>(buffer[0]);
          break;

        case 'V':
          Output("\nWhat is the object's value/reward? ");
          GetInput(buffer, 15, true);
          obj.value = atoi(buffer);
          break;

        case 'W':
          Output("\nWhat is the object's weight (str units/2)? ");
          GetInput(buffer, 15, true);
          if ((obj.weight = atoi(buffer)) < 1) {
            obj.weight = 1;
          }
          break;
      }
      Output("\nMore editing on this object? [Y/N] ");
      GetInput(buffer, 10, true);
      if (toupper(buffer[0]) == 'Y') {
        continue;
      }
      Output("\nSave this altered object back to disk? [Y/N] ");
      GetInput(buffer, 10, true);
      if (toupper(buffer[0]) == 'Y') {
        lseek(fd, -sizeof(dbobject_t), SEEK_CUR);
        write(fd, &obj, sizeof(dbobject_t));

        Output("\nObject saved\n");
      } else {
        Output("\nEdit abandoned\n");
      }
      break;
    }
  }
}

/*-----------------------------------------------------------------------------

  _ O b j L i s t

  List out the contents of the object file to screen.

-----------------------------------------------------------------------------*/

static void
_ObjList(int fd) {
  dbobject_t obj;

  lseek(fd, 0, SEEK_SET);

  while (read(fd, &obj, sizeof(obj)) == sizeof(obj)) {
    _ObjDisplay(&obj);
    Output("\n");
  }
}

/*-----------------------------------------------------------------------------

  _ O b j W r i t e

  Allows the player to write a new object into his/her datafile.

-----------------------------------------------------------------------------*/

static void
_ObjWrite(int fd) {
  dbobject_t obj;
  memset(&obj, '\0', sizeof(obj));

  int count;
  int mobile;
  char answer[15], buffer[250];

  if (lseek(fd, 0, SEEK_END) / sizeof(dbobject_t) >= 14) {
    Output("\nToo many objects!\n");
    return;
  }

  answer[0] = 'y';

  while (answer[0] == 'y') {
    Output("\nEnter object name (up to 15 chars)\n");
    GetInput(buffer, 80, true);
    buffer[15] = '\0';

    if (!validateName(buffer)) {
      continue;
    }

    strcpy(obj.name, buffer);
    obj.number = -1;

    Output("\nWhich flags do you want to set for this object?\n");
    Output("Flags available are:\n");
    Output("  [C] - Object is a cleaning droid\n");
    Output("  [E] - Object is edible\n");
    Output("  [I] - Object is a musical instrument\n");
    Output("  [L] - Object is drinkable\n");
    Output("  [M] - Object is a mobile\n");
    Output("  [S] - Object is a spaceship\n");
    Output("  [T] - Object is a light source\n");
    Output("  [X] - No flags to be set\n");
    Output("Setting the 'S' flag automatically sets the 'M' flag\n");
    Output("Setting the 'C' flag automatically sets the 'M' flag\n");
    Output("Please enter the letters and end with <RETURN>\n");
    mobile = false;
    obj.flags = 0;
    GetInput(buffer, 80, true);
    count = 0;

    while (buffer[count]) {
      switch (tolower(buffer[count++])) {
        case 'c':
          obj.flags |= ofAnimate + ofCleaner;
          obj.flags &= ~ofShip;
          mobile = true;
          break;
        case 'e':
          obj.flags |= ofEdible;
          obj.flags &= ~ofLiquid;
          break;
        case 'i':
          obj.flags |= ofMusic;
          break;
        case 'l':
          obj.flags |= ofLiquid;
          obj.flags &= ~ofEdible;
          break;
        case 'm':
          obj.flags |= ofAnimate;
          mobile = true;
          break;
        case 's':
          obj.flags |= ofAnimate + ofShip;
          obj.flags &= ~ofCleaner;
          mobile = true;
          break;
        case 't':
          obj.flags |= ofLight;
          break;
      }
    }

    Output("\nEnter description for LOOK [up to 80 chars]:\n");
    buffer[0] = '\0';
    editor(buffer, 81);
    strcpy(obj.desc, buffer);
    Output("\nEnter description for EXAMINE [up to 200 chars]\n");
    buffer[0] = '\0';
    editor(buffer, 201);
    strcpy(obj.scan, buffer);

    if (!mobile) {
      obj.sex = static_cast<sex_t>('n');
    } else {
      Output("\nSex of mobile [m/f/n]: ");
      GetInput(buffer, 15, true);
      obj.sex = static_cast<sex_t>(tolower(buffer[0]));
    }

    Output("\nWhat is the object's start location number? ");
    GetInput(buffer, 15, true);
    obj.cur_loc = obj.start_loc = atoi(buffer);

    obj.weight = 0;
    obj.give_event = obj.consume_event = obj.drop_event = 0;

    if (!mobile) {
      Output("\nWhat is the object's weight (in str units/2)? ");
      GetInput(buffer, 15, true);
      if ((obj.weight = atoi(buffer)) < 1) {
        obj.weight = 1;
      }
      obj.kill_event = 0;

      Output("\nValue for GIVE event [0 for no event]: ");
      GetInput(buffer, 15, true);
      obj.give_event = atoi(buffer);

      if ((obj.flags & ofLiquid) != 0 || (obj.flags & ofEdible) != 0) {
        Output("\nValue for EAT/DRINK event: ");
        GetInput(buffer, 15, true);
        obj.consume_event = atoi(buffer);
      } else {
        obj.consume_event = 0;
      }

      Output("\nValue for DROP event: ");
      GetInput(buffer, 15, true);
      obj.drop_event = atoi(buffer);
    }

    Output("\nValue for GET event: ");
    GetInput(buffer, 15, true);
    obj.get_event = atoi(buffer);

    if (mobile) {
      Output("\nWhat is the mobile's reward value?\n");
    } else {
      Output("\nWhat is the value of the object?\n");
    }
    Output("REMEMBER - the money comes from your treasury! ");
    GetInput(buffer, 15, true);
    obj.value = atoi(buffer);

    if (mobile) {
      _GetMobileDetails(&obj);
    } else {
      Output("\nMax offset from start for recycling? ");
      GetInput(buffer, 15, true);
      obj.offset = atoi(buffer);
    }

    Output("\n");
    _ObjDisplay(&obj);
    Output("\nHappy? ");
    GetInput(buffer, 5, true);

    if (buffer[0] == 'y') {
      Output("\nWriting data to disk file...\n");
      lseek(fd, 0, SEEK_END);
      write(fd, &obj, sizeof(dbobject_t));
    } else {
      Output("\nIgnoring this record...\n");
    }

    Output("\n  Another one?");
    GetInput(answer, 5, true);
    Output("\n");
  }
}

/*-----------------------------------------------------------------------------

  _ O b L o c E d i t

  Edit object location limits.

-----------------------------------------------------------------------------*/

static void
_ObLocEdit(dbobject_t *obj) {
  char buffer[20];

  for (;;) {
    Output("\nChange [H]igh, [L]ow, [O]ffset, or\n");
    Output("[S]tart location number, or [Q]uit?\n");
    GetInput(buffer, 5, true);

    switch (toupper(buffer[0])) {
      case 'H':
        Output("\nNew value? ");
        GetInput(buffer, 20, true);
        obj->max_loc = atoi(buffer);
        break;

      case 'L':
        Output("\nNew value? ");
        GetInput(buffer, 20, true);
        obj->min_loc = atoi(buffer);
        break;

      case 'O':
        Output("\nNew value? ");
        GetInput(buffer, 20, true);
        obj->offset = atoi(buffer);
        break;

      case 'Q':
        return;

      case 'S':
        Output("\nNew value? ");
        GetInput(buffer, 20, true);
        obj->cur_loc = obj->start_loc = atoi(buffer);
        break;
    }
    Output("\nValue changed...\n");
  }
}

/*-----------------------------------------------------------------------------

  _ O b M o b E d i t

  Edit mobile attributes.

-----------------------------------------------------------------------------*/

static void
_ObMobEdit(dbobject_t *obj) {
  char buffer[20];

  for (;;) {
    Output("\nWhat do you want to change:\n");
    Output("  [A]ttack %% [M]ovement rate [P]referred object\n");
    Output("  [L]ocation limits [S]tats or [Q]uit\n");
    GetInput(buffer, 5, true);
    switch (toupper(buffer[0])) {
      case 'A':
        obj->attack_percent = PromptForInteger("\nWhat is the new attack percentage?\n");
        break;

      case 'L':
        obj->max_loc = PromptForInteger("\nHighest location allowed? ");
        obj->min_loc = PromptForInteger("\nLowest location allowed? ");
        break;

      case 'M':
        obj->move_counter = PromptForInteger("\nNew movement speed? [-1 if no movement] ");
        obj->max_counter = obj->move_counter;
        break;

      case 'P':
        obj->pref_object = PromptForInteger("\nWhat is the new preferred object? ");
        break;

      case 'Q':
        return;

      case 'S':
        if (obj->flags & ofShip) {
          _ObShipEdit(obj);
          break;
        }
        Output("\nWhich stat do you want to alter?\n");
        Output("  [A] - Stamina      [R] - Strength\n");
        Output("  [I] - Intell       [D] - Dexterity\n");
        GetInput(buffer, 5, true);
        switch (toupper(buffer[0])) {
          case 'A':
            PromptForInteger("\nWhat is the new stamina? ");
            break;
          case 'R':
            PromptForInteger("\nWhat is the new strength? ");
            break;
          case 'I':
            PromptForInteger("\nWhat is the new intelligence? ");
            break;
          case 'D':
            PromptForInteger("\nWhat is the new dexterity? ");
            break;
        }
    }
  }
}

/*-----------------------------------------------------------------------------

  _ O b S h i p E d i t

  Edit ship attributes.

-----------------------------------------------------------------------------*/

static void
_ObShipEdit(dbobject_t *obj) {
  int count;
  int temp, flag = true;
  char buffer[60];

  Output("\nWhat do you want to change:\n");
  Output("  [C]omputer   [H]ull\n");
  Output("  [S]hields    [W]eapons\n");

  GetInput(buffer, 5, true);
  switch (toupper(buffer[0])) {
    case 'C':
      temp = PromptForInteger("\nWhat is the new computer level? ");
      if (temp < 1) {
        temp = 1;
      }
      if (temp > 6) {
        temp = 6;
      }
      obj->computer = temp;
      break;

    case 'H':
      Output("\nWhat is the new hull strength? ");
      GetInput(buffer, 10, true);
      obj->hull = atoi(buffer);
      break;

    case 'S':
      Output("\nWhat is the new shield strength? ");
      GetInput(buffer, 10, true);
      obj->shield = atoi(buffer);
      break;

    case 'W':
      Output("\n[A]dd or [D]elete a weapon? ");
      GetInput(buffer, 5, true);
      if (toupper(buffer[0]) == 'D') {
        for (count = 0; count < 4; count++) {
          if (obj->ship_guns[count].type > 0) {
            sprintf(buffer, " [%d] %s\n", count, obj->ship_guns[count].name);
            Output(buffer);
            flag = false;
          }
        }
        if (flag) {
          Output("\nCraft is unarmed!\n");
          break;
        }
        temp = PromptForInteger("\nWhich one do you want to remove ([-1] = none)? ");
        if ((temp < 0) || (temp > 3)) {
          break;
        }

        for (count = temp; count < 3; count++) {
          memcpy(&(obj->ship_guns[count]),
                 &(obj->ship_guns[count + 1]),
                 sizeof(old_s_guns_t));
        }

        obj->ship_guns[3].type = 0;
        Output("\nWeapon removed...\n");
        break;
      }
      if (toupper(buffer[0]) == 'A') {
        for (count = 0; count < 4; count++) {
          if (obj->ship_guns[count].type == 0) {
            break;
          }
        }
        if (count == 4) {
          Output("\nNo space to install more weapons!\n");
          break;
        }

        Output("\nWeapons are 1=mag/2=mis/3=las/4=d-las\n");
        Output("Weapon type number? [0 is no weapon] ");
        GetInput(buffer, 15, true);
        switch (atoi(buffer)) {
          case 1:
            obj->ship_guns[count].type = 1;
            strcpy(obj->ship_guns[count].name, "Mag Gun");
            obj->ship_guns[count].damage = 2;
            obj->ship_guns[count].power = 2;
            break;

          case 2:
            obj->ship_guns[count].type = 2;
            strcpy(obj->ship_guns[count].name, "Missile");
            obj->ship_guns[count].damage = 4;
            obj->ship_guns[count].power = 0;
            break;

          case 3:
            obj->ship_guns[count].type = 3;
            strcpy(obj->ship_guns[count].name, "Laser");

            obj->ship_guns[count].damage = 5;
            obj->ship_guns[count].power = 15;
            break;

          case 4:
            obj->ship_guns[count].type = 4;
            strcpy(obj->ship_guns[count].name, "Twin Laser");
            obj->ship_guns[count].damage = 10;
            obj->ship_guns[count].power = 30;
            break;
        }
        Output("\nWeapon added...\n");
        break;
      }
      break;
  }
}

/*-----------------------------------------------------------------------------

  v a l i d a t e N a m e

-----------------------------------------------------------------------------*/

static bool
validateName(char *name) {
  if (!isalpha(name[0])) {
    Output("Names must begin with a letter.\n");
    return false;
  }

  const size_t length = strlen(name);

  if (length < 3) {
    Output("Names must be at least 3 characters long.\n");
    return false;
  }

  for (size_t index = 1; index < length; index++) {
    if (!isalnum(name[index]) && name[index] != '-') {
      Output("Names can only contain letters, numbers and hyphens.\n");
      return false;
    }

    if (isupper(name[index])) {
      name[index] = tolower(name[index]);
    }
  }

  return true;
}
