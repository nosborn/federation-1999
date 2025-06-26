package server

// import (
// 	"testing"
// 	"time"
//
// 	"github.com/nosborn/federation-1999/internal/server/global"
// )
//
// func TestTransportationCreation(t *testing.T) {
// 	// Create a test duchy
// 	duchy := NewCoreDuchy("TestDuchy", 10)
//
// 	// Check that transportation was created
// 	if duchy.Transportation == nil {
// 		t.Fatal("Transportation should be created with duchy")
// 	}
//
// 	// Check transportation name
// 	expected := "TestDuchy Transportation"
// 	if duchy.Transportation.name != expected {
// 		t.Errorf("Expected transportation name %q, got %q", expected, duchy.Transportation.name)
// 	}
//
// 	// Check that free slots are initialized (99 slots)
// 	expectedSlots := TRANS_MAX_JOBS - 1 // slot 0 is reserved
// 	if len(duchy.Transportation.freeSlots) != expectedSlots {
// 		t.Errorf("Expected %d free slots, got %d", expectedSlots, len(duchy.Transportation.freeSlots))
// 	}
//
// 	// Check that all job slots are invalid initially
// 	for i, slot := range duchy.Transportation.jobArray {
// 		if slot.job.jobType != JOB_INVALID {
// 			t.Errorf("Job slot %d should be invalid, got type %d", i, slot.job.jobType)
// 		}
// 	}
// }
//
// func TestTransportationSolSpecialCase(t *testing.T) {
// 	// Sol duchy should get special name
// 	solDuchy := NewCoreDuchy("Sol", 25)
//
// 	expected := "Transportation Central"
// 	if solDuchy.Transportation.name != expected {
// 		t.Errorf("Expected Sol transportation name %q, got %q", expected, solDuchy.Transportation.name)
// 	}
// }
//
// func TestTransportationStartStop(t *testing.T) {
// 	duchy := NewCoreDuchy("TestDuchy", 10)
// 	transport := duchy.Transportation
//
// 	// Initially should not have a ticker
// 	if transport.ticker != nil {
// 		t.Error("Transportation should not have ticker before Start()")
// 	}
//
// 	// Start transportation
// 	transport.Start()
//
// 	// Should now have a ticker
// 	if transport.ticker == nil {
// 		t.Error("Transportation should have ticker after Start()")
// 	}
//
// 	// Stop transportation
// 	transport.Stop()
//
// 	// Ticker should still exist but be stopped
// 	if transport.ticker == nil {
// 		t.Error("Transportation should still have ticker after Stop() (but stopped)")
// 	}
// }
//
// func TestTransportationTimerLoop(t *testing.T) {
// 	duchy := NewCoreDuchy("TestDuchy", 10)
// 	transport := duchy.Transportation
//
// 	// Start the transportation system
// 	transport.Start()
// 	defer transport.Stop()
//
// 	// Wait for at least one timer tick (30 seconds is too long for tests,
// 	// but we can verify the timer is running by checking it exists)
// 	time.Sleep(10 * time.Millisecond)
//
// 	// Verify timer is still running
// 	if transport.ticker == nil {
// 		t.Error("Transportation timer should still be running")
// 	}
//
// 	// The timer proc should be safe to call directly
// 	// (this tests the timer handler without waiting 30 seconds)
// 	transport.timerProc()
//
// 	// Should not crash and timer should still be running after timerProc
// 	if transport.ticker == nil {
// 		t.Error("Transportation timer should still be running after timerProc")
// 	}
// }
//
// func TestDuchyTransportationIntegration(t *testing.T) {
// 	// Create duchy
// 	duchy := NewCoreDuchy("IntegrationTest", 15)
//
// 	// Transportation should be automatically created
// 	if duchy.Transportation == nil {
// 		t.Fatal("Duchy should have transportation system")
// 	}
//
// 	// Test that StartTransportation method works
// 	duchy.StartTransportation()
//
// 	// Transportation should now be running
// 	if duchy.Transportation.ticker == nil {
// 		t.Error("Transportation should be running after StartTransportation()")
// 	}
//
// 	// Clean up
// 	duchy.Transportation.Stop()
//
// 	// Test that duchy properly references the transportation
// 	if duchy.Transportation.duchy != duchy {
// 		t.Error("Transportation should reference its parent duchy")
// 	}
// }
//
// func TestSolJobGeneration(t *testing.T) {
// 	// Initialize global haulers count
// 	global.Haulers = 10
//
// 	// Create Sol duchy
// 	solDuchy := NewCoreDuchy("Sol", 25)
// 	transport := solDuchy.Transportation
//
// 	// Verify initial state - all slots should be free
// 	expectedFreeSlots := TRANS_MAX_JOBS - 1 // slot 0 is reserved
// 	if len(transport.freeSlots) != expectedFreeSlots {
// 		t.Errorf("Expected %d free slots initially, got %d", expectedFreeSlots, len(transport.freeSlots))
// 	}
//
// 	// Generate Sol jobs
// 	transport.generateSolJobs()
//
// 	// Should have some jobs allocated now
// 	jobsCreated := expectedFreeSlots - len(transport.freeSlots)
// 	if jobsCreated == 0 {
// 		t.Error("Expected some jobs to be created, but none were allocated")
// 	}
//
// 	t.Logf("Created %d jobs for %d haulers", jobsCreated, global.Haulers)
//
// 	// Check that allocated jobs have correct properties
// 	for i, slot := range transport.jobArray {
// 		if slot.job.jobType == JOB_INVALID {
// 			continue // Skip free slots
// 		}
//
// 		// Verify job type
// 		if slot.job.jobType != JOB_MILKRUN {
// 			t.Errorf("Job %d should be MILKRUN type, got %d", i, slot.job.jobType)
// 		}
//
// 		// Verify planets are different
// 		if slot.job.from == slot.job.to {
// 			t.Errorf("Job %d has same origin and destination: %s", i, slot.job.from)
// 		}
//
// 		// Verify planets are valid Sol planets
// 		validPlanets := map[string]bool{
// 			"Castillo": true, "Titan": true, "Mars": true, "Earth": true,
// 			"Moon": true, "Venus": true, "Mercury": true,
// 		}
// 		if !validPlanets[slot.job.from] {
// 			t.Errorf("Job %d has invalid origin planet: %s", i, slot.job.from)
// 		}
// 		if !validPlanets[slot.job.to] {
// 			t.Errorf("Job %d has invalid destination planet: %s", i, slot.job.to)
// 		}
//
// 		// Verify cargo properties
// 		if slot.job.pallet.Quantity != 75 {
// 			t.Errorf("Job %d should have quantity 75, got %d", i, slot.job.pallet.Quantity)
// 		}
// 		if slot.job.pallet.Origin != slot.job.from {
// 			t.Errorf("Job %d cargo origin should match job origin", i)
// 		}
//
// 		// Verify other job properties
// 		if slot.job.status != JOB_NONE {
// 			t.Errorf("Job %d should have status JOB_NONE, got %d", i, slot.job.status)
// 		}
// 		if slot.job.credits != 2 {
// 			t.Errorf("Job %d should have 2 credits, got %d", i, slot.job.credits)
// 		}
// 		if slot.job.gtu < 9 || slot.job.gtu > 12 {
// 			t.Errorf("Job %d GTU should be 9-12, got %d", i, slot.job.gtu)
// 		}
// 		if slot.job.value <= 0 {
// 			t.Errorf("Job %d should have positive value, got %d", i, slot.job.value)
// 		}
//
// 		t.Logf("Job %d: %s->%s, %s qty=%d, value=%d IG/ton, gtu=%d",
// 			i+1, slot.job.from, slot.job.to,
// 			GoodsArray[slot.job.pallet.GoodsType].name,
// 			slot.job.pallet.Quantity, slot.job.value, slot.job.gtu)
// 	}
// }
//
// func TestJobSlotAllocation(t *testing.T) {
// 	transport := NewTransportation(NewCoreDuchy("Test", 10))
//
// 	// Create a test job
// 	job := Work{
// 		jobType: JOB_MILKRUN,
// 		from:    "Earth",
// 		to:      "Mars",
// 		status:  JOB_NONE,
// 		pallet: Cargo{
// 			GoodsType: commodityGold,
// 			Quantity:  50,
// 			Origin:    "Earth",
// 			Cost:      100,
// 		},
// 		value:   10,
// 		gtu:     10,
// 		credits: 1,
// 	}
//
// 	// Should be able to allocate
// 	if !transport.allocateJobSlot(&job) {
// 		t.Error("Should be able to allocate first job slot")
// 	}
//
// 	// Should have one less free slot
// 	expectedFreeSlots := TRANS_MAX_JOBS - 2 // one allocated
// 	if len(transport.freeSlots) != expectedFreeSlots {
// 		t.Errorf("Expected %d free slots after allocation, got %d", expectedFreeSlots, len(transport.freeSlots))
// 	}
//
// 	// Find the allocated job
// 	found := false
// 	for _, slot := range transport.jobArray {
// 		if slot.job.jobType == JOB_MILKRUN && slot.job.from == "Earth" && slot.job.to == "Mars" {
// 			found = true
// 			break
// 		}
// 	}
// 	if !found {
// 		t.Error("Allocated job should be found in job array")
// 	}
// }
//
// func TestSolTransportationTimerIntegration(t *testing.T) {
// 	// Set up haulers for job generation
// 	global.Haulers = 5
//
// 	// Create Sol duchy (should get special treatment)
// 	solDuchy := NewCoreDuchy("Sol", 25)
// 	transport := solDuchy.Transportation
//
// 	// Verify it's empty initially
// 	initialJobs := 0
// 	for _, slot := range transport.jobArray {
// 		if slot.job.jobType != JOB_INVALID {
// 			initialJobs++
// 		}
// 	}
// 	if initialJobs != 0 {
// 		t.Errorf("Expected 0 initial jobs, got %d", initialJobs)
// 	}
//
// 	// Call timerProc (simulates the 30-second timer)
// 	transport.timerProc()
//
// 	// Should now have some jobs
// 	currentJobs := 0
// 	for _, slot := range transport.jobArray {
// 		if slot.job.jobType != JOB_INVALID {
// 			currentJobs++
// 		}
// 	}
//
// 	if currentJobs == 0 {
// 		t.Error("Expected jobs to be generated after timerProc, but found none")
// 	}
//
// 	t.Logf("Timer proc generated %d jobs", currentJobs)
// }
