package group_anagrams

import (
	"reflect"
	"sort"
	"testing"
)

func sortResult(groups [][]string) [][]string {
	for _, group := range groups {
		sort.Strings(group)
	}
	sort.Slice(groups, func(i, j int) bool {
		if len(groups[i]) != len(groups[j]) {
			return len(groups[i]) < len(groups[j])
		}
		for k := 0; k < len(groups[i]); k++ {
			if groups[i][k] != groups[j][k] {
				return groups[i][k] < groups[j][k]
			}
		}
		return false
	})
	return groups
}

func TestGroupAnagrams(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		got := sortResult(groupAnagrams([]string{"eat", "tea", "tan", "ate", "nat", "bat"}))
		want := sortResult([][]string{{"ate", "eat", "tea"}, {"bat"}, {"nat", "tan"}})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example2", func(t *testing.T) {
		got := sortResult(groupAnagrams([]string{""}))
		want := sortResult([][]string{{""}})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Example3", func(t *testing.T) {
		got := sortResult(groupAnagrams([]string{"a"}))
		want := sortResult([][]string{{"a"}})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("AllSame", func(t *testing.T) {
		got := sortResult(groupAnagrams([]string{"abc", "abc", "abc"}))
		want := sortResult([][]string{{"abc", "abc", "abc"}})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("NoAnagrams", func(t *testing.T) {
		got := sortResult(groupAnagrams([]string{"abc", "def", "ghi"}))
		want := sortResult([][]string{{"abc"}, {"def"}, {"ghi"}})
		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
