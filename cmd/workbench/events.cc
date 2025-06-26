/******************************************************************************

  Copyright (c) 1987-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: events.cc,v 1.2 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <algorithm>

#include <ctype.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <ibgames/goodies.h>
#include <dbevent.hh>
#include <workbench.h>

#include "fedwb.hh"

extern bool needsCheck;

static void _EvDisplay(dbevent_t *, int);
static void _EvEdit(int);
static void _EvList(int);
static void _EvWrite(int);
static short *_GetAttributeChange(dbevent_t *);

/*-----------------------------------------------------------------------------

  E v e n t s

  Finds out which of the event tools the player wishes to use.

-----------------------------------------------------------------------------*/

void
Events() {
  const int fd = OpenEvFile();

  if (fd < 0) {
    return;
  }

  for (;;) {
    switch (doMenu(mnEventMenu, 3)) {
      case 1:
        _EvWrite(fd);
        needsCheck = true;
        break;

      case 2:
        _EvEdit(fd);
        needsCheck = true;
        break;

      case 3:
        _EvList(fd);
        break;

      case 0:
        close(fd);
        return;
    }
  }
}

/*-----------------------------------------------------------------------------

  O p e n E v F i l e

  Opens the player's events file.

-----------------------------------------------------------------------------*/

int
OpenEvFile() {
  char pathname[PATH_MAX];

  getEventPathname(GetUserID(), pathname, sizeof(pathname));
  dbgCheck(dbgIsValidString(pathname));

  return open(pathname, O_RDWR);
}

/*-----------------------------------------------------------------------------

  _ E v D i s p l a y

  Displays the selected event for the user's inspection.

-----------------------------------------------------------------------------*/

static void
_EvDisplay(dbevent_t *event, int index) {
  char buffer[120];

  Output("---------------------------------\n");
  Output(FormatText(mn92, index));
  buffer[0] = '\0';

  switch (event->type) {
    case 1:
      if (event->field_1 != 0) {
        Output(FormatText(mn93, event->field_1));
      }

      if (event->field_2 != 0) {
        Output(FormatText(mn94, event->field_2));
      }

      if (event->field_3 != 0) {
        Output(FormatText(mn95, event->field_3));
      }

      if (event->field_4 != 0) {
        Output(FormatText(mn95, event->field_4));
      }

      break;

    case 3:
      if (event->field_1 != 0) {
        Output(FormatText(mn97, event->field_1, event->field_7));
      }

      if (event->field_2 != 0) {
        Output(FormatText(mn98, event->field_2, event->field_7));
      }

      if (event->field_3 != 0) {
        Output(FormatText(mn99, event->field_3, event->field_7));
      }

      if (event->field_4 != 0) {
        Output(FormatText(mn100, event->field_4, event->field_7));
      }

      break;

    case 8:
      Output(mn101);
      break;

    case 9:
      if (event->field_8 >= 0) {
        Output(FormatText(mn102, event->field_8));
      } else {
        Output(FormatText(mn103, event->field_8));
      }

      Output(mn104);

      if (event->field_1 != 0) {
        Output(FormatText(mn105, event->field_1));
      }

      if (event->field_2 != 0) {
        Output(FormatText(mn106, event->field_2));
      }

      if (event->field_3 != 0) {
        Output(FormatText(mn107, event->field_3));
      }

      if (event->field_4 != 0) {
        Output(FormatText(mn108, event->field_4));
      }

      break;
  }

  Output(mn109);
  Output(event->desc);

  Output(FormatText(mn110, event->new_loc));
}

/*-----------------------------------------------------------------------------

  _ E v E d i t

  Allows the player to edit existing events int the file.

-----------------------------------------------------------------------------*/

