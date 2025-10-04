/******************************************************************************

  Copyright (c) 1987-1997 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: locations.cc,v 1.2 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <ctype.h>
#include <fcntl.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <unistd.h>

#include <dblocation.hh>
#include <workbench.h>

#include "fedwb.hh"

extern bool needsCheck;

static void _ChangeEvent(dblocation_t *);
static void _ChangeFlags(dblocation_t *);
static void _ChangeLocs(int);
static void _ChangeMvt(dblocation_t *);
static void _ListFlags(const dblocation_t *);
static void _ListLocations(int);
static void _LocDisplay(dblocation_t *);
static void _LocEd(int);
static void _LocWrite(int);
static void _NewLoc(int, int);

/*-----------------------------------------------------------------------------

  L o c a t i o n s

  Finds out which of the location tools the player wishes to use.

-----------------------------------------------------------------------------*/

void
Locations() {
  const int fd = OpenLocFile();

  if (fd < 0) {
    return;
  }

  for (;;) {
    switch (doMenu(mnLocationMenu, 4)) {
      case 1:
        _LocWrite(fd);
        needsCheck = true;
        break;

      case 2:
        _LocEd(fd);
        needsCheck = true;
        break;

      case 3:
        _ListLocations(fd);
        break;

      case 4:
        _ChangeLocs(fd);
        needsCheck = true;
        break;

      case 0:
        close(fd);
        return;
    }
  }
}

/*-----------------------------------------------------------------------------

  O p e n L o c F i l e

  Opens the player's location file.

-----------------------------------------------------------------------------*/

int
OpenLocFile() {
  char pathname[PATH_MAX];

  getLocationPathname(GetUserID(), pathname, sizeof(pathname));
  dbgCheck(dbgIsValidString(pathname));

  return open(pathname, O_RDWR);
}

/*-----------------------------------------------------------------------------

  _ C h a n g e E v e n t

  Allows the player to change the events for the specified location.

-----------------------------------------------------------------------------*/

static void
_ChangeEvent(dblocation_t *loc) {
  int count;
  char buffer[30];

  Output("Events:");
  sprintf(buffer, "  In: %2d   Out:  %2d\n", loc->events[0], loc->events[1]);
  Output(buffer);
  Output("Which event do you want to alter -  [I]n or [O]ut?");
  GetInput(buffer, 30, true);

  switch (toupper(buffer[0])) {
    case 'I':
      count = 0;
      break;
    case 'O':
      count = 1;
      break;
    default:
      Output("Leaving events unaltered");
      return;
  }

  loc->events[count] = PromptForInteger("New value for event?");
}

/*-----------------------------------------------------------------------------

  _ C h a n g e F l a g s

  Allows the player to toggle the flags for the specified location.

-----------------------------------------------------------------------------*/

