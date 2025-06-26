/******************************************************************************

  Copyright (c) 1987-1997 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  -----------------------------------------------------------------------------

  Federation - online mini editor.

  Designed to be called from other functions. Copies the edited text into a
  buffer provided by the calling function. It is the responsibility of the
  calling function to ensure that the buffer is large enough to contain the
  resulting text. Max line capacity of the editor is MAXLINE.

  $Id: editor.cc,v 1.2 1998/12/18 14:18:32 nick Exp $

******************************************************************************/

#include <string>

#include <ctype.h>
#include <errno.h>
#include <math.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>

#include <ibgames/log.h>

#include "fedwb.hh"

#define MAXLINE 40
#define MAXINPUT 300

static char *_ChangeText(const char *);

// Top level function for mini-editor.
bool
editor(char *output, size_t size) {
  dbgPrecondition(output != NULL);
  dbgPrecondition(size > 0);

  char *buffer[MAXLINE];
  int count, index;
  int line = 0, number, flag;
  unsigned char_count;
  char input[MAXINPUT], *new_text;

  for (count = 0; count < MAXLINE; count++) {
    buffer[count] = NULL;
  }

  if (*output != '\0') {
    for (index = 0;;) {
      buffer[line] = static_cast<char *>(calloc(80, sizeof(char)));

      if (buffer[line] == NULL) {
        for (count = 0; count < line; count++) {
          free(buffer[line]);
        }

        IB_log("editor: calloc() failed errno=%d", errno);
        return false;
      }

      flag = false;
      strncpy(input, output + index, 74);
      input[74] = '\0';

      for (count = 0; count < 74 || input[count] != '\0'; count++) {
        if (input[count] == '\n') {
          input[count] = '\0';
          strcpy(buffer[line], input);
          index += (count + 1);
          line++;

          buffer[line] = static_cast<char *>(calloc(2, sizeof(char)));

          if (buffer[line] == NULL) {
            for (count = 0; count < line; count++) {
              free(buffer[line]);
            }

            IB_log("editor: calloc() failed errno=%d", errno);
            return false;
          }

          line++;
          flag = true;
          break;
        }
      }

      if (flag) {
        continue;
      }

      if (strlen(input) < 74) {
        strcpy(buffer[line++], input);
        break;
      }

      for (count = 74; count >= 0; count--) {
        if (isspace(input[count])) {
          input[count] = '\0';
          strcpy(buffer[line++], input);
          index += (count + 1);
          break;
        }
      }
    }
  }

  for (;;) {
    printf("%2d> ", line + 1);
    GetInput(input, MAXINPUT, false);

    if (input[0] == '*') {
      bool addSpace;
      std::string cookedText;

      switch (tolower(input[1])) {
        case 'c': // Change a line
          number = atoi(&input[3]) - 1;

          if (number < 0 || number >= line) {
            break;
          }

          new_text = _ChangeText(buffer[number]);

          if (new_text == NULL) {
            break;
          }

          free(buffer[number]);
          buffer[number] = new_text;

          break;

        case 'd': // Delete a line
          number = atoi(&input[3]) - 1;

          if (number < 0 || number >= line) {
            break;
          }

          for (count = number; count < line; count++) {
            buffer[count] = buffer[count + 1];
          }

          free(buffer[line]);
          buffer[line] = NULL;
          line--;

          break;

        case 'h': // Request for help
          Output(mnEditorHelp);
          break;

#if defined(BROKEN_INSERT_COMMAND)
        case 'i': // Insert
          if (line == MAXLINE) {
            Output(mnEditorFull);
            continue;
          }

          strncpy(temp, &input[3], 9);
          temp[9] = '\0';

          for (count = 0; count < 9; count++) {
            if (temp[count] == ',') {
              temp[count] = '\0';
              break;
            }
          }

          text_start = count + 4;

          if (((number = atoi(temp)) >= line) || (number < 0)) {
            break;
          }

          for (count = line; count > number; count--) {
            buffer[count] = buffer[count - 1];
          }

          buffer[number] = static_cast<char *>(malloc(strlen(input)));

          if (buffer[number] == NULL) {
            IB_log("editor: malloc() failed errno=%d", errno);
            return false;
          }

          strcpy(buffer[number], &input[text_start]);
          line++;
          break;
#endif // defined(BROKEN_INSERT_COMMAND)

        case 'l': // List contents
          Output("\n");

          for (count = 0; count < line; count++) {
            printf("%2d> %s\n", count + 1, buffer[count]);
          }

          Output("\n");
          break;

        case 's': // Store the edited text
          addSpace = false;
          cookedText = "";

          for (count = 0; count < line; count++) {
            dbgCheck(buffer[count] != NULL);

            if (isBlank(buffer[count])) {
              cookedText += '\n';
              addSpace = false;
            } else {
              if (addSpace) {
                cookedText += ' ';
              }

              cookedText += buffer[count];
              addSpace = true;
            }

            free(buffer[count]);
          }

          if (cookedText.size() > size) {
            Output(FormatText(mnEditorTextTooBig, cookedText.size()));
            break;
          }

          strcpy(output, cookedText.c_str());
          return true;

        case 't': // Total character count
          addSpace = false;
          char_count = 0;

          for (count = 0; count < line; count++) {
            dbgCheck(buffer[count] != NULL);

            if (isBlank(buffer[count])) {
              char_count += 1;
              addSpace = false;
            } else {
              if (addSpace) {
                char_count += 1;
              }

              char_count += strlen(buffer[count]);
              addSpace = true;
            }
          }

          Output(FormatText(mnEditorCharacterCount, char_count));
          break;

        case 'w': // Clear buffer
          for (count = 0; count < line; count++) {
            if (buffer[line] != NULL) {
              free(buffer[line]);
              buffer[line] = NULL;
            }
          }

          line = 0;
          break;

        case 'x': // Exit abandoning the edit
          for (count = 0; count < line; count++) {
            free(buffer[count]);
          }

          return false;
      }
    } else {
      if (line == MAXLINE) {
        Output(mnEditorFull);
        continue;
      }

      dbgCheck(buffer[line] == NULL);
      buffer[line] = static_cast<char *>(malloc(strlen(input) + 1));

      if (buffer[line] == NULL) {
        for (count = 0; count < line; count++) {
          free(buffer[line]);
        }

        IB_log("editor: malloc() failed errno=%d", errno);
        return false;
      }

      strcpy(buffer[line], input);
      line++;
      continue;
    }
  }
}

