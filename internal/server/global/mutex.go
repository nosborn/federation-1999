package global

import (
	"sync"
)

var mu sync.Mutex

func Lock() {
	mu.Lock()
}

func Unlock() {
	mu.Unlock()
}
