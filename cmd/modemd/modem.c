/*
 * Modem rate limiting for telnetd
 * Simulates vintage modem speeds with PPP overhead
 */

#define _DEFAULT_SOURCE

#include <sys/socket.h>
#include <sys/time.h>
#include <errno.h>
#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <unistd.h>

#include "modem.h"

// Modem specifications - time per 8 bits in microseconds
static const struct {
  long read_usec;  // Time per 8 bits received
  long write_usec; // Time per 8 bits sent
} modem_specs[MODEM_COUNT] = {
  [MODEM_V23] = { 66667, 106667 },    // 1200/75 bps: 8000000/1200, 8000000/75
  [MODEM_V32] = { 8411, 8411 },       // 9.6K: 8000000/9600*0.89 ≈ 8411
  [MODEM_V32_BIS] = { 5556, 5556 },   // 14.4K: 8000000/14400*0.89 ≈ 5556
  [MODEM_V32_TERBO] = { 4167, 4167 }, // 19.2K: 8000000/19200*0.89 ≈ 4167
  [MODEM_V34_28K] = { 2778, 2778 },   // 28.8K: 8000000/28800*0.89 ≈ 2778
  [MODEM_V34_33K] = { 2381, 2381 },   // 33.6K: 8000000/33600*0.89 ≈ 2381
  [MODEM_V90] = { 1429, 2381 },       // 56K/33.6K: 8000000/56000*0.89, 8000000/33600*0.89
  [MODEM_ISDN] = { 1124, 1124 },      // 64K: 8000000/64000*0.89 ≈ 1124
};

// Internal state
static struct timeval read_interval;
static struct timeval write_interval;
static struct timeval next_read = { 0, 0 };
static struct timeval next_write = { 0, 0 };
static struct timeval timeout;

static void set_modem(modem_t type);

void
modem_init(void) {
  // Initialize with V.34 28.8K default
  set_modem(MODEM_V34_28K);

  // Initialize timing - allow immediate first byte
  gettimeofday(&next_read, NULL);
  next_write = next_read;
}

static void
set_modem(modem_t type) {
  read_interval = (struct timeval) { 0, modem_specs[type].read_usec };
  write_interval = (struct timeval) { 0, modem_specs[type].write_usec };
}

void
set_modem_speeds(int xspeed, int rspeed) {
  // Use V.23 if either speed is below bounds
  if (xspeed <= 1200 || rspeed <= 75) {
    set_modem(MODEM_V23);
    return;
  }

  // Cap at V.90 speeds
  if (xspeed > 56000) {
    xspeed = 56000;
  }
  if (rspeed > 33600) {
    rspeed = 33600;
  }

  // Calculate intervals - 8000000 microseconds per 8 bits, with PPP overhead
  read_interval = (struct timeval) { 0, 8000000L / ((rspeed * 89) / 100) };
  write_interval = (struct timeval) { 0, 8000000L / ((xspeed * 89) / 100) };
}

struct timeval *
modem_read_timeout(void) {
  struct timeval now;
  gettimeofday(&now, NULL);

  if (!timercmp(&now, &next_read, <)) {
    return NULL; // Can read immediately
  }

  timersub(&next_read, &now, &timeout);
  return &timeout;
}

struct timeval *
modem_write_timeout(void) {
  struct timeval now;
  gettimeofday(&now, NULL);

  if (!timercmp(&now, &next_write, <)) {
    return NULL; // Can write immediately
  }

  timersub(&next_write, &now, &timeout);
  return &timeout;
}

void
modem_on_read(void) {
  // Set next allowed read time
  struct timeval now;
  gettimeofday(&now, NULL);
  timeradd(&now, &read_interval, &next_read);
}

void
modem_on_write(void) {
  // Set next allowed write time
  struct timeval now;
  gettimeofday(&now, NULL);
  timeradd(&now, &write_interval, &next_write);
}

ssize_t
modem_send(int fd, const void *buf, size_t len, int flags) {
  const unsigned char *data = (const unsigned char *) buf;
  size_t total_sent = 0;

  for (size_t i = 0; i < len; i++) {
    struct timeval write_time;

    struct timeval now;
    gettimeofday(&now, NULL);
    if (timercmp(&now, &next_write, <)) {
      struct timeval wait;
      timersub(&next_write, &now, &wait);
      usleep(wait.tv_usec); // TODO: check errors
      write_time = next_write;
    } else {
      write_time = now;
    }

    // Send exactly one byte
    ssize_t result = send(fd, &data[i], 1, flags);
    if (result < 0) {
      return total_sent > 0 ? total_sent : result;
    }
    if (result == 0) {
      break;
    }

    timeradd(&write_time, &write_interval, &next_write);
    total_sent++;
  }

  return total_sent;
}