static void
_EvEdit(int fd) {
  dbevent_t this_event;
  short *event_field;
  char buffer[200];

  const size_t size = std::min((size_t) IB_fileSize(fd) / sizeof(dbevent_t), EVENT_LIMIT);
  const int event_no = PromptForInteger("Which event do you want to edit? ");

  if (event_no < 1 || static_cast<size_t>(event_no) > size) {
    Output(mn112);
    return;
  }

  lseek(fd, (event_no - 1) * sizeof(dbevent_t), SEEK_SET);
  read(fd, &this_event, sizeof(dbevent_t));

  _EvDisplay(&this_event, event_no);
  Output("\nIs this the one you wanted to edit? [Y/N] ");
  GetInput(buffer, 25, true);
  if (toupper(buffer[0]) != 'Y') {
    Output("\nOK - returning to edit tools...\n");
    return;
  }

  for (;;) {
    Output(mn115);
    GetInput(buffer, 25, true);

    if (toupper(buffer[0]) == 'Q' || buffer[0] == '5') {
      break;
    }

    switch (toupper(buffer[0])) {
      case '1':
      case 'C':
        event_field = _GetAttributeChange(&this_event);
        if (this_event.type == 3) {
          *event_field = PromptForInteger("\nWhat is the new value to test for? ");
          this_event.field_7 = PromptForInteger("\nWhat is the amount you want to change it by? ");
        } else {
          *event_field = PromptForInteger("\nWhat is the amount you want to change it by? ");
        }
        break;

      case '2':
      case 'D':
        strcpy(buffer, this_event.desc);
        Output("\n*l to list existing text for change - 159 chars!\n");
        if (editor(buffer, sizeof(this_event.desc) - 1)) {
          strcpy(this_event.desc, buffer);
        }
        break;

      case '3':
      case 'N':
        this_event.new_loc = PromptForInteger("\nWhat is the number of the new location?\n");
        break;

      case '4':
      case 'O':
        Output("\nWhat is the number of the object?\n");
        this_event.field_8 = PromptForInteger("\n[-ve if you want to check object is absent]\n");
        break;

      case '6':
      case 'R':
        _EvDisplay(&this_event, event_no);
        break;

      case '7':
      case 'T':
        Output(mn121);
        GetInput(buffer, 25, true);

        switch (toupper(buffer[0])) {
          case 'C':
            this_event.type = 1;
            break;
          case 'T':
            this_event.type = 3;
            break;
          case 'S':
            this_event.type = 8;
            break;
          case 'O':
            this_event.type = 9;
            break;
          default:
            Output(MessageText(mn122));
            break;
        }
        break;
    }
  }

  Output("\nWrite the changes back to disk? [Y/N] ");
  GetInput(buffer, 25, true);

  if (toupper(buffer[0]) == 'Y') {
    lseek(fd, (long) ((event_no - 1) * sizeof(dbevent_t)), SEEK_SET);

    if (write(fd, &this_event, sizeof(dbevent_t)) > 0) {
      Output("\nChanges written out to disk...\n");
    } else {
      Output("\nUnable to write changes to disk!\n");
    }
  } else {
    Output("\nIgnoring changes...\n");
  }
}

/*-----------------------------------------------------------------------------

  _ E v L i s t

  Allows the player to list the events in his/her datafile.

-----------------------------------------------------------------------------*/

static void
_EvList(int fd) {
  dbevent_t event;
  int index = 1;

  lseek(fd, 0, SEEK_SET);

  while (read(fd, &event, sizeof(event)) == sizeof(event)) {
    _EvDisplay(&event, index++);
  }
}

/*-----------------------------------------------------------------------------

  _ E v W r i t e

  Allows the player to write new events into his/her datafile.

-----------------------------------------------------------------------------*/

