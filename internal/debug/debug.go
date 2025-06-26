package debug

import (
	"log"
)

const (
	check uint = 1 << iota
	precondition
	trace
)

var flags uint

func Check(condition bool) {
	if !isEnabled(check) {
		return
	}
	if condition {
		return
	}
	log.Panic("debug.Check: condition is false")
}

func EnableCheck() {
	flags |= check
}

func EnablePrecondition() {
	flags |= precondition
}

func EnableTrace() {
	flags |= trace
}

func Flags() uint {
	return flags
}

func Precondition(condition bool) {
	if !isEnabled(precondition) {
		return
	}
	if condition {
		return
	}
	log.Panic("debug.Precondition: condition is false")
}

func Trace(format string, v ...any) {
	if !isEnabled(trace) {
		return
	}
	log.Printf(format, v...)
}

func isEnabled(flag uint) bool {
	return (flags & flag) != 0
}
