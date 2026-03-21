package insert_delete_getrandom

import "math/rand"

type RandomizedSet struct {
	arr []int
	set map[int]int
}

func Constructor() RandomizedSet {
	return RandomizedSet{
		arr: make([]int, 0),
		set: make(map[int]int),
	}
}
func (rs *RandomizedSet) Insert(val int) bool {
	_, ok := rs.set[val]
	if ok {
		return false
	}

	rs.arr = append(rs.arr, val)
	rs.set[val] = len(rs.arr) - 1

	return true
}
func (rs *RandomizedSet) Remove(val int) bool {
	index, ok := rs.set[val]
	if !ok {
		return false
	}

	lastElement := rs.arr[len(rs.arr)-1]
	rs.arr[index] = lastElement

	if val != lastElement {
		rs.set[lastElement] = index
	}

	rs.arr = rs.arr[:len(rs.arr)-1]
	delete(rs.set, val)

	return true
}
func (rs *RandomizedSet) GetRandom() int {
	index := rand.Intn(len(rs.arr))
	return rs.arr[index]
}
