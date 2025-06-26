//go:build linux

/*
 * This function implements the "bigcrypt" algorithm.
 *
 * Based on the implementation from libxcrypt by Thorsten Kukuk, Andy Phillips,
 * and others.
 *
 * This simplified version is a non-reentrant, drop-in replacement for HP-UX
 * bigcrypt(), and it calls the standard non-reentrant crypt().
 */

#define _GNU_SOURCE
#include <crypt.h>
#include <string.h>

#include "bigcrypt.h"

#define MAX_PASS_LEN 16
#define SEGMENT_SIZE 8
#define SALT_SIZE 2
#define KEYBUF_SIZE ((MAX_PASS_LEN * SEGMENT_SIZE) + SALT_SIZE)
#define ESEGMENT_SIZE 11
#define CBUF_SIZE ((MAX_PASS_LEN * ESEGMENT_SIZE) + SALT_SIZE + 1)

const char *bigcrypt(const char *key, const char *salt) {
  static char dec_c2_cryptbuf[CBUF_SIZE]; /* static storage area */

  size_t keylen;
  unsigned long int n_seg, j;
  char *cipher_ptr, *plaintext_ptr, *tmp_ptr, *salt_ptr;
  char keybuf[KEYBUF_SIZE + 1];

  /* reset arrays */
  memset(keybuf, 0, KEYBUF_SIZE + 1);
  memset(dec_c2_cryptbuf, 0, CBUF_SIZE);

  /* fill KEYBUF_SIZE with key */
  strncpy(keybuf, key, KEYBUF_SIZE);

  /* deal with case that we are doing a password check for a
     conventially encrypted password: the salt will be
     SALT_SIZE+ESEGMENT_SIZE long. */
  if (strlen(salt) == (SALT_SIZE + ESEGMENT_SIZE))
    keybuf[SEGMENT_SIZE] = '\0'; /* terminate password early(?) */

  keylen = strlen(keybuf);

  if (!keylen) {
    n_seg = 1;
  } else {
    /* work out how many segments */
    n_seg = 1 + ((keylen - 1) / SEGMENT_SIZE);
  }

  if (n_seg > MAX_PASS_LEN)
    n_seg = MAX_PASS_LEN; /* truncate at max length */

  /* set up some pointers */
  cipher_ptr = dec_c2_cryptbuf;
  plaintext_ptr = keybuf;

  /* do the first block with supplied salt */
  tmp_ptr = crypt(plaintext_ptr, salt);

  /* and place in the static area */
  strncpy(cipher_ptr, tmp_ptr, 13);
  cipher_ptr += ESEGMENT_SIZE + SALT_SIZE;
  plaintext_ptr += SEGMENT_SIZE; /* first block of SEGMENT_SIZE */

  /* change the salt (1st 2 chars of previous block) */
  salt_ptr = cipher_ptr - ESEGMENT_SIZE;

  /* if there is more than one block encrypt them... */
  if (n_seg > 1) {
    for (j = 2; j <= n_seg; j++) {

      tmp_ptr = crypt(plaintext_ptr, salt_ptr);

      /* skip the salt for seg!=0 */
      strncpy(cipher_ptr, (tmp_ptr + SALT_SIZE), ESEGMENT_SIZE);

      cipher_ptr += ESEGMENT_SIZE;
      plaintext_ptr += SEGMENT_SIZE;
      salt_ptr = cipher_ptr - ESEGMENT_SIZE;
    }
  }
  /* this is the <NUL> terminated encrypted password */
  return dec_c2_cryptbuf;
}
