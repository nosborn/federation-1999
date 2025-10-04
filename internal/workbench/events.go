package workbench

import (
	"os"

	"github.com/nosborn/federation-1999/internal/text"
)

// Finds out which of the event tools the player wishes to use.
func events() {
	file, err := OpenEvFile()
	if err != nil {
		return
	}
	defer file.Close()

	for {
		switch doMenu(text.Workbench_EventMenu, 3) {
		case 1:
			evWrite(file)
			NeedsCheck = true
		case 2:
			evEdit(file)
			NeedsCheck = true
		case 3:
			evList(file)
		case 0:
			return
		}
	}
}

// Opens the player's events file.
func OpenEvFile() (*os.File, error) {
	pathname := EventPathname(UserID)
	return os.OpenFile(pathname, os.O_RDWR, 0)
}

// // Displays the selected event for the user's inspection.
// func evDisplay(event *model.DBEvent, index int) {
// 	// TODO
// }

// Allows the player to edit existing events int the file.
func evEdit(file *os.File) {
	// TODO
}

// Allows the player to list the events in his/her datafile.
func evList(file *os.File) {
	// TODO
}

// Allows the player to write new events into his/her datafile.
func evWrite(file *os.File) {
	// TODO
}

// // Finds out which attribute the player wishes to change or test.
// func getAttributeChange(event model.DBEvent) *int16 {
// 	for {
// 		fmt.Print(text.Msg(text.Workbench_AttributeChangePrompt))
// 		input := getInput(true)
//
// 		if strings.EqualFold(input, "STR") {
// 			return &event.Field1
// 		}
// 		if strings.EqualFold(input, "STA") {
// 			return &event.Field2
// 		}
// 		if strings.EqualFold(input, "INT") {
// 			return &event.Field3
// 		}
// 		if strings.EqualFold(input, "DEX") {
// 			return &event.Field4
// 		}
//
// 		fmt.Print(text.Msg(text.Workbench_AttributeChangeInvalid))
// 	}
// }
