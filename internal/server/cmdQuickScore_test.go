package server

import (
	"fmt"
	"testing"

	"github.com/dustin/go-humanize"
	"github.com/nosborn/federation-1999/internal/model"
	"github.com/nosborn/federation-1999/internal/testutil"
)

func TestCmdQuickScore(t *testing.T) {
	tests := []struct {
		name        string
		player      *Player
		expectedMsg string
	}{
		{
			name: "with wallet and with insurance",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankCaptain)
				p.balance = 12345
				p.Flags0 |= model.PL0_INSURED
				p.Sta = PlayerStat{Cur: 10, Max: 20}
				p.Str = PlayerStat{Cur: 15}
				p.Int = PlayerStat{Cur: 25}
				p.Dex = PlayerStat{Cur: 30}
				return p
			}(),
			expectedMsg: fmt.Sprintf("Stats: IG:%s Sta:10/20 Str:15 Int:25 Dex:30 Ins:Y\n", humanize.Comma(12345)),
		},
		{
			name: "with wallet and without insurance",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankTrader)
				p.balance = 12345
				p.Flags0 &^= model.PL0_INSURED
				p.Sta = PlayerStat{Cur: 10, Max: 20}
				p.Str = PlayerStat{Cur: 15}
				p.Int = PlayerStat{Cur: 25}
				p.Dex = PlayerStat{Cur: 30}
				return p
			}(),
			expectedMsg: fmt.Sprintf("Stats: IG:%s Sta:10/20 Str:15 Int:25 Dex:30 Ins:N\n", humanize.Comma(12345)),
		},
		{
			name: "without wallet and with insurance",
			player: func() *Player {
				p := createTestPlayerWithSession(666000, "TestActor", model.RankSenator)
				p.balance = 0
				p.Flags0 |= model.PL0_INSURED
				p.Sta = PlayerStat{Cur: 10, Max: 20}
				p.Str = PlayerStat{Cur: 15}
				p.Int = PlayerStat{Cur: 25}
				p.Dex = PlayerStat{Cur: 30}
				return p
			}(),
			expectedMsg: "Stats: Sta:10/20 Str:15 Int:25 Dex:30 Ins:Y\n",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			saver := &testutil.MockPlayerSaver[*Player]{}
			tt.player.saveFunc = saver.Save

			tt.player.CmdQuickScore()

			assertMessageEquals(t, tt.player, tt.expectedMsg)
			testutil.AssertPlayerNotSaved(t, saver)
		})
	}
}