static void
_ChangeFlags(dblocation_t *loc) {
  char buffer[30];

  Output("[U]nlit                 [S]pace\n");
  Output("[T]rading Exchange      [I]nterstellar Link\n");
  Output("[G]eneral Store         [W]eapon Shop\n");
  Output("[F] - Clothing Shop     [R]estaurant/Cafe/Bar\n");
  Output("[P]eace (no fighting)   [H]ospital\n");
  Output("[B] - Ship Repair yard  [Y] - Ship Building Yard\n");
  Output("[J] - Insurance Broker  [E]lectronics Shop\n");
  Output("[O]rbit                 [L]anding pad\n");
  Output("[D]eath (kills the player)\n");
  Output("[A] - Lockable (dropped objects recycle)\n");
  Output("[Z] - Teleport-shielded area\n");

  Output("\nFlags currently set:\n");
  _ListFlags(loc);

  Output("Which flag you wish to change? ");
  GetInput(buffer, 30, true);

  switch (toupper(buffer[0])) {
    case 'A':
      loc->map_flag ^= lfLock;
      break;
    case 'B':
      loc->map_flag ^= lfRep;
      break;
    case 'D':
      loc->map_flag ^= lfDeath;
      break;
    case 'E':
      loc->map_flag ^= lfCom;
      break;
    case 'F':
      loc->map_flag ^= lfClth;
      break;
    case 'G':
      loc->map_flag ^= lfGen;
      break;
    case 'H':
      loc->map_flag ^= lfHospital;
      break;
    case 'I':
      loc->map_flag ^= lfLink;
      if (loc->map_flag & lfLink) {
        loc->map_flag |= (lfPeace | lfSpace);
      }
      break;
    case 'J':
      loc->map_flag ^= lfIns;
      break;
    case 'L':
      loc->map_flag ^= lfLanding;
      break;
    case 'O':
      loc->map_flag ^= lfOrbit;
      if (loc->map_flag & lfOrbit) {
        loc->map_flag |= (lfPeace | lfSpace);
      }
      break;
    case 'P':
      loc->map_flag ^= lfPeace;
      break;
    case 'R':
      loc->map_flag ^= lfCafe;
      break;
    case 'S':
      loc->map_flag ^= lfSpace;
      break;
    case 'T':
      loc->map_flag ^= lfTrade;
      break;
    case 'U':
      loc->map_flag ^= lfDark;
      break;
    case 'V':
      loc->map_flag ^= lfVacuum;
      break;
    case 'W':
      loc->map_flag ^= lfWeap;
      break;
    case 'Y':
      loc->map_flag ^= lfYard;
      break;
    case 'Z':
      loc->map_flag ^= lfShield;
      break;
  }

  Output("\nFlag changed...\n");
}

/*-----------------------------------------------------------------------------

  _ C h a n g e L o c s

  Changes all movement table references to the specified loc to a new one.

-----------------------------------------------------------------------------*/

static void
_ChangeLocs(int f_num) {
  dblocation_t loc;
  int count;
  int new_number, old_number, index = 1;
  bool flag;
  char buffer[200];

  strcpy(buffer, "\nThis command will go though the movement tables of ");
  strcat(buffer, "all the locations and change all references to the ");
  strcat(buffer, "location you specify to a new number. Continue? [Y/N]");
  Output(buffer);
  GetInput(buffer, 30, true);
  if (toupper(buffer[0]) != 'Y') {
    return;
  }

  Output("What location number would you like to change? ");
  GetInput(buffer, 30, true);
  old_number = atoi(buffer);
  Output("What do you want to change it to? ");
  GetInput(buffer, 30, true);
  new_number = atoi(buffer);

  if (old_number < 9 || new_number < 9) {
    Output("You can't renumber ship locations!\n");
    return;
  }

  lseek(f_num, 0, SEEK_SET);
  while (read(f_num, &loc, sizeof(loc)) > 0) {
    flag = false;
    sprintf(buffer, "Updating location number: %d...\n", index);
    Output(buffer);
    for (count = 0; count < 13; count++) {
      if (loc.mov_tab[count] == old_number) {
        loc.mov_tab[count] = new_number;
        flag = true;
      }
    }
    if (flag) { /* need to write loc back to disk */
      lseek(f_num, (long) (sizeof(loc) * (index - 1)), SEEK_SET);
      write(f_num, &loc, sizeof(loc));
      lseek(f_num, 0, SEEK_CUR);
    }
    index++;
  }
}

/*-----------------------------------------------------------------------------

  _ C h a n g e M v t

  Allows the player to alter the movement table for the specified location.

-----------------------------------------------------------------------------*/

