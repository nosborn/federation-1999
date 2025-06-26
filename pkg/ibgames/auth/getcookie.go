package auth

import (
	"database/sql"
	"net"
	"time"

	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/nosborn/federation-1999/pkg/ibgames/db"
)

func GetCookie(sid string, _ net.Addr, uidp *ibgames.AccountID) CookieResult {
	*uidp = ibgames.AccountID(0)

	var uid ibgames.AccountID
	var expire int64
	err := db.QueryRow("SELECT uid, expire FROM cookies WHERE sid = ?", sid).Scan(&uid, &expire)
	if err != nil {
		if err == sql.ErrNoRows {
			return CookieNotFound
		}
		return CookieError
	}

	now := time.Now().Unix()

	if expire < now {
		_, err = db.Exec("DELETE FROM cookies WHERE sid = ?", sid)
		if err != nil {
			return CookieError
		}
		return CookieNotFound
	}

	expire = now + (30 * 60)

	_, err = db.Exec("UPDATE cookies SET expire = ? WHERE sid = ?", expire, sid)
	if err != nil {
		return CookieError
	}

	*uidp = uid
	return CookieOK
}
