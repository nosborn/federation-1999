package server

// import (
// 	"testing"
//
// 	"github.com/nosborn/federation-1999/internal/model"
// 	"github.com/nosborn/federation-1999/internal/text"
// 	"github.com/stretchr/testify/require"
// )
//
// func TestCmdInsureQuote(t *testing.T) {
// 	tests := []struct {
// 		name        string
// 		setupPlayer func() *Player
// 		expectedMsg string
// 	}{
// 		{
// 			name: "quote for player with 0 deaths",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 0
// 				p.Balance = 10000
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureQuotation, "1,000"),
// 		},
// 		{
// 			name: "quote for player with 1 death",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 1
// 				p.Balance = 10000
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureQuotation, "4,000"),
// 		},
// 		{
// 			name: "quote for player with 5 deaths",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 5
// 				p.Balance = 20000
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureQuotation, "16,000"),
// 		},
// 		{
// 			name: "quote for player with 9 deaths (boundary case)",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 9
// 				p.Balance = 50000
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureQuotation, "28,000"),
// 		},
// 		{
// 			name: "quote for player with 10 deaths (higher rate)",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 10
// 				p.Balance = 100000
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureQuotation, "510,000"),
// 		},
// 		{
// 			name: "quote for player with 15 deaths (higher rate)",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 15
// 				p.Balance = 1000000
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureQuotation, "760,000"),
// 		},
// 		{
// 			name: "quote at wrong location",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 0
// 				p.Balance = 10000
// 				p.curLoc = &Location{Flags: 0} // No model.LfIns flag
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureWrongLocation),
// 		},
// 		{
// 			name: "quote when already insured",
// 			setupPlayer: func() *Player {
// 				p := createTestPlayer()
// 				p.Deaths = 0
// 				p.Balance = 13000
// 				p.Flags0 = model.PL0_INSURED
// 				p.curLoc = &Location{Flags: model.LfIns}
// 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// 				return p
// 			},
// 			expectedMsg: text.Msg(text.InsureAlreadyInsured),
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			player := tt.setupPlayer()
// 			player.Insure(InsureGetQuote)
// 			assertMessageEquals(t, player, tt.expectedMsg)
// 		})
// 	}
// }
//
// // func TestInsureBuyPolicy(t *testing.T) {
// // 	tests := []struct {
// // 		name            string
// // 		setupPlayer     func() *Player
// // 		expectedMsg     string
// // 		expectedFlags   uint32
// // 		expectedBalance int
// // 		expectInsured   bool
// // 	}{
// // 		{
// // 			name: "buy policy with 0 deaths",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 0
// // 				p.Balance = 13000
// // 				p.curLoc = &Location{Flags: model.LfIns}
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureOK),
// // 			expectedFlags:   model.PL0_INSURED,
// // 			expectedBalance: 12000, // 13000 - 1000
// // 			expectInsured:   true,
// // 		},
// // 		{
// // 			name: "buy policy with 5 deaths",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 5
// // 				p.Balance = 20000
// // 				p.curLoc = &Location{Flags: model.LfIns}
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureOK),
// // 			expectedFlags:   model.PL0_INSURED,
// // 			expectedBalance: 4000, // 20000 - 16000
// // 			expectInsured:   true,
// // 		},
// // 		{
// // 			name: "buy policy with 10 deaths (higher rate)",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 10
// // 				p.Balance = 600000
// // 				p.curLoc = &Location{Flags: model.LfIns}
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureOK),
// // 			expectedFlags:   model.PL0_INSURED,
// // 			expectedBalance: 90000, // 600000 - 510000
// // 			expectInsured:   true,
// // 		},
// // 		{
// // 			name: "insufficient funds",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 0
// // 				p.Balance = 500 // Less than 1000 needed
// // 				p.curLoc = &Location{Flags: model.LfIns}
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureInsufficientFunds, "1,000"),
// // 			expectedFlags:   0,
// // 			expectedBalance: 500, // Unchanged
// // 			expectInsured:   false,
// // 		},
// // 		{
// // 			name: "insufficient funds for high death count",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 15
// // 				p.Balance = 500000 // Less than 760000 needed
// // 				p.curLoc = &Location{Flags: model.LfIns}
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureInsufficientFunds, "760,000"),
// // 			expectedFlags:   0,
// // 			expectedBalance: 500000, // Unchanged
// // 			expectInsured:   false,
// // 		},
// // 		{
// // 			name: "already insured",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 0
// // 				p.Balance = 13000
// // 				p.Flags0 = model.PL0_INSURED
// // 				p.curLoc = &Location{Flags: model.LfIns}
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureAlreadyInsured),
// // 			expectedFlags:   model.PL0_INSURED,
// // 			expectedBalance: 13000, // Unchanged
// // 			expectInsured:   true,
// // 		},
// // 		{
// // 			name: "wrong location",
// // 			setupPlayer: func() *Player {
// // 				p := createTestPlayer()
// // 				p.Deaths = 0
// // 				p.Balance = 10000
// // 				p.curLoc = &Location{Flags: 0} // No model.LfIns flag
// // 				p.curSys = &System{name: "Test", loadState: SystemOnline}
// // 				return p
// // 			},
// // 			expectedMsg:     text.Msg(text.InsureWrongLocation),
// // 			expectedFlags:   0,
// // 			expectedBalance: 10000, // Unchanged
// // 			expectInsured:   false,
// // 		},
// // 	}
// //
// // 	for _, tt := range tests {
// // 		t.Run(tt.name, func(t *testing.T) {
// // 			player := tt.setupPlayer()
// // 			initialBalance := player.Balance
// //
// // 			player.Insure(InsureBuyPolicy)
// //
// // 			// Check message output
// // 			assertMessageEquals(t, player, tt.expectedMsg)
// //
// // 			// Check flags
// // 			if tt.expectInsured {
// // 				testutil.AssertFlagSet(t, player.Flags0, model.PL0_INSURED)
// // 			} else {
// // 				testutil.AssertFlagNotSet(t, player.Flags0, model.PL0_INSURED)
// // 			}
// //
// // 			// Check balance
// // 			require.Equal(t, tt.expectedBalance, player.Balance, "Balance should be updated correctly")
// //
// // 			// Check insurance status
// // 			require.Equal(t, tt.expectInsured, player.IsInsured(), "Insurance status should match expected")
// //
// // 			// For successful purchases, verify cost calculation
// // 			if tt.expectInsured && tt.expectedBalance != initialBalance {
// // 				expectedCost := initialBalance - tt.expectedBalance
// // 				var actualCost int
// // 				if player.Deaths < 10 {
// // 					actualCost = 1000 + (3000 * player.Deaths)
// // 				} else {
// // 					actualCost = 10000 + (50000 * player.Deaths)
// // 				}
// // 				require.Equal(t, expectedCost, actualCost, "Cost calculation should match expected formula")
// // 			}
// // 		})
// // 	}
// // }
//
// func TestInsuranceCostCalculation(t *testing.T) {
// 	tests := []struct {
// 		name         string
// 		deaths       int
// 		expectedCost int
// 	}{
// 		{
// 			name:         "0 deaths",
// 			deaths:       0,
// 			expectedCost: 1000, // 1000 + (3000 * 0)
// 		},
// 		{
// 			name:         "1 death",
// 			deaths:       1,
// 			expectedCost: 4000, // 1000 + (3000 * 1)
// 		},
// 		{
// 			name:         "5 deaths",
// 			deaths:       5,
// 			expectedCost: 16000, // 1000 + (3000 * 5)
// 		},
// 		{
// 			name:         "9 deaths (boundary)",
// 			deaths:       9,
// 			expectedCost: 28000, // 1000 + (3000 * 9)
// 		},
// 		{
// 			name:         "10 deaths (higher rate starts)",
// 			deaths:       10,
// 			expectedCost: 510000, // 10000 + (50000 * 10)
// 		},
// 		{
// 			name:         "15 deaths (higher rate)",
// 			deaths:       15,
// 			expectedCost: 760000, // 10000 + (50000 * 15)
// 		},
// 		{
// 			name:         "20 deaths (higher rate)",
// 			deaths:       20,
// 			expectedCost: 1010000, // 10000 + (50000 * 20)
// 		},
// 	}
//
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			var actualCost int
// 			if tt.deaths < 10 {
// 				actualCost = 1000 + (3000 * tt.deaths)
// 			} else {
// 				actualCost = 10000 + (50000 * tt.deaths)
// 			}
// 			require.Equal(t, tt.expectedCost, actualCost, "Cost calculation should match expected formula")
// 		})
// 	}
// }
//
// func TestInsureEdgeCases(t *testing.T) {
// 	// t.Run("exact balance match", func(t *testing.T) {
// 	// 	player := createTestPlayer()
// 	// 	player.Deaths = 0
// 	// 	player.Balance = 1000 // Exactly enough for insurance
// 	// 	player.curLoc = &Location{Flags: model.LfIns}
// 	// 	player.curSys = &System{name: "Test", loadState: SystemOnline}
// 	//
// 	// 	player.Insure(InsureBuyPolicy)
// 	//
// 	// 	assertMessageEquals(t, player, text.Msg(text.InsureOK))
// 	// 	require.Equal(t, 0, player.Balance)
// 	// 	require.True(t, player.IsInsured())
// 	// })
//
// 	t.Run("one credit short", func(t *testing.T) {
// 		player := createTestPlayer()
// 		player.Deaths = 0
// 		player.Balance = 999 // One credit short
// 		player.curLoc = &Location{Flags: model.LfIns}
// 		player.curSys = &System{name: "Test", loadState: SystemOnline}
//
// 		player.Insure(InsureBuyPolicy)
//
// 		assertMessageEquals(t, player, text.Msg(text.InsureInsufficientFunds, "1,000"))
// 		require.Equal(t, 999, player.Balance)
// 		require.False(t, player.IsInsured())
// 	})
//
// 	t.Run("flying spaceship location check", func(t *testing.T) {
// 		player := createTestPlayer()
// 		player.Deaths = 0
// 		player.Balance = 13000
// 		player.Flags0 = model.PL0_FLYING
// 		player.ShipLoc = 100
// 		player.curLoc = &Location{Flags: 0} // No insurance broker
// 		player.curSys = &System{name: "Test", loadState: SystemOnline}
//
// 		player.Insure(InsureGetQuote)
//
// 		assertMessageEquals(t, player, text.Msg(text.InsureWrongLocation))
// 	})
// }