static void
_ChangeMvt(dblocation_t *loc) {
  static const char a[] = " /------\\";
  static const char b[] = "         ";
  static const char c[] = "|        ";
  static const char d[] = " \\------/";

  int count;
  char buffer[100], value[30], temp[30];

  buffer[0] = '\0';

  if (loc->mov_tab[mvNW] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  if (loc->mov_tab[mvNorth] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  if (loc->mov_tab[mvNE] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  strcat(buffer, "\n");
  Output(buffer);

  /* line 2/3/4 */
  for (count = 0; count < 3; count++) {
    buffer[0] = '\0';

    if (loc->mov_tab[mvNW] == 0) {
      strcat(buffer, b);
    } else {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvNW]);
      strcat(buffer, temp);
    }

    if (loc->mov_tab[mvNorth] != 0) {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvNorth]);
      strcat(buffer, temp);
    }

    if ((loc->mov_tab[mvNW] != 0) && (loc->mov_tab[mvNorth] == 0)) {
      strcat(buffer, c);
    }

    if ((loc->mov_tab[mvNW] == 0) && (loc->mov_tab[mvNorth] == 0)) {
      strcat(buffer, b);
    }

    if (loc->mov_tab[mvNE] != 0) {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvNE]);
      strcat(buffer, temp);
    }

    if ((loc->mov_tab[mvNorth] != 0) && (loc->mov_tab[mvNE] == 0)) {
      strcat(buffer, c);
    }

    if ((loc->mov_tab[mvNorth] == 0) && (loc->mov_tab[mvNE] == 0)) {
      strcat(buffer, b);
    }

    if (loc->mov_tab[mvNE] != 0) {
      strcat(buffer, c);
    } else {
      strcat(buffer, b);
    }

    strcat(buffer, "\n");
    Output(buffer);
  }

  /* line 5 */

  buffer[0] = '\0';

  if (loc->mov_tab[mvNW] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  if (loc->mov_tab[mvNorth] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  if (loc->mov_tab[mvNE] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  strcat(buffer, "\n");
  Output(buffer);

  /* line 6 */
  buffer[0] = '\0';

  if (loc->mov_tab[mvWest] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  strcat(buffer, a);

  if (loc->mov_tab[mvEast] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  strcat(buffer, "\n");
  Output(buffer);

  /* line 7/8/9 */
  for (count = 0; count < 3; count++) {
    buffer[0] = '\0';

    if (loc->mov_tab[mvWest] == 0) {
      strcat(buffer, b);
    } else {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvWest]);
      strcat(buffer, temp);
    }

    strcat(buffer, "|  HOME  ");

    if (loc->mov_tab[mvEast] != 0) {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvEast]);
      strcat(buffer, temp);
    } else {
      strcat(buffer, c);
    }

    if (loc->mov_tab[mvEast] != 0) {
      strcat(buffer, c);
    } else {
      strcat(buffer, b);
    }

    strcat(buffer, "\n");
    Output(buffer);
  }

  /* line 10 */
  buffer[0] = '\0';

  if (loc->mov_tab[mvWest] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  strcat(buffer, d);

  if (loc->mov_tab[mvEast] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  strcat(buffer, "\n");
  Output(buffer);

  /* line 11 */
  buffer[0] = '\0';

  if (loc->mov_tab[mvSW] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  if (loc->mov_tab[mvSouth] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }

  if (loc->mov_tab[mvSE] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, a);
  }
  strcat(buffer, "\n");
  Output(buffer);

  /* line 12/13/14 */
  for (count = 0; count < 3; count++) {
    buffer[0] = '\0';
    if (loc->mov_tab[mvSW] == 0) {
      strcat(buffer, b);
    } else {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvSW]);
      strcat(buffer, temp);
    }

    if (loc->mov_tab[mvSouth] != 0) {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvSouth]);
      strcat(buffer, temp);
    }
    if ((loc->mov_tab[mvSW] != 0) && (loc->mov_tab[mvSouth] == 0)) {
      strcat(buffer, c);
    }
    if ((loc->mov_tab[mvSW] == 0) && (loc->mov_tab[mvSouth] == 0)) {
      strcat(buffer, b);
    }

    if (loc->mov_tab[mvSE] != 0) {
      sprintf(temp, "|  %4d  ", loc->mov_tab[mvSE]);
      strcat(buffer, temp);
    }
    if ((loc->mov_tab[mvSouth] != 0) && (loc->mov_tab[mvSE] == 0)) {
      strcat(buffer, c);
    }
    if ((loc->mov_tab[mvSouth] == 0) && (loc->mov_tab[mvSE] == 0)) {
      strcat(buffer, b);
    }

    if (loc->mov_tab[mvSE] != 0) {
      strcat(buffer, c);
    } else {
      strcat(buffer, b);
    }
    strcat(buffer, "\n");
    Output(buffer);
  }

  /* line 15 */
  buffer[0] = '\0';
  if (loc->mov_tab[mvSW] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  if (loc->mov_tab[mvSouth] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }

  if (loc->mov_tab[mvSE] == 0) {
    strcat(buffer, b);
  } else {
    strcat(buffer, d);
  }
  strcat(buffer, "\n");
  Output(buffer);

  sprintf(buffer, "Up  = %3d  ", loc->mov_tab[mvUp]);
  sprintf(temp, "Dn  = %3d  ", loc->mov_tab[mvDown]);
  strcat(buffer, temp);
  sprintf(temp, "In  = %3d", loc->mov_tab[mvIn]);
  strcat(buffer, temp);
  strcat(buffer, "\n");
  Output(buffer);
  sprintf(buffer, "Out = %3d  ", loc->mov_tab[mvOut]);
  sprintf(temp, "Pl  = %3d", loc->mov_tab[mvPlanet]);
  strcat(buffer, temp);
  strcat(buffer, "\n");
  Output(buffer);

  Output("\nWhich direction do you want to change? ");
  GetInput(buffer, 30, true);
  Output("\nWhat is it's new value? ");
  GetInput(value, 30, true);

  switch (tolower(buffer[0])) {
    case 'n':
      if (buffer[1] == '\0') {
        loc->mov_tab[mvNorth] = atoi(value);
      }
      if (buffer[1] == 'e') {
        loc->mov_tab[mvNE] = atoi(value);
      }
      if (buffer[1] == 'w') {
        loc->mov_tab[mvNW] = atoi(value);
      }
      break;

    case 'e':
      loc->mov_tab[mvEast] = atoi(value);
      break;

    case 's':
      if (buffer[1] == '\0') {
        loc->mov_tab[mvSouth] = atoi(value);
      }
      if (buffer[1] == 'e') {
        loc->mov_tab[mvSE] = atoi(value);
      }
      if (buffer[1] == 'w') {
        loc->mov_tab[mvSW] = atoi(value);
      }
      break;

    case 'w':
      loc->mov_tab[mvWest] = atoi(value);
      break;

    case 'u':
      loc->mov_tab[mvUp] = atoi(value);
      break;

    case 'd':
      loc->mov_tab[mvDown] = atoi(value);
      break;

    case 'i':
      loc->mov_tab[mvIn] = atoi(value);
      break;

    case 'o':
      loc->mov_tab[mvOut] = atoi(value);
      break;

    case 'p':
      loc->mov_tab[mvPlanet] = atoi(value);
      break;

    default:
      Output("\nNo such direction!!!!\n");
  }
}

