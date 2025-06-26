/******************************************************************************

  Copyright (c) 1987-1997 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: drvio.cc,v 1.2 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <ctype.h>
#include <limits.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>

#include "fedwb.hh"

bool
AnswerIsYes(msg_id_t idPrompt) {
  const char *pszYes = MessageText(mnYes);
  const char *pszNo = MessageText(mnNo);

  bool yes;

  for (;;) {
    char szAnswer[4];

    printf("%s [%s/%s] ", MessageText(idPrompt), pszYes, pszNo);
    GetInput(szAnswer, sizeof(szAnswer), true);

    if (strlen(szAnswer) > 0) {
      if (strncasecmp(pszYes, szAnswer, strlen(szAnswer)) == 0) {
        yes = true;
        break;
      } else if (strncasecmp(pszNo, szAnswer, strlen(szAnswer)) == 0) {
        yes = false;
        break;
      }
    }
  }

  return yes;
}

void
Output(msg_id_t msg) {
  Output(MessageText(msg));
}

void
Output(const char *text) {
  printf("%s", text);
  fflush(stdout);
}

int
PromptForInteger(const char *pszPrompt) {
  bool fGoodInput = false;
  long lInput = 0;

  Output(pszPrompt);

  do {
    char szInput[12];
    char *pszEnd;

    GetInput(szInput, sizeof(szInput), true);

    if (strlen(szInput) > 0) {
      lInput = strtol(szInput, &pszEnd, 10);

      if (*pszEnd == '\0') {
        fGoodInput = true;
      }
    }
  } while (!fGoodInput);

  if (lInput > INT_MAX) {
    lInput = INT_MAX;
  } else if (lInput < INT_MIN) {
    lInput = INT_MIN;
  }

  return static_cast<int>(lInput);
}

int
PromptForInteger(msg_id_t idPrompt) {
  return PromptForInteger(MessageText(idPrompt));
}
