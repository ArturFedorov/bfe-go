package next_closest_time

import "testing"

func TestNextClosestTime(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := nextClosestTime("19:34")
		if got != "19:39" {
			t.Errorf("got %q, want %q", got, "19:39")
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := nextClosestTime("23:59")
		if got != "22:22" {
			t.Errorf("got %q, want %q", got, "22:22")
		}
	})

	t.Run("AllSameDigits", func(t *testing.T) {
		got := nextClosestTime("11:11")
		if got != "11:11" {
			t.Errorf("got %q, want %q", got, "11:11")
		}
	})

	t.Run("Midnight", func(t *testing.T) {
		got := nextClosestTime("00:00")
		if got != "00:00" {
			t.Errorf("got %q, want %q", got, "00:00")
		}
	})

	t.Run("WrapToNextDay", func(t *testing.T) {
		got := nextClosestTime("13:55")
		if got != "15:11" {
			t.Errorf("got %q, want %q", got, "15:11")
		}
	})
}