/*-----------------------------------------------------------------------------

  _ L i s t F l a g s

-----------------------------------------------------------------------------*/

static void
_ListFlags(const dblocation_t *loc) {
  if (loc->map_flag & lfCafe) {
    Output(" Cafe\n");
  }

  if (loc->map_flag & lfClth) {
    Output(" Clothes Shop\n");
  }

  if (loc->map_flag & lfCom) {
    Output(" Comms Shop\n");
  }

  if (loc->map_flag & lfDark) {
    Output(" Unlit\n");
  }

  if (loc->map_flag & lfDeath) {
    Output(" Death\n");
  }

  if (loc->map_flag & lfGen) {
    Output(" General Store\n");
  }

  if (loc->map_flag & lfHospital) {
    Output(" Hospital\n");
  }

  if (loc->map_flag & lfIns) {
    Output(" Insurance Broker\n");
  }

  if (loc->map_flag & lfLanding) {
    Output(" Landing pad\n");
  }

  if (loc->map_flag & lfLink) {
    Output(" Interstellar Link\n");
  }

  if (loc->map_flag & lfLock) {
    Output(" Lockable Area\n");
  }

  if (loc->map_flag & lfOrbit) {
    Output(" Planetary orbit\n");
  }

  if (loc->map_flag & lfPeace) {
    Output(" Peace\n");
  }

  if (loc->map_flag & lfRep) {
    Output(" Ship Repairs\n");
  }

  if (loc->map_flag & lfShield) {
    Output(" Teleport-shielded Area\n");
  }

  if (loc->map_flag & lfSpace) {
    Output(" Space\n");
  }

  if (loc->map_flag & lfTrade) {
    Output(" Trading Exchange\n");
  }

  if (loc->map_flag & lfVacuum) {
    Output(" Vacuum\n");
  }

  if (loc->map_flag & lfWeap) {
    Output(" Weapon Shop\n");
  }

  if (loc->map_flag & lfYard) {
    Output(" Shipyard\n");
  }
}

