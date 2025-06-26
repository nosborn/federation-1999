package server

import (
	"fmt"
	"log"
	"math/rand/v2"
	"os/exec"
	"path/filepath"
	"sync"
	"time"

	"github.com/nosborn/federation-1999/internal/config"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/server/database"
	"github.com/nosborn/federation-1999/internal/server/global"
	"github.com/nosborn/federation-1999/internal/text"
)

type SystemLoadState int

const (
	SystemOffline SystemLoadState = iota
	SystemLoading
	SystemOnline
	SystemUnloading
)

// Loader manages the queue of systems waiting to be loaded
type Loader struct {
	// autostart       bool
	checker *exec.Cmd
	// checkerPid    int // FIXME
	duchiesLoaded bool
	frozen        bool
	idleTimer     *time.Timer
	queue         []*PlayerSystem
	// running         bool
	startListener func()
	timer         *time.Timer
}

var (
	loader     *Loader
	loaderOnce sync.Once
)

const (
	timerPeriod = 300 * time.Second
)

func Dequeue(system *PlayerSystem) {
	GetLoader().Dequeue(system)
}

func Enqueue(system *PlayerSystem) {
	GetLoader().Enqueue(system)
}

func GetLoader() *Loader {
	return loader
}

func NewLoader(startListener func()) *Loader {
	loaderOnce.Do(func() {
		loader = &Loader{
			queue:         make([]*PlayerSystem, 0),
			startListener: startListener,
		}
	})
	return loader
}

func LoaderQueueLength() int {
	return GetLoader().QueueLength()
}

func LoaderQueuePosition(system *PlayerSystem) int {
	return GetLoader().QueuePosition(system)
}

func StartLoader() {
	GetLoader().Start()
}

func StopLoader() {
	GetLoader().Stop()
}

func (l *Loader) Dequeue(system *PlayerSystem) {
	if system == nil {
		log.Panic("loader.Dequeue: system == nil")
		// log.Print("PANIC: loader.Dequeue: system == nil")
	}

	for i, s := range l.queue {
		if s == system {
			l.queue = append(l.queue[:i], l.queue[i+1:]...)
		}
	}
}

func (l *Loader) Enqueue(system *PlayerSystem) {
	if system == nil {
		log.Panic("loader.Enqueue: system == nil")
		// log.Print("PANIC: loader.Enqueue: system == nil")
	}
	if !system.IsLoading() {
		log.Panic("loader.Enqueue: !system.IsLoading()")
		// log.Print("PANIC: loader.Enqueue: !system.IsLoading()")
	}

	l.queue = append(l.queue, system)
	// if l.checkerPid == 0 && l.idleTimer == nil && l.timer == nil { -- FIXME
	// 	log.Print("loader.Enqueue: starting timer")
	// 	l.timer = time.AfterFunc(timerPeriod, l.timerHandler)
	// }
	if l.checker == nil && l.idleTimer == nil { // FIXME
		// log.Printf("loader.Enqueue: requesting idle call")
		l.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), l.idleHandler)
	}
}

func (l *Loader) IsFrozen() bool {
	return l.frozen
}

// QueueLength returns the current queue length
// Caller must hold global lock
func (l *Loader) QueueLength() int {
	return len(l.queue)
}

// QueuePosition returns the position of a system in the load queue (1-based)
// Returns 0 if the system is not in the queue
// Caller must hold global lock
func (l *Loader) QueuePosition(system *PlayerSystem) int {
	for i, queued := range l.queue {
		if queued == system {
			return i + 1
		}
	}
	return 0
}

func (l *Loader) Start() {
	if l.idleTimer == nil {
		// log.Printf("loader.Start: requesting idle call")
		l.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), l.idleHandler)
	}
}

func (l *Loader) Stop() {
	if l.idleTimer != nil {
		l.idleTimer.Stop()
		l.idleTimer = nil
	}
	if l.timer != nil {
		l.timer.Stop()
		l.timer = nil
	}
}

