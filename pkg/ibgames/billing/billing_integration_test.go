package billing

import (
	"fmt"
	"testing"
	"time"

	"github.com/nosborn/federation-1999/pkg/ibgames"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

// Integration tests - skip with: go test -short
// These tests use real time passage and take several minutes to complete

func TestRealTimeBilling(t *testing.T) {
	if testing.Short() {
		t.Skip("Skipping integration test in short mode")
	}

	setup := setupTestDB(t)
	t.Cleanup(resetBillingState)

	t.Logf("Using temporary database: %s", setup.FilePath)

	err := Init(Federation)
	require.NoError(t, err)

	t.Run("paying account accrues and charges minutes over real time", func(t *testing.T) {
		FreePeriod(false)
		AutoCommit(true)

		uid := ibgames.AccountID(666200)
		initialMinutes := 1000
		setup.CreateTestAccount(t, uid, fmt.Sprintf("user%d", uid), "N", initialMinutes)

		session, err := BeginSession(uid, "192.0.2.1")
		require.NoError(t, err)

		// Verify initial state
		minutes := getAccountMinutes(t, setup, uid)
		assert.Equal(t, initialMinutes-minimumCharge, minutes)
		assert.Equal(t, minimumCharge, session.lastCharge)
		assert.True(t, session.ticking)

		t.Logf("Sleeping for 130 seconds to cross 2+ minute boundaries...")
		time.Sleep(130 * time.Second) // 2+ minutes to ensure charging

		// Call Tick to process accumulated time
		result := session.Tick()
		assert.Equal(t, 1, result)

		// Should have accumulated roughly 130 seconds (allow ±10 seconds margin)
		minExpected := int64(120)
		maxExpected := int64(140)
		assert.True(t, session.seconds >= minExpected && session.seconds <= maxExpected,
			"Expected between %d and %d, got %d", minExpected, maxExpected, session.seconds)

		// Should have charged for at least 2 minutes
		expectedMinutes := int(session.seconds / 60)
		assert.GreaterOrEqual(t, expectedMinutes, 2,
			"Expected at least 2 minutes charged, got %d", expectedMinutes)

		t.Logf("Accumulated %d seconds, charged %d minutes", session.seconds, expectedMinutes)
	})

	t.Run("stop and start clock controls billing over minute boundaries", func(t *testing.T) {
		FreePeriod(false)

		uid := ibgames.AccountID(666201)
		setup.CreateTestAccount(t, uid, fmt.Sprintf("user%d", uid), "N", 1000)

		session, err := BeginSession(uid, "192.0.2.1")
		require.NoError(t, err)

		t.Logf("Running for 70 seconds with ticking=true...")
		time.Sleep(70 * time.Second) // Cross 1 minute boundary
		session.Tick()

		timeAfterFirst := session.seconds
		chargeAfterFirst := session.lastCharge
		t.Logf("After 70s: %v accumulated, %d minutes charged", timeAfterFirst, chargeAfterFirst)

		// Should have accumulated ~70s and charged for 1+ minute
		assert.GreaterOrEqual(t, timeAfterFirst, int64(60), "Should have at least 60 seconds")
		assert.GreaterOrEqual(t, chargeAfterFirst, 1, "Should have charged at least 1 minute")

		// Stop the clock
		session.StopClock()
		assert.False(t, session.ticking)

		t.Logf("Running for 80 seconds with ticking=false...")
		time.Sleep(80 * time.Second) // Would cross another minute if ticking
		session.Tick()

		// Time and charges should not have advanced during stopped period
		timeAfterStop := session.seconds
		chargeAfterStop := session.lastCharge
		t.Logf("After stop: %v accumulated, %d minutes charged", timeAfterStop, chargeAfterStop)

		assert.Equal(t, timeAfterFirst, timeAfterStop, "Time should not advance when stopped")
		assert.Equal(t, chargeAfterFirst, chargeAfterStop, "Charges should not advance when stopped")

		// Restart the clock
		session.StartClock()
		assert.True(t, session.ticking)

		t.Logf("Running for 70 seconds with ticking=true again...")
		time.Sleep(70 * time.Second) // Cross another minute boundary
		session.Tick()

		finalTime := session.seconds
		finalCharge := session.lastCharge
		t.Logf("Final: %v accumulated, %d minutes charged", finalTime, finalCharge)

		// Should have accumulated roughly another 70 seconds
		additionalTime := finalTime - timeAfterStop
		additionalCharge := finalCharge - chargeAfterStop
		assert.GreaterOrEqual(t, additionalTime, int64(60),
			"Should have accumulated at least another 60 seconds, got %d", additionalTime)
		assert.GreaterOrEqual(t, additionalCharge, 1,
			"Should have charged at least 1 more minute, got %d", additionalCharge)
	})

	t.Run("complimentary account never accrues charges even over minutes", func(t *testing.T) {
		t.Logf("Starting complimentary account test...")

		uid := ibgames.AccountID(666202)
		initialMinutes := 1000
		setup.CreateTestAccount(t, uid, fmt.Sprintf("user%d", uid), "Y", initialMinutes)
		t.Logf("Created complimentary test account")

		session, err := BeginSession(uid, "192.0.2.1")
		require.NoError(t, err)
		t.Logf("BeginSession completed successfully")

		// Verify no initial charge
		minutes := getAccountMinutes(t, setup, uid)
		assert.Equal(t, initialMinutes, minutes)
		assert.Equal(t, 0, session.lastCharge)
		t.Logf("Initial verification completed - minutes: %d, lastCharge: %d", minutes, session.lastCharge)

		t.Logf("About to sleep for 90 seconds with complimentary account...")
		time.Sleep(90 * time.Second) // Cross minute boundary
		t.Logf("Sleep completed, about to call Tick()...")

		session.Tick()
		t.Logf("Tick() completed")

		// Should still have no time accumulated and no charges
		assert.Equal(t, int64(0), session.seconds)
		assert.Equal(t, 0, session.lastCharge)
		t.Logf("Time accumulation check completed - seconds: %d, lastCharge: %d", session.seconds, session.lastCharge)

		// Account should be unchanged
		t.Logf("About to call final getAccountMinutes()...")
		finalMinutes := getAccountMinutes(t, setup, uid)
		t.Logf("Final getAccountMinutes() completed - result: %d", finalMinutes)

		assert.Equal(t, initialMinutes, finalMinutes)

		t.Logf("Complimentary account test completed successfully")
	})
}