/*-----------------------------------------------------------------------------

  _ L i s t L o c a t i o n s

  Lists out the player's location file, suitably formatted, onto the screen.

-----------------------------------------------------------------------------*/

static void
_ListLocations(int f_num) {
  char buffer[60];

  Output("\nLocations lister\n");
  Output("Start location? ");
  GetInput(buffer, 30, true);
  int start = atoi(buffer);
  Output("\nStop location? ");
  GetInput(buffer, 30, true);
  int stop = atoi(buffer);

  if (start < 9) {
    start = 9;
  }

  dblocation_t loc;
  int counter = start;

  lseek(f_num, sizeof(loc) * (start - 1), SEEK_SET);

  while (read(f_num, &loc, sizeof(loc)) > 0) {
    sprintf(buffer, "\nLocation number: %d\n", counter++);
    Output(buffer);
    _LocDisplay(&loc);
    Output("\n-------------------------------------\n");

    if (counter > stop) {
      break;
    }
  }
}

/*-----------------------------------------------------------------------------

  _ L o c D i s p l a y

  Displays the specified location.

-----------------------------------------------------------------------------*/

static void
_LocDisplay(dblocation_t *loc) {
  int count;
  char buffer[80], temp[50];

  Output(loc->desc);
  Output("\n");

  Output("Movement table:\n    n   ne    e   se    s   sw    w   nw   up    d   in  out plan  sys\n");
  buffer[0] = '\0';
  for (count = 0; count < 13; count++) {
    sprintf(temp, "%5d", loc->mov_tab[count]);
    strcat(buffer, temp);
  }
  sprintf(temp, "%5d\n", loc->sys_loc);
  strcat(buffer, temp);
  Output(buffer);

  Output("Events:");
  sprintf(buffer, "  In: %2d   Out:  %2d\n", loc->events[0], loc->events[1]);
  Output(buffer);

  Output("Flags:\n");
  _ListFlags(loc);
}

/*-----------------------------------------------------------------------------

  _ L o c E d

  Allows the player to edit his/her location database.

-----------------------------------------------------------------------------*/

