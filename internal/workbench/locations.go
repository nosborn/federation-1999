package workbench

import (
	"os"

	"github.com/nosborn/federation-1999/internal/text"
)

// Finds out which of the location tools the player wishes to use.
func locations() {
	file, err := OpenLocFile()
	if err != nil {
		return
	}
	defer file.Close()

	for {
		switch doMenu(text.Workbench_LocationMenu, 4) {
		case 1:
			locWrite(file)
			NeedsCheck = true
		case 2:
			locEd(file)
			NeedsCheck = true
		case 3:
			listLocations(file)
		case 4:
			changeLocs(file)
			NeedsCheck = true
		case 0:
			return
		}
	}
}

// Opens the player's location file.
func OpenLocFile() (*os.File, error) {
	pathname := LocationPathname(UserID)
	return os.OpenFile(pathname, os.O_RDWR, 0)
}

// // Allows the player to change the events for the specified location.
// func changeEvent(loc *model.DBLocation) {
// 	//
// }

// // Allows the player to toggle the flags for the specified location.
// func changeFlags(loc *model.DBLocation) {
// 	//
// }

// Changes all movement table references to the specified loc to a new one.
func changeLocs(file *os.File) {
	//
}

// // Allows the player to alter the movement table for the specified location.
// func changeMvt(loc *model.DBLocation) {
// 	//
// }

// func listFlags(loc *model.DBLocation) {
// 	//
// }

// Lists out the player's location file, suitably formatted, onto the screen.
func listLocations(file *os.File) {
	// char buffer[60];
	//
	// fmt.Print("\nLocations lister\n");
	// fmt.Print("Start location? ");
	// input := getInput(true);
	// start, _ = strconv.Atoi(input);
	// fmt.Print("\nStop location? ");
	// input = getInput(true);
	// stop, _ = strconv.Atoi(input);
	//
	// if start < 9 {
	// 	start = 9
	// }
	//
	// var loc DBLocation
	// counter := start
	//
	// // lseek(f_num, sizeof(loc) * (start - 1), SEEK_SET);
	//
	// while (read(f_num, &loc, sizeof(loc)) > 0) {
	// 	fmt.Printf("\nLocation number: %d\n", counter++)
	// 	locDisplay(&loc);
	// 	fmt.Print("\n-------------------------------------\n")
	//
	// 	if counter > stop {
	// 		break;
	// 	}
	// }
}

// // Displays the specified location.
// func locDisplay(loc *model.DBLocation) {
// }

// Allows the player to edit his/her location database.
func locEd(file *os.File) {
}

// Writes new locations into the the player's location file.
func locWrite(file *os.File) {
	//
}

// // Gets all the information for a new location from the player.
// func newLoc(index int, file *os.File) {
// 	//
// }
