/******************************************************************************

  Copyright (c) 1987-1998 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

******************************************************************************/

#ifndef WORKBENCH_H
#define WORKBENCH_H

#include <stddef.h>

#include <ibgames.h>

// File creation flags for CreateWorkbenchFiles.

enum {
  WB_CREATE_EVT = 0x00000001,
  WB_CREATE_LOC = 0x00000002,
  WB_CREATE_OBJ = 0x00000004
};

#define WB_CREATE_ALL (WB_CREATE_EVT | WB_CREATE_LOC | WB_CREATE_OBJ)

// Return values from WorkbenchAccess.

enum {
  WB_ACCESS_OK,
  WB_NO_FILES,
  WB_CANT_WRITE
};

// Size limits for player planets.
const size_t EVENT_LIMIT = 25;
const size_t LOCATION_LIMIT = 120;
const size_t OBJECT_LIMIT = 14;

extern bool createWorkbenchFiles(account_id_t, unsigned, const char *);
extern void deleteWorkbenchFiles(account_id_t);
extern void getEventPathname(account_id_t, char *, size_t);
extern void getLocationPathname(account_id_t, char *, size_t);
extern void getObjectPathname(account_id_t, char *, size_t);
extern int workbenchAccess(account_id_t);

#endif /* WORKBENCH_H */