static void
_LocEd(int f_num) {
  dblocation_t loc;
  char buffer[50];
  int loc_no;
  size_t loc_size;

  Output("\nLocation editor\n");
  loc_size = sizeof(dblocation_t);

  for (;;) {
    Output("Location to edit? [Q = quit] ");
    GetInput(buffer, 30, true);
    if (toupper(buffer[0]) == 'Q') {
      return;
    }
    if ((loc_no = atoi(buffer)) < 1) {
      Output("\nThe location number must be greater than zero!\n");
      continue;
    } else if (loc_no < 9) {
      Output("\nYou can't edit that location!\n");
      continue;
    }
    if ((lseek(f_num, 0, SEEK_END) / loc_size) < (unsigned) loc_no) {
      Output("\nYou haven't got that many locations written!\n");
      continue;
    }
    lseek(f_num, (long) (loc_size * (loc_no - 1)), SEEK_SET);
    read(f_num, &loc, loc_size);
    sprintf(buffer, "\nLocation number: %d\n", loc_no);
    Output(buffer);
    _LocDisplay(&loc);

    Output("\nDo you want to edit this location? ");
    GetInput(buffer, 30, true);
    if (toupper(buffer[0]) == 'N') {
      continue;
    }

    for (;;) {
      Output("\nWhat would you like to change:\n");
      Output("  [C]lear description & replace with 'xxx'\n");
      Output("  [D]escription    [E]vents\n");
      Output("  [F]lags          [M]ovement table\n");
      Output("  [S]ystem message [R]e-display location\n");
      Output("  [Q]uit\n");
      GetInput(buffer, 30, true);
      switch (toupper(buffer[0])) {
        case 'C':
          strcpy(loc.desc, "xxx\n");
          Output("\nDescription replaced with 'xxx'\n");
          break;

        case 'D':
          sprintf(buffer, "\nTotal characters may not exceed %d!\n", DESC_SIZE - 5);
          Output("Type '*h' at the start of a line for help.\n");
          Output(buffer);
          editor(loc.desc, DESC_SIZE - 5);
          break;

        case 'E':
          _ChangeEvent(&loc);
          break;

        case 'F':
          _ChangeFlags(&loc);
          break;

        case 'M':
          _ChangeMvt(&loc);
          break;

        case 'Q':
          Output("\nWrite altered location back to disk? ");
          GetInput(buffer, 30, true);
          if (toupper(buffer[0]) == 'Y') {
            lseek(f_num, (long) (loc_size * (loc_no - 1)), SEEK_SET);
            write(f_num, &loc, loc_size);
            Output("\nLocation written to disk...\n");
          } else {
            Output("\nIgnoring alterations...\n");
          }
          return;

        case 'R':
          _LocDisplay(&loc);
          break;

        case 'S':
          Output("\nNew value for System message?\n");
          GetInput(buffer, 30, true);
          loc.sys_loc = atoi(buffer);
          break;
      }
    }
  }
}

/*-----------------------------------------------------------------------------

  _ L o c W r i t e

  Writes new locations into the the player's location file.

-----------------------------------------------------------------------------*/

static void
_LocWrite(int fd) {
  int next_loc = (lseek(fd, 0, SEEK_END) / sizeof(dblocation_t)) + 1;

  if (next_loc > 8 + 120) {
    Output("\nToo many locations!\n");
    return;
  }

  Output("Location writer\n");
  _NewLoc(next_loc, fd);
}

/*-----------------------------------------------------------------------------

  _ N e w L o c

  Gets all the information for a new location from the player.

-----------------------------------------------------------------------------*/

