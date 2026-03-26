package design_hashmap

import "testing"

func TestMyHashMap(t *testing.T) {
	t.Run("put and get", func(t *testing.T) {
		m := Constructor()
		m.Put(1, 1)
		m.Put(2, 2)
		if got := m.Get(1); got != 1 {
			t.Errorf("Get(1) = %d, want 1", got)
		}
		if got := m.Get(3); got != -1 {
			t.Errorf("Get(3) = %d, want -1", got)
		}
	})

	t.Run("put updates existing key", func(t *testing.T) {
		m := Constructor()
		m.Put(2, 2)
		m.Put(2, 1)
		if got := m.Get(2); got != 1 {
			t.Errorf("Get(2) = %d, want 1", got)
		}
	})

	t.Run("remove key", func(t *testing.T) {
		m := Constructor()
		m.Put(2, 2)
		m.Remove(2)
		if got := m.Get(2); got != -1 {
			t.Errorf("Get(2) = %d, want -1", got)
		}
	})

	t.Run("empty map get returns -1", func(t *testing.T) {
		m := Constructor()
		if got := m.Get(1); got != -1 {
			t.Errorf("Get(1) = %d, want -1", got)
		}
	})
}
