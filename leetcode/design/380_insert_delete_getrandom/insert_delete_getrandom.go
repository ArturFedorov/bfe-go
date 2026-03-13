package insert_delete_getrandom

type RandomizedSet struct{}

func Constructor() RandomizedSet              { return RandomizedSet{} }
func (rs *RandomizedSet) Insert(val int) bool { return false }
func (rs *RandomizedSet) Remove(val int) bool { return false }
func (rs *RandomizedSet) GetRandom() int      { return 0 }