static void
_EvWrite(int fd) {
  dbevent_t event;
  memset(&event, '\0', sizeof(event));

  int count;
  int number;
  short *attribute;
  char buffer[60], answer[25], ok = 'y', *block;

  int size = lseek(fd, 0, SEEK_END);
  int index = (size / sizeof(dbevent_t)) + 1;

  if (index > EVENT_LIMIT) {
    Output("\nToo many events!\n");
    return;
  }

  size = sizeof(dbevent_t);

  Output("\nEvent record writer\n");
  while (ok == 'y') {
    block = (char *) &event;
    for (count = 0; count < size; count++, *(block++) = '\0')
      ;
    Output("\n---------------------------------------------------------\n");
    sprintf(buffer, "Event number %d\n", index);
    Output(buffer);
    Output("The following types of events are currently available:\n");
    Output("   [C]hange a persona attribute\n");
    Output("   [T]est a persona attribute for a minimum value\n");
    Output("   [O]bject carried test\n");
    Output("Event type: ");
    GetInput(answer, sizeof(answer), true);
    switch (toupper(answer[0])) {
      case 'C':
        event.type = 1;
        break;
      case 'T':
        event.type = 3;
        break;
      case 'S':
        event.type = 8;
        break;
      case 'O':
        event.type = 9;
        break;
      default:
        event.type = 0;
        break;
    }
    switch (event.type) {
      case 0:
        Output("\nNot a legitimate event type!\n");
        return;

      case 1:
        Output("\nWhich attribute do you wish to change? ");
        attribute = _GetAttributeChange(&event);
        Output("\nHow much do you want to alter it by? ");
        *attribute = PromptForInteger("[a negative amount will reduce the attribute] ");
        break;

      case 3:
        Output("\nWhich attribute do you wish to test? ");
        attribute = _GetAttributeChange(&event);
        *attribute = PromptForInteger("\nWhat is the minimum value for the attribute? ");
        Output("\nBy how much do you want to change the ");
        Output("attribute if the player fails the test? ");
        event.field_7 = PromptForInteger("[a negative amount will reduce the attribute] ");
        break;

      case 9:
        Output("\nWhat is the number of the object you wish to ");
        Output("check for? A positive number indicates that ");
        Output("the player SHOULD be carrying the object, a ");
        Output("negative value indicates that the player ");
        event.field_8 = PromptForInteger("SHOULD NOT be carrying the object. ");
        Output("\nWhich attribute do you wish to change if the test fails? ");
        attribute = _GetAttributeChange(&event);
        Output("\nHow much do you want to alter it by? ");
        *attribute = PromptForInteger("[a negative amount will reduce the attribute] ");
        break;
    }

    Output("\nText to go to player:[159 characters]\n");
    if (!editor(event.desc, sizeof(event.desc) - 1)) {
      event.desc[0] = '\0';
    }

    event.new_loc = PromptForInteger("\nNew location to move to: [0 = no move] ");

    _EvDisplay(&event, index);

    Output("\nHappy? ");
    GetInput(answer, sizeof(answer), true);
    if (answer[0] == 'y') {
      number = write(fd, &event, sizeof(dbevent_t));

      if (number <= 0) {
        Output("\nUnable to write to file!\n");
        return;
      } else {
        Output("\nRecord written to file...\n");
        index++;
      }
    } else {
      Output("\nIgnoring input...\n");
    }

    Output("Another event to enter? ");
    GetInput(answer, sizeof(answer), true);
    ok = tolower(answer[0]);
  }
}

/*-----------------------------------------------------------------------------

  _ G e t A t t r i b u t e C h a n g e

  Finds out which attribute the player wishes to change or test.

-----------------------------------------------------------------------------*/

static short *
_GetAttributeChange(dbevent_t *event) {
  for (;;) {
    char buffer[4];

    Output("You can change the following attributes:\n");
    Output("  [STA] - Stamina\n");
    Output("  [STR] - Strength\n");
    Output("  [DEX] - Dexterity\n");
    Output("  [INT] - Intelligence\n");
    Output("Which would you like to change? ");
    GetInput(buffer, sizeof(buffer), true);
    buffer[3] = '\0';

    if (strcasecmp(buffer, "STR") == 0) {
      return &event->field_1;
    }

    if (strcasecmp(buffer, "STA") == 0) {
      return &event->field_2;
    }

    if (strcasecmp(buffer, "INT") == 0) {
      return &event->field_3;
    }

    if (strcasecmp(buffer, "DEX") == 0) {
      return &event->field_4;
    }

    Output("\nSTA, STR, DEX or INT. Please try again\n");
  }
}
