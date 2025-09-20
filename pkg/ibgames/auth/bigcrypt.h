//go:build linux

#ifndef __IBGAMES_BIGCRYPT_H
#define __IBGAMES_BIGCRYPT_H

#ifdef __cplusplus
extern "C" {
#endif

const char *bigcrypt(const char *key, const char *salt);

#ifdef __cplusplus
}
#endif

#endif /* __IBGAMES_BIGCRYPT_H */