static void
_NewLoc(int index, int f_num) {
  dblocation_t loc;
  memset(&loc, '\0', sizeof(loc));

  char buffer[DESC_SIZE];
  int count;
  char answer = 'Y';
  char temp[80], s_temp[16];

  while (answer == 'Y') {
    sprintf(buffer, "Location number %d\n", index);
    Output(buffer);
    Output("Please enter description ");
    sprintf(buffer, "(No more than %d characters):\n", DESC_SIZE - 5);
    Output(buffer);
    Output("Type '*h' at the start of a line for help.\n");
    buffer[0] = '\0';
    editor(buffer, sizeof(buffer) - 5);
    strcpy(loc.desc, buffer);

    Output("\nMovement table:");
    for (count = 0; count < 13; count++) {
      switch (count) {
        case 0:
          Output("  n:  ");
          break;
        case 1:
          Output("\n  ne: ");
          break;
        case 2:
          Output("\n  e:  ");
          break;
        case 3:
          Output("\n  se: ");
          break;
        case 4:
          Output("\n  s:  ");
          break;
        case 5:
          Output("\n  sw: ");
          break;
        case 6:
          Output("\n  w:  ");
          break;
        case 7:
          Output("\n  nw: ");
          break;
        case 8:
          Output("\n  up: ");
          break;
        case 9:
          Output("\n  d:  ");
          break;
        case 10:
          Output("\n  in: ");
          break;
        case 11:
          Output("\n  out:");
          break;
        case 12:
          Output("\n  planet:  ");
          break;
      }
      GetInput(s_temp, 16, true);
      loc.mov_tab[count] = atoi(s_temp);
    }
    Output("\n  Non-movement message: ");
    GetInput(s_temp, 16, true);
    loc.sys_loc = atoi(s_temp);

    for (count = 0; count < EVENT_SIZE; count++) {
      loc.events[count] = 0;
    }

    loc.map_flag = 0;
    Output("\nDo you want to set any flags for the location?\n");
    GetInput(s_temp, 16, true);
    if (toupper(s_temp[0]) == 'Y') {
      strcpy(buffer, "Please identify the flags you wish to set by ");
      strcat(buffer, "entering the identifying letters one after the ");
      strcat(buffer, "other and ending with <RETURN>.\nValid flags are:\n");
      Output(buffer);

      Output("[U]nlit                 [S]pace\n");
      Output("[T]rading Exchange      [I]nterstellar Link\n");
      Output("[G]eneral Store         [W]eapon Shop\n");
      Output("[F] - Clothing Shop     [R]estaurant/Cafe/Bar\n");
      Output("[P]eace (no fighting)   [H]ospital\n");
      Output("[B] - Ship Repair yard  [Y] - Ship Building Yard\n");
      Output("[J] - Insurance Broker  [E]lectronics Shop\n");
      Output("[O]rbit                 [L]anding pad\n");
      Output("[D]eath (kills the player)\n");
      Output("[A] - Lockable (dropped objects recycle)\n");
      Output("[Z] - Teleport-shielded area\n");

      GetInput(temp, 80, true);

      count = 0;
      while (temp[count]) {
        switch (toupper(temp[count++])) {
          case 'A':
            loc.map_flag |= lfLock;
            break;
          case 'B':
            loc.map_flag |= lfRep;
            break;
          case 'D':
            loc.map_flag |= lfDeath;
            break;
          case 'E':
            loc.map_flag |= lfCom;
            break;
          case 'F':
            loc.map_flag |= lfClth;
            break;
          case 'G':
            loc.map_flag |= lfGen;
            break;
          case 'H':
            loc.map_flag |= lfHospital;
            break;
          case 'I':
            loc.map_flag |= (lfLink | lfPeace | lfSpace);
            break;
          case 'J':
            loc.map_flag |= lfIns;
            break;
          case 'L':
            loc.map_flag |= lfLanding;
            break;
          case 'O':
            loc.map_flag |= (lfOrbit | lfPeace | lfSpace);
            break;
          case 'P':
            loc.map_flag |= lfPeace;
            break;
          case 'R':
            loc.map_flag |= lfCafe;
            break;
          case 'S':
            loc.map_flag |= lfSpace;
            break;
          case 'T':
            loc.map_flag |= lfTrade;
            break;
          case 'U':
            loc.map_flag |= lfDark;
            break;
          case 'V':
            loc.map_flag |= lfVacuum;
            break;
          case 'W':
            loc.map_flag |= lfWeap;
            break;
          case 'Y':
            loc.map_flag |= lfYard;
            break;
          case 'Z':
            loc.map_flag |= lfShield;
            break;
        }
      }
    }

    /* display location just created */
    sprintf(buffer, "\nLocation number: %d\n", index);
    Output(buffer);
    _LocDisplay(&loc);

    Output("\n\n   Happy? ");
    GetInput(s_temp, 16, true);
    if (toupper(s_temp[0]) == 'Y') {
      Output("\nSaving location...");
      lseek(f_num, 0, SEEK_END);
      write(f_num, &loc, sizeof(loc));
      Output("Location saved.\n");
      index++;
    } else {
      Output("\nIgnoring this location...\n");
    }

    Output("  Another one? ");
    GetInput(s_temp, 16, true);
    answer = toupper(s_temp[0]);
    Output("-----------------\n\n");
  }
}
