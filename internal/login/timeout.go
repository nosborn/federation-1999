package login

import (
	"fmt"
	"os"
	"time"
)

type Timeout struct {
	timer *time.Timer
}

func SetTimeout() *Timeout {
	return &Timeout{
		timer: time.AfterFunc(300*time.Second, timeoutFunc),
	}
}

func timeoutFunc() {
	fmt.Print("\n\nLogin timed out.\n")
	os.Exit(0)
}

func (t *Timeout) Cancel() {
	t.timer.Stop()
}
