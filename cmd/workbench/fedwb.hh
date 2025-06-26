/******************************************************************************

  Copyright (c) 1987-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

******************************************************************************/

#ifndef FEDWB_HH
#define FEDWB_HH

#include <sys/types.h>

#include <fed.hh>

#include "messages.hh"

typedef struct dbevent dbevent_t;
typedef struct dblocation dblocation_t;
typedef struct dbobject dbobject_t;

const size_t SHIP_SIZE = 8; // Locations reserved for spaceship

extern bool AnswerIsYes(msg_id_t);
extern bool checkPlanet();
extern unsigned doMenu(msg_id_t, unsigned);
extern void download();
extern bool editor(char *, size_t);
extern void Events();
extern const char *FormatText(msg_id_t, ...);
extern const dbevent_t *getEvent(int);
extern int getEventCount();
extern void GetInput(char *, size_t, bool);
extern const dblocation_t *getLocation(int);
extern int getLocationCount();
extern int getNumber(const char *);
extern int getNumber(msg_id_t);
extern const dbobject_t *getObject(const char *, int *);
extern const dbobject_t *getObject(int);
extern int getObjectCount();
extern account_id_t GetUserID();
extern bool isBlank(const char *);
extern void locationMenu();
extern void Locations();
extern const char *MessageText(msg_id_t);
extern void objectMenu();
extern void Objects();
extern int OpenEvFile();
extern int OpenLocFile();
extern int OpenObjFile();
extern void Output(msg_id_t);
extern void Output(const char *);
extern int PromptForInteger(const char *);
extern int PromptForInteger(msg_id_t);
extern bool saveEvent(const dbevent_t &, int);
extern bool saveLocation(const dblocation_t &, int);
extern bool saveObject(const dbobject_t &, int);
extern void upload();

#endif /* FEDWB_HH */
