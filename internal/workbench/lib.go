package workbench

import (
	"fmt"
	"os"

	"github.com/nosborn/federation-1999/pkg/ibgames"
	"golang.org/x/sys/unix"
)

const ( // File creation flags for CreateFiles.
	WB_CREATE_EVT = 0x00000001
	WB_CREATE_LOC = 0x00000002
	WB_CREATE_OBJ = 0x00000004
	WB_CREATE_ALL = WB_CREATE_EVT + WB_CREATE_LOC + WB_CREATE_OBJ
)

const ( // Return values from Access.
	WB_ACCESS_OK = iota
	WB_NO_FILES
	WB_CANT_WRITE
)

const ( // Size limits for player planets.
	EVENT_LIMIT    = 25
	LOCATION_LIMIT = 120
	OBJECT_LIMIT   = 14
)

func createFiles(uid ibgames.AccountID, which int, mini string) bool {
	if (which & WB_CREATE_EVT) == WB_CREATE_EVT {
		pathname := EventPathname(uid)
		evFile, err := os.OpenFile(pathname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o660)
		if err != nil {
			return false
		}
		defer evFile.Close()
	}

	if (which & WB_CREATE_LOC) == WB_CREATE_LOC {
		pathname := LocationPathname(uid)
		locFile, err := os.OpenFile(pathname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o660)
		if err != nil {
			return false
		}
		defer locFile.Close()
	}

	if (which & WB_CREATE_OBJ) == WB_CREATE_OBJ {
		pathname := ObjectPathname(uid)
		objFile, err := os.OpenFile(pathname, os.O_CREATE|os.O_RDWR|os.O_TRUNC, 0o660)
		if err != nil {
			return false
		}
		defer objFile.Close()
	}

	return true
}

func DeleteFiles(uid ibgames.AccountID) {
	os.Remove(EventPathname(uid))
	os.Remove(LocationPathname(uid))
	os.Remove(ObjectPathname(uid))
}

func EventPathname(uid ibgames.AccountID) string {
	return fmt.Sprintf("%s.e", basename(uid))
}

func LocationPathname(uid ibgames.AccountID) string {
	return fmt.Sprintf("%s.l", basename(uid))
}

func ObjectPathname(uid ibgames.AccountID) string {
	return fmt.Sprintf("%s.o", basename(uid))
}

func Access(uid ibgames.AccountID) int {
	eventFile := EventPathname(uid)
	locationFile := LocationPathname(uid)
	objectFile := ObjectPathname(uid)

	// See if the files exist.
	if err := unix.Access(eventFile, unix.F_OK); err != nil {
		return WB_NO_FILES
	}
	if err := unix.Access(locationFile, unix.F_OK); err != nil {
		return WB_NO_FILES
	}
	if err := unix.Access(objectFile, unix.F_OK); err != nil {
		return WB_NO_FILES
	}

	// OK, now see if we'll be able to write to them.
	if err := unix.Access(eventFile, unix.W_OK); err != nil {
		return WB_CANT_WRITE
	}
	if err := unix.Access(locationFile, unix.W_OK); err != nil {
		return WB_CANT_WRITE
	}
	if err := unix.Access(objectFile, unix.W_OK); err != nil {
		return WB_CANT_WRITE
	}

	// All OK.
	return WB_ACCESS_OK
}

func basename(uid ibgames.AccountID) string {
	return fmt.Sprintf("data/workbench%d/%d", uid%10, uid)
}
