/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: workbench.cc,v 1.4 1999/04/24 12:36:00 nick Exp $

******************************************************************************/

#include <ctype.h>
#include <errno.h>
#include <stdarg.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include <ibgames/goodies.h>
#include <ibgames/log.h>
#include <workbench.h>

#include "fedwb.hh"

#define TEXTMAX (16 * 1024)

unsigned debugEnableFlags = 0;
bool needsCheck = false;
bool noExchange = false;

static bool checkOnly = false;
static account_id_t userID;

static void clearAllFiles();
static void mainMenu();
static bool parseArguments(int, char **);

/*-----------------------------------------------------------------------------

  m a i n

-----------------------------------------------------------------------------*/

int
main(int argc, char **argv) {
  IB_setLogOptions("wb", LOG_PID);

  //

  setvbuf(stdout, NULL, _IOLBF, 0);
  setvbuf(stderr, NULL, _IOLBF, 0);

  //

  if (!parseArguments(argc, argv)) {
    IB_log("usage: %s [-c] [-n] id", argv[0]);
    return EXIT_FAILURE;
  }

  dbgCheck(userID >= 100000);

  //

  if (checkOnly) {
    return (checkPlanet() ? EXIT_SUCCESS : EXIT_FAILURE);
  }

  Output(mnHello);
  mainMenu();
  Output(mnGoodbye);

  fflush(stdout);
  fflush(stderr);
  sleep(1);

  return EXIT_SUCCESS;
}

/*-----------------------------------------------------------------------------

  d o M e n u

-----------------------------------------------------------------------------*/

unsigned
doMenu(msg_id_t messageNo, unsigned maxChoice) {
  Output(messageNo);

  for (;;) {
    char input[256];

    Output(mnMenuPrompt);
    GetInput(input, sizeof(input), true);

    size_t inputPos = 0;

    while (isspace(input[inputPos])) {
      inputPos++;
    }

    bool badInput = true;
    unsigned choice = 0;

    while (isdigit(input[inputPos])) {
      badInput = false;

      if (choice > UINT_MAX / 10) {
        badInput = true;
        break;
      }

      choice = (choice * 10) + (input[inputPos] - '0');
      inputPos++;
    }

    if (badInput || choice > maxChoice) {
      continue;
    }

    while (isspace(input[inputPos])) {
      inputPos++;
    }

    if (input[inputPos] != '\0') {
      continue;
    }

    return choice;
  }
}

/*-----------------------------------------------------------------------------

  G e t I n p u t

-----------------------------------------------------------------------------*/

void
GetInput(char *buf, size_t size, bool trim) {
  fflush(stdout);

  size_t pos = 0;

  for (;;) {
    const int ch = getchar();

    if (ch == EOF) {
      fprintf(stderr, "GetInput: EOF on stdin\n");
      exit(EXIT_FAILURE);
    } else if (ch == '\n') {
      buf[pos] = '\0';
      break;
    }

    if (pos == 0 && ch == ' ' && trim) {
      continue;
    }

    buf[pos++] = ch;

    if (pos == size) {
      pos = size - 1;
    }
  }

  if (trim) {
    while (pos > 0 && buf[--pos] == ' ') {
      buf[pos] = '\0';
    }
  }

  puts(buf);
}

/*-----------------------------------------------------------------------------

  G e t U s e r I D

-----------------------------------------------------------------------------*/

account_id_t
GetUserID() {
  return userID;
}

/*-----------------------------------------------------------------------------

  m a i n M e n u

  Finds out which of the explorer tools the player wishes to use.

-----------------------------------------------------------------------------*/

static void
mainMenu() {
  if (workbenchAccess(userID) != WB_ACCESS_OK) {
    Output(mn77);
    return;
  }

  for (;;) {
    switch (doMenu(mnMainMenu, 7)) {
      case 0:
        if (needsCheck) {
          Output(mnUncheckedPreamble);
          if (!AnswerIsYes(mnUncheckedPrompt)) {
            break;
          }
        }
        return;

      case 1:
        Events();
        break;

      case 2:
        Objects();
        break;

      case 3:
        Locations();
        break;

      case 4:
        checkPlanet();
        break;

      case 5:
        clearAllFiles();
        break;

#if 0
	 case 6:
	    download();
	    break;

	 case 7:
	    upload();
	    break;
#endif
    }
  }
}

/*-----------------------------------------------------------------------------

  p a r s e A r g u m e n t s

-----------------------------------------------------------------------------*/

static bool
parseArguments(int argc, char **argv) {
  optind = 1;

  for (;;) {
    const int c = getopt(argc, argv, ":cn");

    if (c == -1) {
      break;
    }

    switch (c) {
      case 'c':
        checkOnly = true;
        break;

      case 'n':
        noExchange = true;
        break;

      case ':':
        IB_log("Option -%c requires an argument", optopt);
        return false;

      case '?':
      default:
        IB_log("Unrecognized argument: -%c", optopt);
        return false;
    }
  }

  if (optind == argc - 1) {
    userID = atoi(argv[optind]);

    if (userID < MIN_ACCOUNT_ID || userID >= MAX_ACCOUNT_ID) {
      IB_log("Bad id argument");
      return false;
    }
  } else {
    if (optind == argc) {
      IB_log("Missing id argument");
    }

    return false;
  }

  return true;
}

/*-----------------------------------------------------------------------------

  c l e a r A l l F i l e s

  Clear all the workbench files to allow the player to start again.

-----------------------------------------------------------------------------*/

static void
clearAllFiles() {
  /* The events file... */

  Output(mn484);

  if (AnswerIsYes(mnAreYouSure)) {
    if (!createWorkbenchFiles(GetUserID(), WB_CREATE_EVT, NULL)) {
      Output(mn486);
      return;
    }

    Output(mnFileCleared);
    needsCheck = true;
  } else {
    Output(mn485);
  }

  /* ...the locations file... */

  Output(mn490);

  if (AnswerIsYes(mnAreYouSure)) {
    if (!createWorkbenchFiles(GetUserID(), WB_CREATE_LOC, NULL)) {
      Output(mn487);
      return;
    }

    Output(mnFileCleared);
    needsCheck = true;
  } else {
    Output(mn491);
  }

  /* ...and the objects file */

  Output(mn492);

  if (AnswerIsYes(mnAreYouSure)) {
    if (!createWorkbenchFiles(GetUserID(), WB_CREATE_OBJ, NULL)) {
      Output(mn493);
      return;
    }

    Output(mnFileCleared);
    needsCheck = true;
  } else {
    Output(mn493);
  }
}

/*-----------------------------------------------------------------------------

  F o r m a t T e x t

-----------------------------------------------------------------------------*/

const char *
FormatText(msg_id_t messageNo, ...) {
  static char aszBuffer[16][TEXTMAX];
  static unsigned short iszBuffer = 0;

  char *pszBuffer = aszBuffer[iszBuffer];
  iszBuffer = (iszBuffer + 1) % (sizeof(aszBuffer) / sizeof(aszBuffer[0]));

  va_list ap;
  va_start(ap, messageNo);

  if (vsprintf(pszBuffer, MessageText(messageNo), ap) < 0) {
    sprintf(pszBuffer, "<<VSPRINTF %08x>>", messageNo);
  }

  va_end(ap);

  return pszBuffer;
}

/*-----------------------------------------------------------------------------

  M e s s a g e T e x t

-----------------------------------------------------------------------------*/

const char *
MessageText(msg_id_t messageNo) {
  return messages[messageNo].text;
}
