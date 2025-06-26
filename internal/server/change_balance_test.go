package server

import (
	"math"
	"testing"

	"github.com/nosborn/federation-1999/internal/model"
)

func TestChangeBalance(t *testing.T) {
	tests := []struct {
		name            string
		initialBalance  int32
		amount          int32
		expectedBalance int32
	}{
		{
			name:            "normal positive addition",
			initialBalance:  1000,
			amount:          500,
			expectedBalance: 1500,
		},
		{
			name:            "normal negative addition",
			initialBalance:  1000,
			amount:          -300,
			expectedBalance: 700,
		},
		{
			name:            "zero amount",
			initialBalance:  1000,
			amount:          0,
			expectedBalance: 1000,
		},
		{
			name:            "positive overflow protection - exact limit",
			initialBalance:  model.MAX_BALANCE - 100,
			amount:          100,
			expectedBalance: model.MAX_BALANCE,
		},
		{
			name:            "positive overflow protection - exceeds limit",
			initialBalance:  model.MAX_BALANCE - 50,
			amount:          100,
			expectedBalance: model.MAX_BALANCE,
		},
		{
			name:            "positive overflow protection - max int32 amount",
			initialBalance:  1000,
			amount:          math.MaxInt32,
			expectedBalance: model.MAX_BALANCE,
		},
		{
			name:            "positive overflow protection - already at max",
			initialBalance:  model.MAX_BALANCE,
			amount:          1,
			expectedBalance: model.MAX_BALANCE,
		},
		{
			name:            "negative underflow protection - exact limit",
			initialBalance:  model.MIN_BALANCE + 100,
			amount:          -100,
			expectedBalance: model.MIN_BALANCE,
		},
		{
			name:            "negative underflow protection - exceeds limit",
			initialBalance:  model.MIN_BALANCE + 50,
			amount:          -100,
			expectedBalance: model.MIN_BALANCE,
		},
		{
			name:            "negative underflow protection - already at min",
			initialBalance:  model.MIN_BALANCE,
			amount:          -1,
			expectedBalance: model.MIN_BALANCE,
		},
		{
			name:            "large positive amount within bounds",
			initialBalance:  0,
			amount:          1000000,
			expectedBalance: 1000000,
		},
		{
			name:            "large negative amount within bounds",
			initialBalance:  0,
			amount:          -1000000,
			expectedBalance: -1000000,
		},
		{
			name:            "near max balance with small addition",
			initialBalance:  model.MAX_BALANCE - 1,
			amount:          1,
			expectedBalance: model.MAX_BALANCE,
		},
		{
			name:            "near min balance with small subtraction",
			initialBalance:  model.MIN_BALANCE + 1,
			amount:          -1,
			expectedBalance: model.MIN_BALANCE,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			balance := tt.initialBalance
			changeBalance(&balance, tt.amount)

			if balance != tt.expectedBalance {
				t.Errorf("changeBalance(&%d, %d) resulted in %d, want %d",
					tt.initialBalance, tt.amount, balance, tt.expectedBalance)
			}

			// Ensure balance is always within valid range
			if balance < model.MIN_BALANCE {
				t.Errorf("balance %d is below MIN_BALANCE %d", balance, model.MIN_BALANCE)
			}
			if balance > model.MAX_BALANCE {
				t.Errorf("balance %d is above MAX_BALANCE %d", balance, model.MAX_BALANCE)
			}
		})
	}
}

func TestChangeBalanceBoundaryConditions(t *testing.T) {
	// Test edge cases around the 32-bit integer boundaries
	t.Run("max int32 initial balance", func(t *testing.T) {
		balance := int32(math.MaxInt32)
		changeBalance(&balance, 1)
		if balance > model.MAX_BALANCE {
			t.Errorf("balance %d exceeds MAX_BALANCE %d", balance, model.MAX_BALANCE)
		}
	})

	t.Run("min int32 initial balance", func(t *testing.T) {
		balance := int32(math.MinInt32)
		changeBalance(&balance, -1)
		if balance < model.MIN_BALANCE {
			t.Errorf("balance %d is below MIN_BALANCE %d", balance, model.MIN_BALANCE)
		}
	})

	t.Run("max int32 amount positive", func(t *testing.T) {
		balance := int32(0)
		changeBalance(&balance, math.MaxInt32)
		if balance < model.MIN_BALANCE || balance > model.MAX_BALANCE {
			t.Errorf("balance %d is outside valid range [%d, %d]",
				balance, model.MIN_BALANCE, model.MAX_BALANCE)
		}
	})

	t.Run("min int32 amount negative", func(t *testing.T) {
		balance := int32(0)
		changeBalance(&balance, math.MinInt32)
		if balance < model.MIN_BALANCE || balance > model.MAX_BALANCE {
			t.Errorf("balance %d is outside valid range [%d, %d]",
				balance, model.MIN_BALANCE, model.MAX_BALANCE)
		}
	})
}
