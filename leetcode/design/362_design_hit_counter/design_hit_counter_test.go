package design_hit_counter

import "testing"

func TestHitCounter(t *testing.T) {
	t.Run("hit at 1,2,3 getHits(4) -> 3", func(t *testing.T) {
		hc := Constructor()
		hc.Hit(1)
		hc.Hit(2)
		hc.Hit(3)
		if got := hc.GetHits(4); got != 3 {
			t.Errorf("GetHits(4) = %d, want 3", got)
		}
	})

	t.Run("getHits(300) -> 3", func(t *testing.T) {
		hc := Constructor()
		hc.Hit(1)
		hc.Hit(2)
		hc.Hit(3)
		if got := hc.GetHits(300); got != 3 {
			t.Errorf("GetHits(300) = %d, want 3", got)
		}
	})

	t.Run("getHits(301) -> 2", func(t *testing.T) {
		hc := Constructor()
		hc.Hit(1)
		hc.Hit(2)
		hc.Hit(3)
		if got := hc.GetHits(301); got != 2 {
			t.Errorf("GetHits(301) = %d, want 2", got)
		}
	})

	t.Run("no hits -> 0", func(t *testing.T) {
		hc := Constructor()
		if got := hc.GetHits(1); got != 0 {
			t.Errorf("GetHits(1) = %d, want 0", got)
		}
	})
}