func (l *Loader) fail(system *PlayerSystem) {
	if system == nil {
		log.Panic("loader.fail: system == nil")
		// log.Print("PANIC: loader.fail: system == nil")
	}
	if !system.IsLoading() {
		log.Panic("loader.fail: !system.IsLoading()")
		// log.Print("PANIC: loader.fail: !system.IsLoading()")
	}

	system.SetLoadState(SystemOffline)

	// if system.Owner() != nil {
	// 	if system.Owner.IsPlaying() {
	// 		system.Owner().Outputm(text.PlanetCheckFailed)
	// 		system.Owner().FlushOutput()
	// 	}
	// 	if system.Owner().Rank() < RankSquire {
	// 		system.Owner().ownSystem = nil
	// 		system.ClearOwner()
	// 	} else {
	// 		system.Owner().Save(database.SaveNow)
	// 	}
	// }

	// if system.Owner() == nil {
	// 	// delete system
	// }
}

func (l *Loader) idleHandler() {
	global.Lock()
	defer global.Unlock()

	defer database.CommitDatabase()
	l.idleProc()
}

func (l *Loader) idleProc() {
	// log.Printf("Loader idle proc")

	if l.checker != nil {
		log.Panic("loader.idleProc: l.checker != nil")
		// log.Print("PANIC: loader.idleProc: l.checker != nil")
	}
	if l.timer != nil {
		log.Panic("loader.idleProc: l.timer != nil")
		// log.Print("PANIC: loader.idleProc: l.timer != nil")
	}
	l.idleTimer = nil

	//
	if !l.duchiesLoaded {
		if len(l.queue) == 0 || !l.queue[0].IsCapital() {
			l.duchiesLoaded = true
			l.startListener()
		}
	}

	//
	if len(l.queue) == 0 {
		return
	}

	//
	l.frozen = l.shouldFreeze()
	if l.frozen {
		if l.timer == nil {
			// log.Printf("idleProc: starting timer")
			l.timer = time.AfterFunc(timerPeriod, l.timerHandler)
		}
		return
	}

	//
	system := l.queue[0]
	if !system.IsLoading() {
		log.Panic("loader.idleProc: !system.IsLoading()")
		// log.Print("PANIC: loader.idleProc: !system.IsLoading()")
	}
	l.queue = l.queue[1:]

	log.Printf("Running planet checker for %s [%d]", system.Name(), system.Owner().UID())

	args := []string{"-c", fmt.Sprintf("%d", system.Owner().UID())}
	noExchange := false // TODO
	if noExchange {
		args[0] = "-cn"
	}
	l.checker = exec.Command(filepath.Join(config.BinDir, "workbench"), args...) //nolint:gosec,noctx

	// Start the command asynchronously. This is like fork/exec.
	if err := l.checker.Start(); err != nil {
		log.Printf("failed to start workbench for %s: %v", system.Name(), err)
		l.fail(system)
		l.checker = nil
		if l.idleTimer == nil {
			// log.Printf("loader.idleProc: requesting idle call")
			l.timer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), l.idleHandler)
		}
		return
	}

	// This goroutine is the "reaper". It waits for the child to exit.
	go func(cmd *exec.Cmd, system *PlayerSystem) {
		// Wait for the command to finish. This is like waiting for SIGCHLD.
		err := cmd.Wait()

		// Now, do the work that the SIGCHLD handler would do.
		// This must be under the global lock.
		global.Lock()
		defer global.Unlock()

		if err == nil {
			// Workbench exited successfully.
			system.Load()
			if system.Owner().IsPlaying() {
				system.Owner().Outputm(text.PlanetLoadComplete)
				system.Owner().FlushOutput()
			}
			if system.Owner().Rank() == model.RankExplorer {
				system.Owner().PromoteToSquire()
			}
		} else {
			// Workbench failed.
			log.Printf("Planet checker failed for %s [%d]: %v", system.Name(), system.Owner().UID(), err)
			l.fail(system)
		}

		// Cleanup and schedule the next idle check.
		l.checker = nil
		if l.idleTimer == nil {
			// log.Printf("loader.reaper: requesting idle call")
			l.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), l.idleHandler)
		}
	}(l.checker, system)
}

