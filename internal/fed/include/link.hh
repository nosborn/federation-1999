/******************************************************************************

  Copyright (c) 1987-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: link.hh,v 1.4 1999/04/24 21:11:15 nick Exp $

******************************************************************************/

#ifndef _LINK_HH
#define _LINK_HH

#define LINK_ADDRESS_SIZE 16
#define LINK_HOSTNAME_SIZE 64

#define LINK_FORMAT "%u %s %s %d\n"
#define LINK_TOKENS 4
#define LINK_SOCKET "%s/.fedtpd.socket" /* add IB_homeDir() */

struct connect_info {
  char addr[LINK_ADDRESS_SIZE];
  char name[LINK_HOSTNAME_SIZE];
};

typedef struct connect_info connect_info_t;

enum {
  DLE = 16
};

enum {
  leAck = 1,
  leSpy,
  leTrace
};

#endif /* _LINK_HH */
