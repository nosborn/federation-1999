/******************************************************************************

  Copyright (c) 1995-1999 Alan Lenton & Interactive Broadcasting Ltd.
  All Rights Reserved.

  No part of this software may be reproduced, transmitted, transcribed, stored
  in a retrieval system, or translated into any human or computer language, in
  any form or by any means, electronic, mechanical, magnetic, optical, manual
  or otherwise, without the express written permission of the copyright holder.

  $Id: goodies.h,v 1.3 1999/03/17 12:16:22 nick Exp $

******************************************************************************/

#ifndef __IBGAMES_GOODIES_H
#define __IBGAMES_GOODIES_H

#include <sys/types.h>
#include <fcntl.h>
#include <stdio.h>
#include <unistd.h>

/* Enumeration type for IB_lseek() */

typedef enum {
  seekSet = SEEK_SET,
  seekCur = SEEK_CUR,
  seekEnd = SEEK_END
} seek_whence_t;

#ifdef __cplusplus
extern "C" {
#endif

/*
 * Trivial wrapper functions for system calls. Return values and errno are as
 * per the corresponding system call except that EINTR is handled and not
 * returned to the caller.
 *
 * IB_unlink treats an attempt to delete a non-existent file as success.
 */
int IB_close(int fildes);
int IB_creat(const char *path, mode_t mode);
int IB_dup(int fildes);
int IB_dup2(int fildes, int fildes2);
int IB_ftruncate(int fildes, off_t length);
int IB_lockf(int fildes, int function, off_t size);
off_t IB_lseek(int fildes, off_t offset, seek_whence_t whence);
int IB_open(const char *path, int oflag, ... /* [mode_t mode] */);
ssize_t IB_read(int fildes, void *buf, size_t nbyte);
ssize_t IB_recv(int s, void *buf, size_t len, int flags);
ssize_t IB_send(int s, const void *msg, size_t len, int flags);
int IB_unlink(const char *path);
pid_t IB_wait(int *stat_loc);
pid_t IB_waitpid(pid_t pid, int *stat_loc, int options);
ssize_t IB_write(int fildes, const void *buf, size_t nbyte);

/*
 * Trivial functions to manipulate file descriptors. Return values and errno
 * are as per fcntl().
 */
int IB_blocking(int fildes);
int IB_closeOnExec(int fildes);
int IB_nonBlocking(int fildes);

/*
 * Close all open files except for the first keepfds (ie, 0 to keepfds-1). This
 * is typically only a useful thing to do when daemon-ising or in a fork/exec
 * window. Returns -1 with errno from pstat() or (possibly) close(). The latter
 * is unlikely; EBADF and EINTR are eaten, which in theory eliminates
 * everything except problems involving NFS filesystems.
 */
int IB_closeFiles(int keepfds);

/*
 * Get the size of a file. Returns -1 with errno from fstat() on error, file
 * size in bytes on success.
 */
off_t IB_fileSize(int fildes);

/* Functions which do real work! */

int IB_getRunLock(const char *);

/*
 * Return the home directory for the effective uid. The path is obtained from
 * /etc/passwd on the first call and cached; this is convenient for the caller
 * but makes it unsuitable for use when the euid may be change.
 */
const char *IB_homeDir(void);

/*
 * Change the value of the 'command line' as shown by ps. Useful if you want
 * some sort of per-process info to show up - user id or IP address for a
 * driver process, number of active connections for a server, that kind of
 * thing. Keep the title fairly short - less than 64 bytes for HP-UX.
 */
void IB_setProcTitle(const char *, ...);

/*
 * Remove leading and trailing whitespace from a string. Returns a pointer to
 * the first non-whitespace charecter. The input string is modified.
 */
char *IB_trimString(char *s);

/*
 * Send RFC-822 mail.
 */
FILE *IB_beginMail(void);
int IB_endMail(FILE *stream);
int IB_sendMail(const char *args, ...);

/*
 * Sandbox setup.
 */
void IB_sandbox(const char *path, const char *name);

/*
 * Go-faster (we hope) gethostbyname. Don't use this for lookups from sources
 * that lack DNS-style trailing dot magic, such as local files or NIS maps.
 */
struct hostent *IB_getHostByName(const char *name);

/*
 * Check for bad substrings like 'ogin:' and 'word:' in a string.
 */
int IB_containsPrompt(const char *s);

#ifdef __cplusplus
}
#endif

/*
 * Non-overflowing math functions. The return value is limited to
 *
 *   (type_MIN + 1) <= x <= (signed_type_MAX - 1)
 *
 * or
 *
 *   0 <= x <= (unsigned_type_MAX - 1)
 *
 * which conveniently leaves type_MIN and type_MAX available for use as
 * indicator values. Note that it's up to the caller to do any necessary
 * interpretation; neither are treated as special if you pass them.
 *
 * Overloading obviously precludes these from use in C code. You're on your
 * own if you want safe long long functions!
 */

#ifdef __cplusplus
extern short IB_add(const short, const short);
extern unsigned short IB_add(const unsigned short, const unsigned short);
extern int IB_add(const int, const int);
extern unsigned int IB_add(const unsigned int, const unsigned int);
extern long IB_add(const long, const long);
extern unsigned long IB_add(const unsigned long, const unsigned long);
#endif /* __cplusplus */

#endif /* __IBGAMES_GOODIES_H */