// // processLoop is the main loader goroutine that processes the queue
// func (l *Loader) processLoop() {
// 	ticker := time.NewTicker(100 * time.Millisecond) // FIXME: 1 * time.Second
// 	defer ticker.Stop()
//
// 	for {
// 		select {
// 		case <-ticker.C:
// 			l.processNextSystem()
// 		}
// 	}
// }

// // processNextSystem processes the next system in the queue
// func (l *Loader) processNextSystem() {
// 	// Acquire global lock for the entire operation
// 	global.Lock()
// 	defer global.Unlock()
//
// 	if len(l.queue) == 0 {
// 		return
// 	}
//
// 	// Get the first system in queue
// 	system := l.queue[0]
//
// 	if !l.duchiesLoaded {
// 		if !system.IsCapital() {
// 			l.duchiesLoaded = true
// 			go l.startListener()
// 		}
// 	}
//
// 	// Orchestrate the loading process
// 	success := l.orchestrateLoad(system)
//
// 	if success {
// 		// Remove from queue - system is now responsible for its own state
// 		l.queue = l.queue[1:]
// 	} else {
// 		l.queue = l.queue[1:]
// 		l.fail(system)
// 	}
//
// 	if len(l.queue) == 0 {
// 		if !l.duchiesLoaded {
// 			l.duchiesLoaded = true
// 			go l.startListener()
// 		}
// 		l.Stop()
// 		return
// 	}
//
// 	//nolint:gosec // "It's Just A Game"
// 	time.Sleep(time.Duration(rand.IntN(500)+1) * time.Millisecond) // emulate idle processing
// }

// // orchestrateLoad manages the loading process for a system
// func (l *Loader) orchestrateLoad(system *PlayerSystem) bool {
// 	// TODO: In full implementation:
// 	// 1. Run workbench --check to validate planet files
// 	// 2. If validation passes, call system.LoadPlayerSystem()
// 	log.Printf("Running planet checker for %s [%d]", system.Name(), system.Owner().UID())
//
// 	// For now, simulate validation and immediately call system to load
// 	// In reality, this would check planet files first
// 	return system.LoadPlayerSystem()
// }

func (l *Loader) shouldFreeze() bool {
	if global.TestFeaturesEnabled && l.duchiesLoaded && global.CurrentPlayers == 0 {
		return true
	}
	return false
}

func (l *Loader) timerHandler() {
	global.Lock()
	defer global.Unlock()

	defer database.CommitDatabase() // Not needed
	l.timerProc()
}

func (l *Loader) timerProc() {
	// log.Printf("loader timer proc")

	if l.checker != nil {
		log.Panic("loader.timerProc: l.checker != nil")
		// log.Print("PANIC: loader.timerProc: l.checker != nil")
	}
	if l.idleTimer != nil {
		log.Panic("loader.timerProc: l.idleTimer != nil")
		// log.Print("PANIC: loader.timerProc: l.idleTimer != nil")
	}
	l.timer = nil

	if !l.duchiesLoaded {
		if len(l.queue) == 0 || !l.queue[0].IsCapital() {
			l.duchiesLoaded = true
			l.startListener()
		}
	}

	//
	if len(l.queue) == 0 {
		return
	}

	//
	if l.frozen = l.shouldFreeze(); l.frozen {
		// log.Printf("loader.timerProc: starting timer")
		l.timer = time.AfterFunc(timerPeriod, l.timerHandler)
		return
	}

	//
	// log.Printf("loader.timerProc: requesting idle call")
	l.idleTimer = time.AfterFunc(time.Duration(rand.IntN(int(time.Second))), l.idleHandler)
}
