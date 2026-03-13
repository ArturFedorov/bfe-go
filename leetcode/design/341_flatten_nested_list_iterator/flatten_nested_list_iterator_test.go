package flatten_nested_list_iterator

import "testing"

func intPtr(v int) *int { return &v }

func collectAll(it *NestedIterator) []int {
	var result []int
	for it.HasNext() {
		result = append(result, it.Next())
	}
	return result
}

func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func TestNestedIterator(t *testing.T) {
	t.Run("[[1,1],2,[1,1]] -> [1,1,2,1,1]", func(t *testing.T) {
		input := []*NestedInteger{
			{list: []*NestedInteger{{val: intPtr(1)}, {val: intPtr(1)}}},
			{val: intPtr(2)},
			{list: []*NestedInteger{{val: intPtr(1)}, {val: intPtr(1)}}},
		}
		it := Constructor(input)
		got := collectAll(it)
		want := []int{1, 1, 2, 1, 1}
		if !equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("[1,[4,[6]]] -> [1,4,6]", func(t *testing.T) {
		input := []*NestedInteger{
			{val: intPtr(1)},
			{list: []*NestedInteger{
				{val: intPtr(4)},
				{list: []*NestedInteger{{val: intPtr(6)}}},
			}},
		}
		it := Constructor(input)
		got := collectAll(it)
		want := []int{1, 4, 6}
		if !equal(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("empty list", func(t *testing.T) {
		input := []*NestedInteger{}
		it := Constructor(input)
		got := collectAll(it)
		if len(got) != 0 {
			t.Errorf("got %v, want empty", got)
		}
	})
}
