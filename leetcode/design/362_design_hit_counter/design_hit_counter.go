package design_hit_counter

type HitCounter struct{}

func Constructor() HitCounter                    { return HitCounter{} }
func (hc *HitCounter) Hit(timestamp int)         {}
func (hc *HitCounter) GetHits(timestamp int) int { return 0 }
