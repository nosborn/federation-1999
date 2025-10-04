package text

import (
	"fmt"
	"log"
)

func Msg(num MsgNum, args ...any) string {
	if num < 1 || num > maxMsgNum {
		// log.Panicf("message not found: %#v", n)
		log.Printf("message not found: %#v", num)
		return fmt.Sprintf("<<OOPS %#v>>", num)
	}
	msg := msgs[num]
	if msg.f {
		return fmt.Sprintf(msg.t, args...)
	}
	return msg.t
}

func NoExitMsgNum(sysLoc int) (MsgNum, bool) {
	switch sysLoc {
	case 1:
		return NO_MOVEMENT_1, true
	case 2:
		return NO_MOVEMENT_2, true
	case 7:
		return NO_MOVEMENT_7, true
	case 12:
		return NO_MOVEMENT_12, true
	case 23:
		return NO_MOVEMENT_23, true
	case 26:
		return NO_MOVEMENT_26, true
	case 30:
		return NO_MOVEMENT_30, true
	case 32:
		return NO_MOVEMENT_32, true
	case 33:
		return NO_MOVEMENT_33, true
	case 351:
		return NO_MOVEMENT_351, true
	case 352:
		return NO_MOVEMENT_352, true
	case 364:
		return NoExitHeatRay, true
	default:
		return 0, false
	}
}