/*-----------------------------------------------------------------------------

  _ C h a n g e T e x t

  Handles changing a string in a line.

-----------------------------------------------------------------------------*/

static char *
_ChangeText(const char *text) {
  dbgPrecondition(text != NULL);

  char oldString[MAXINPUT], newString[MAXINPUT];

  Output(FormatText(mnOldStringPrompt, text));
  GetInput(oldString, MAXINPUT, false);

  const char *subString = strstr(text, oldString);

  if (subString == NULL) {
    Output(FormatText(mnOldStringNotFound, oldString));
    return NULL;
  }

  Output(mnNewStringPrompt);
  GetInput(newString, MAXINPUT, false);

  const size_t newSize = strlen(text) - strlen(oldString) + strlen(newString);
  char *newText = static_cast<char *>(malloc(newSize + 1));

  if (newText == NULL) {
    Output("\nInsufficient memory - abandoning insert.\n");
    return NULL;
  }

  strncpy(newText, text, subString - text);
  strcpy(newText + (subString - text), newString);
  strcat(newText, subString + strlen(oldString));

  // strcpy(new_text, text);
  // sub_string = strstr(new_text, old_string);
  //*sub_string = '\0';
  // strcat(new_text, new_string);
  // sub_string = strstr(text, old_string);
  // strcat(new_text, sub_string + strlen(old_string));

  dbgCheck(memchr(newText, '\0', newSize + 1) != NULL);
  return newText;
}
