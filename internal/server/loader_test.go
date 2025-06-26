package server

// import (
// 	"testing"
// 	"time"
//
// 	"github.com/nosborn/federation-1999/internal/server/global"
// )
//
// // mockSystem implements LoadableSystem for testing
// type mockSystem struct {
// 	name          string
// 	isCapital     bool
// 	shouldSucceed bool
// 	loadState     LoadState
// }
//
// func (m *mockSystem) Name() string {
// 	return m.name
// }
//
// func (m *mockSystem) LoadPlayerSystem() bool {
// 	if m.shouldSucceed {
// 		m.loadState = StateOnline
// 		return true
// 	}
// 	m.loadState = StateOffline
// 	return false
// }
//
// func (m *mockSystem) SetLoadState(state LoadState) {
// 	m.loadState = state
// }
//
// func (m *mockSystem) IsCapital() bool {
// 	return m.isCapital
// }
//
// func TestLoaderSingleton(t *testing.T) {
// 	loader1 := NewLoader(func() {})
// 	loader2 := GetLoader()
//
// 	if loader1 != loader2 {
// 		t.Error("NewLoader/GetLoader should return the same instance (singleton)")
// 	}
// }
//
// func TestLoaderStartStop(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Should not be running initially
// 	if loader.running {
// 		t.Error("Loader should not be running initially")
// 	}
//
// 	// Start loader
// 	loader.Start()
// 	if !loader.running {
// 		t.Error("Loader should be running after Start()")
// 	}
//
// 	// Starting again should be safe
// 	loader.Start()
// 	if !loader.running {
// 		t.Error("Loader should still be running after second Start()")
// 	}
//
// 	// Stop loader
// 	loader.Stop()
// 	if loader.running {
// 		t.Error("Loader should not be running after Stop()")
// 	}
//
// 	// Stopping again should be safe
// 	loader.Stop()
// 	if loader.running {
// 		t.Error("Loader should still be stopped after second Stop()")
// 	}
// }
//
// func TestLoaderEnqueueSystem(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Create test systems
// 	system1 := &mockSystem{name: "TestSystem1", shouldSucceed: true}
// 	system2 := &mockSystem{name: "TestSystem2", shouldSucceed: true}
//
// 	// Test with global lock as callers would have
// 	global.Lock()
//
// 	// Initially queue should be empty
// 	if length := loader.GetQueueLength(); length != 0 {
// 		t.Errorf("Expected empty queue, got length %d", length)
// 	}
//
// 	// Enqueue first system
// 	loader.EnqueueSystem(system1)
// 	if length := loader.GetQueueLength(); length != 1 {
// 		t.Errorf("Expected queue length 1, got %d", length)
// 	}
//
// 	// Enqueue second system
// 	loader.EnqueueSystem(system2)
// 	if length := loader.GetQueueLength(); length != 2 {
// 		t.Errorf("Expected queue length 2, got %d", length)
// 	}
//
// 	// Enqueue same system again - should be ignored
// 	loader.EnqueueSystem(system1)
// 	if length := loader.GetQueueLength(); length != 2 {
// 		t.Errorf("Expected queue length to remain 2, got %d", length)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderQueuePosition(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Create test systems
// 	system1 := &mockSystem{name: "First", shouldSucceed: true}
// 	system2 := &mockSystem{name: "Second", shouldSucceed: true}
// 	system3 := &mockSystem{name: "Third", shouldSucceed: true}
// 	systemNotInQueue := &mockSystem{name: "NotInQueue", shouldSucceed: true}
//
// 	global.Lock()
//
// 	// Clear queue for clean test
// 	loader.queue = []LoadableSystem{}
//
// 	// System not in queue should return position 0
// 	if pos := loader.GetQueuePosition(systemNotInQueue); pos != 0 {
// 		t.Errorf("Expected position 0 for system not in queue, got %d", pos)
// 	}
//
// 	// Add systems and check positions
// 	loader.EnqueueSystem(system1)
// 	loader.EnqueueSystem(system2)
// 	loader.EnqueueSystem(system3)
//
// 	// Check positions (1-based)
// 	if pos := loader.GetQueuePosition(system1); pos != 1 {
// 		t.Errorf("Expected position 1 for first system, got %d", pos)
// 	}
// 	if pos := loader.GetQueuePosition(system2); pos != 2 {
// 		t.Errorf("Expected position 2 for second system, got %d", pos)
// 	}
// 	if pos := loader.GetQueuePosition(system3); pos != 3 {
// 		t.Errorf("Expected position 3 for third system, got %d", pos)
// 	}
//
// 	// System not in queue should still return 0
// 	if pos := loader.GetQueuePosition(systemNotInQueue); pos != 0 {
// 		t.Errorf("Expected position 0 for system not in queue, got %d", pos)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderProcessing(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Create a test system that succeeds loading
// 	system := &mockSystem{
// 		name:          "ProcessingTest",
// 		shouldSucceed: true,
// 		loadState:     StateLoading,
// 	}
//
// 	global.Lock()
//
// 	// Clear queue and add test system
// 	loader.queue = []LoadableSystem{}
// 	loader.EnqueueSystem(system)
//
// 	// Verify system is in queue
// 	if length := loader.GetQueueLength(); length != 1 {
// 		t.Fatalf("Expected 1 system in queue, got %d", length)
// 	}
//
// 	global.Unlock()
//
// 	// Process the system directly (simulating timer tick)
// 	loader.processNextSystem()
//
// 	global.Lock()
//
// 	// System should be removed from queue after processing
// 	if length := loader.GetQueueLength(); length != 0 {
// 		t.Errorf("Expected empty queue after processing, got length %d", length)
// 	}
//
// 	// System should be marked as online
// 	if system.loadState != StateOnline {
// 		t.Errorf("Expected system to be online after processing, got %v", system.loadState)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderProcessingEmpty(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	global.Lock()
//
// 	// Clear queue
// 	loader.queue = []LoadableSystem{}
//
// 	global.Unlock()
//
// 	// Processing empty queue should not panic
// 	loader.processNextSystem()
//
// 	global.Lock()
//
// 	// Queue should still be empty
// 	if length := loader.GetQueueLength(); length != 0 {
// 		t.Errorf("Expected queue to remain empty, got length %d", length)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderOrchestration(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Create test system that succeeds
// 	system := &mockSystem{
// 		name:          "OrchestrationTest",
// 		shouldSucceed: true,
// 		loadState:     StateLoading,
// 	}
//
// 	// Test orchestration directly
// 	success := loader.orchestrateLoad(system)
//
// 	if !success {
// 		t.Error("Expected orchestration to succeed")
// 	}
//
// 	// System should be marked as online
// 	if system.loadState != StateOnline {
// 		t.Errorf("Expected system to be online after orchestration, got %v", system.loadState)
// 	}
// }
//
// func TestLoaderIntegration(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Start the loader
// 	loader.Start()
// 	defer loader.Stop()
//
// 	// Create test systems
// 	system1 := &mockSystem{name: "Integration1", shouldSucceed: true, loadState: StateLoading}
// 	system2 := &mockSystem{name: "Integration2", shouldSucceed: true, loadState: StateLoading}
//
// 	global.Lock()
//
// 	// Clear queue
// 	loader.queue = []LoadableSystem{}
//
// 	// Enqueue systems
// 	loader.EnqueueSystem(system1)
// 	loader.EnqueueSystem(system2)
//
// 	// Verify both systems are queued
// 	if length := loader.GetQueueLength(); length != 2 {
// 		t.Fatalf("Expected 2 systems in queue, got %d", length)
// 	}
//
// 	global.Unlock()
//
// 	// Wait a bit for processing
// 	time.Sleep(100 * time.Millisecond)
//
// 	// Process manually to avoid waiting for timer
// 	loader.processNextSystem()
// 	loader.processNextSystem()
//
// 	global.Lock()
//
// 	// Queue should be empty
// 	if length := loader.GetQueueLength(); length != 0 {
// 		t.Errorf("Expected empty queue after processing, got length %d", length)
// 	}
//
// 	// Both systems should be online
// 	if system1.loadState != StateOnline {
// 		t.Errorf("Expected system1 to be online, got %v", system1.loadState)
// 	}
// 	if system2.loadState != StateOnline {
// 		t.Errorf("Expected system2 to be online, got %v", system2.loadState)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderGlobalFunctions(t *testing.T) {
// 	// Test global convenience functions
// 	system := &mockSystem{name: "GlobalTest", shouldSucceed: true}
//
// 	// Initialize loader first
// 	NewLoader(func() {})
//
// 	// These should not panic and should work with the singleton
// 	Start()
// 	defer Stop()
//
// 	global.Lock()
//
// 	// Clear queue
// 	GetLoader().queue = []LoadableSystem{}
//
// 	Enqueue(system)
// 	position := QueuePosition(system)
//
// 	if position != 1 {
// 		t.Errorf("Expected position 1, got %d", position)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderQueueOrdering(t *testing.T) {
// 	loader := NewLoader(func() {})
//
// 	// Create multiple systems
// 	systems := []LoadableSystem{
// 		&mockSystem{name: "First", shouldSucceed: true},
// 		&mockSystem{name: "Second", shouldSucceed: true},
// 		&mockSystem{name: "Third", shouldSucceed: true},
// 		&mockSystem{name: "Fourth", shouldSucceed: true},
// 		&mockSystem{name: "Fifth", shouldSucceed: true},
// 	}
//
// 	global.Lock()
//
// 	// Clear queue
// 	loader.queue = []LoadableSystem{}
//
// 	// Enqueue in order
// 	for _, system := range systems {
// 		loader.EnqueueSystem(system)
// 	}
//
// 	// Verify correct ordering
// 	for i, expectedSystem := range systems {
// 		expectedPos := i + 1
// 		actualPos := loader.GetQueuePosition(expectedSystem)
// 		if actualPos != expectedPos {
// 			t.Errorf("Expected system %s at position %d, got %d",
// 				expectedSystem.Name(), expectedPos, actualPos)
// 		}
// 	}
//
// 	global.Unlock()
//
// 	// Process first system
// 	loader.processNextSystem()
//
// 	global.Lock()
//
// 	// Verify positions shifted correctly
// 	for i, expectedSystem := range systems[1:] { // Skip first system
// 		expectedPos := i + 1 // Should now be at positions 1,2,3,4
// 		actualPos := loader.GetQueuePosition(expectedSystem)
// 		if actualPos != expectedPos {
// 			t.Errorf("After processing, expected system %s at position %d, got %d",
// 				expectedSystem.Name(), expectedPos, actualPos)
// 		}
// 	}
//
// 	// First system should no longer be in queue
// 	if pos := loader.GetQueuePosition(systems[0]); pos != 0 {
// 		t.Errorf("Expected processed system to not be in queue, got position %d", pos)
// 	}
//
// 	global.Unlock()
// }
//
// func TestLoaderDuchiesLoadedCallback(t *testing.T) {
// 	callbackCalled := false
// 	loader := NewLoader(func() {
// 		callbackCalled = true
// 	})
//
// 	// Create a non-capital system (should trigger duchiesLoaded callback)
// 	nonCapitalSystem := &mockSystem{
// 		name:          "NonCapital",
// 		isCapital:     false,
// 		shouldSucceed: true,
// 		loadState:     StateLoading,
// 	}
//
// 	global.Lock()
// 	loader.queue = []LoadableSystem{}
// 	loader.duchiesLoaded = false // Ensure we start with duchies not loaded
// 	loader.EnqueueSystem(nonCapitalSystem)
// 	global.Unlock()
//
// 	// Process the system - should trigger callback
// 	loader.processNextSystem()
//
// 	// Give the goroutine time to execute
// 	time.Sleep(50 * time.Millisecond)
//
// 	if !callbackCalled {
// 		t.Error("Expected startListener callback to be called when processing non-capital system")
// 	}
//
// 	if !loader.duchiesLoaded {
// 		t.Error("Expected duchiesLoaded to be true after processing non-capital system")
// 	}
// }
//
// func TestLoaderCapitalSystemNoCallback(t *testing.T) {
// 	callbackCalled := false
// 	loader := NewLoader(func() {
// 		callbackCalled = true
// 	})
//
// 	// Create a capital system (should NOT trigger duchiesLoaded callback)
// 	capitalSystem := &mockSystem{
// 		name:          "Capital",
// 		isCapital:     true,
// 		shouldSucceed: true,
// 		loadState:     StateLoading,
// 	}
//
// 	global.Lock()
// 	loader.queue = []LoadableSystem{}
// 	loader.duchiesLoaded = false // Ensure we start with duchies not loaded
// 	loader.EnqueueSystem(capitalSystem)
// 	global.Unlock()
//
// 	// Process the system - should NOT trigger callback yet
// 	loader.processNextSystem()
//
// 	// Give any potential goroutine time to execute
// 	time.Sleep(50 * time.Millisecond)
//
// 	if callbackCalled {
// 		t.Error("Expected startListener callback NOT to be called when processing capital system")
// 	}
//
// 	if loader.duchiesLoaded {
// 		t.Error("Expected duchiesLoaded to remain false after processing capital system")
// 	}
// }
