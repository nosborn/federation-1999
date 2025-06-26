package workbench

import (
	"fmt"
	"strings"

	"github.com/nosborn/federation-1999/internal/text"
)

func answerIsYes(msgNo text.MsgNum) bool {
	yes := text.Msg(text.Workbench_Yes)
	no := text.Msg(text.Workbench_No)

	for {
		fmt.Printf("%s [%s/%s] ", text.Msg(msgNo), yes, no)
		answer := getInput(true)

		if answer != "" {
			if strings.EqualFold(answer, yes[:len(answer)]) {
				return true
			}
			if strings.EqualFold(answer, no[:len(answer)]) {
				return false
			}
		}
	}
}

// func promptForInteger(prompt string) int {
// 	fmt.Print(prompt)
//
// 	// TODO
//
// 	return 0
// }
