package workbench

import (
	"os"

	"github.com/nosborn/federation-1999/internal/text"
)

// Finds out which of the object tools the player wishes to use.
func objects() {
	file, err := OpenObjFile()
	if err != nil {
		return
	}
	defer file.Close()

	for {
		switch doMenu(text.Workbench_ObjectMenu, 3) {
		case 1:
			objWrite(file)
			NeedsCheck = true
		case 2:
			objEdit(file)
			NeedsCheck = true
		case 3:
			objList(file)
		case 0:
			return
		}
	}
}

// Opens the player's object file.
func OpenObjFile() (*os.File, error) {
	pathname := ObjectPathname(UserID)
	return os.OpenFile(pathname, os.O_RDWR, 0)
}

// func getMobileDetails(obj *model.DBObject) {
// 	//
// }

// // Allows player to edit object descs.
// func obDesEdit(obj *model.DBObject) {
// 	//
// }

// // Edit events associated with an object.
// func obEvEdit(obj *model.DBObject) {
// 	//
// }

// // Edit flags associated with an object.
// func obFlagsEdit(obj *model.DBObject) {
// 	//
// }

// // Displays an object's details.
// func objDisplay(obj *model.DBObject) {
// 	//
// }

// Allows a player to edit the attributes of objects in the object file.
func objEdit(file *os.File) {
	//
}

// List out the contents of the object file to screen.
func objList(file *os.File) {
	//
}

// Allows the player to write a new object into his/her datafile.
func objWrite(file *os.File) {
	//
}

// // Edit object location limits.
// func objLocEdit(obj *model.DBObject) {
// 	//
// }

// // Edit mobile attributes.
// func obMobEdit(obj *model.DBObject) {
// 	//
// }

// // Edit ship attributes.
// func obShipEdit(obj *model.DBObject) {
// 	//
// }

// func validateName(name string) bool {
// 	// TODO
// 	return true
// }
