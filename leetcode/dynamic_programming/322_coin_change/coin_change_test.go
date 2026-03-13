package coin_change

import "testing"

func TestCoinChange(t *testing.T) {
	tests := []struct {
		name     string
		coins    []int
		amount   int
		expected int
	}{
		{"example 1", []int{1, 2, 5}, 11, 3},
		{"impossible", []int{2}, 3, -1},
		{"zero amount", []int{1}, 0, 0},
		{"large coin value", []int{1, 2147483647}, 2, 2},
		{"single coin exact", []int{5}, 5, 1},
		{"single coin one", []int{1}, 1, 1},
		{"multiple options", []int{1, 3, 4}, 6, 2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := coinChange(tt.coins, tt.amount)
			if result != tt.expected {
				t.Errorf("coinChange(%v, %d) = %d, want %d", tt.coins, tt.amount, result, tt.expected)
			}
		})
	}
}
