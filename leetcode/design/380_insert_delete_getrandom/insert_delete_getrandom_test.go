package insert_delete_getrandom

import "testing"

func TestRandomizedSet(t *testing.T) {
	t.Run("insert returns true", func(t *testing.T) {
		rs := Constructor()
		if got := rs.Insert(1); !got {
			t.Errorf("Insert(1) = %v, want true", got)
		}
	})

	t.Run("duplicate insert returns false", func(t *testing.T) {
		rs := Constructor()
		rs.Insert(1)
		if got := rs.Insert(1); got {
			t.Errorf("Insert(1) duplicate = %v, want false", got)
		}
	})

	t.Run("remove returns true", func(t *testing.T) {
		rs := Constructor()
		rs.Insert(1)
		if got := rs.Remove(1); !got {
			t.Errorf("Remove(1) = %v, want true", got)
		}
	})

	t.Run("remove missing returns false", func(t *testing.T) {
		rs := Constructor()
		if got := rs.Remove(1); got {
			t.Errorf("Remove(1) missing = %v, want false", got)
		}
	})

	t.Run("getRandom returns inserted element", func(t *testing.T) {
		rs := Constructor()
		rs.Insert(42)
		if got := rs.GetRandom(); got != 42 {
			t.Errorf("GetRandom() = %d, want 42", got)
		}
	})
}
