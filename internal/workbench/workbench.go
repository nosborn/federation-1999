package workbench

import (
	"bufio"
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/nosborn/federation-1999/internal/text"
	"github.com/nosborn/federation-1999/pkg/ibgames"
)

var (
	CheckOnly  bool
	NeedsCheck bool
	NoExchange bool
	UserID     ibgames.AccountID
)

func doMenu(msgNo text.MsgNum, maxChoice int) int {
	fmt.Print(text.Msg(msgNo))

	for {
		fmt.Print(text.Msg(text.Workbench_MenuPrompt))
		input := getInput(true)
		choice, err := strconv.Atoi(input)
		if err != nil || choice > maxChoice {
			continue
		}
		return choice
	}
}

func getInput(trim bool) string {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	if err := scanner.Err(); err != nil {
		fmt.Fprintf(os.Stderr, "getInput: EOF on stdin\n")
		os.Exit(1)
	}
	input := scanner.Text()
	if trim {
		input = strings.TrimSpace(input)
	}
	fmt.Println(input)
	return input
}

func MainMenu() {
	if Access(UserID) != WB_ACCESS_OK {
		fmt.Print(text.Msg(text.Workbench_77))
		return
	}

	for {
		switch doMenu(text.Workbench_MainMenu, 7) {
		case 0:
			if NeedsCheck {
				fmt.Print(text.Msg(text.Workbench_UncheckedPreamble))
				if !answerIsYes(text.Workbench_UncheckedPrompt) {
					break
				}
			}
			return
		case 1:
			events()
		case 2:
			objects()
		case 3:
			locations()
		case 4:
			CheckPlanet()
		case 5:
			clearAllFiles()
		case 6:
			// download()
		case 7:
			// upload()
		}
	}
}

func ParseArguments() bool {
	flag.BoolVar(&CheckOnly, "c", false, "help message for flag c")
	flag.BoolVar(&NoExchange, "n", false, "help message for flag n")
	flag.Parse()

	if flag.NArg() != 1 {
		return false
	}

	arg, err := strconv.Atoi(flag.Arg(0))
	if err != nil || arg < ibgames.MinAccountID || arg >= ibgames.MaxAccountID {
		log.Print("Bad id argument")
		return false
	}
	UserID = ibgames.AccountID(arg)

	return true
}

func clearAllFiles() {
	// The events file...
	fmt.Print(text.Msg(text.Workbench_484))
	if answerIsYes(text.Workbench_AreYouSure) {
		if !createFiles(UserID, WB_CREATE_EVT, "") {
			fmt.Print(text.Msg(text.Workbench_486))
			return
		}
		fmt.Print(text.Msg(text.Workbench_FileCleared))
		NeedsCheck = true
	} else {
		fmt.Print(text.Msg(text.Workbench_485))
	}

	// ...the locations file...
	fmt.Print(text.Msg(text.Workbench_490))
	if answerIsYes(text.Workbench_AreYouSure) {
		if !createFiles(UserID, WB_CREATE_LOC, "") {
			fmt.Print(text.Msg(text.Workbench_487))
			return
		}
		fmt.Print(text.Msg(text.Workbench_FileCleared))
		NeedsCheck = true
	} else {
		fmt.Print(text.Msg(text.Workbench_491))
	}

	// ...and the objects file.
	fmt.Print(text.Msg(text.Workbench_492))
	if answerIsYes(text.Workbench_AreYouSure) {
		if !createFiles(UserID, WB_CREATE_OBJ, "") {
			fmt.Print(text.Msg(text.Workbench_487))
			return
		}
		fmt.Print(text.Msg(text.Workbench_FileCleared))
		NeedsCheck = true
	} else {
		fmt.Print(text.Msg(text.Workbench_491))
	}
}
