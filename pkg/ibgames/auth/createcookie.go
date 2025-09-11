package auth

import (
	"net"

	"github.com/nosborn/federation-1999/pkg/ibgames"
)

func CreateCookie(_ net.Addr, uid ibgames.AccountID, sid *string) CookieResult {
	// TODO: implementation
	return CookieOK
}
