/*
 * Modem rate limiting for telnetd
 * Simulates vintage modem speeds with PPP overhead
 */

#ifndef MODEM_H
#define MODEM_H

#include <sys/time.h>
#include <sys/types.h>

/*
 * Modem types
 */
typedef enum {
  MODEM_V23,
  MODEM_V32,
  MODEM_V32_BIS,
  MODEM_V32_TERBO,
  MODEM_V34_28K,
  MODEM_V34_33K,
  MODEM_V90,
  MODEM_ISDN,
  MODEM_COUNT
} modem_t;

void modem_init(void);
void set_modem_speeds(int xspeed, int rspeed);

struct timeval *modem_read_timeout(void);
struct timeval *modem_write_timeout(void);

void modem_on_read(void);
void modem_on_write(void);

ssize_t modem_send(int fd, const void *buf, size_t len, int flags);

#endif // MODEM_H
