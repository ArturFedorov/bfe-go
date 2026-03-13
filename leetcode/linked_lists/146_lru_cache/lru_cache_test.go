package lru_cache

import (
	"testing"
)

func TestLRUCache(t *testing.T) {
	t.Run("Example1", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		assertGet(t, &cache, 1, 1)
		cache.Put(3, 3)
		assertGet(t, &cache, 2, -1)
		cache.Put(4, 4)
		assertGet(t, &cache, 1, -1)
		assertGet(t, &cache, 3, 3)
		assertGet(t, &cache, 4, 4)
	})

	t.Run("GetMissing", func(t *testing.T) {
		cache := Constructor(2)
		assertGet(t, &cache, 1, -1)
	})

	t.Run("CapacityOne", func(t *testing.T) {
		cache := Constructor(1)
		cache.Put(1, 10)
		assertGet(t, &cache, 1, 10)
		cache.Put(2, 20)
		assertGet(t, &cache, 1, -1)
		assertGet(t, &cache, 2, 20)
	})

	t.Run("UpdateExistingKey", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(1, 10)
		assertGet(t, &cache, 1, 10)
		assertGet(t, &cache, 2, 2)
	})

	t.Run("UpdateRefreshesOrder", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(1, 1)
		cache.Put(3, 3)
		assertGet(t, &cache, 1, 1)
		assertGet(t, &cache, 2, -1)
		assertGet(t, &cache, 3, 3)
	})

	t.Run("GetRefreshesOrder", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Get(1)
		cache.Put(3, 3)
		assertGet(t, &cache, 1, 1)
		assertGet(t, &cache, 2, -1)
		assertGet(t, &cache, 3, 3)
	})

	t.Run("EvictionOrder", func(t *testing.T) {
		cache := Constructor(3)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)
		cache.Put(4, 4)
		assertGet(t, &cache, 1, -1)
		assertGet(t, &cache, 2, 2)
		assertGet(t, &cache, 3, 3)
		assertGet(t, &cache, 4, 4)
	})

	t.Run("EvictionAfterGet", func(t *testing.T) {
		cache := Constructor(3)
		cache.Put(1, 1)
		cache.Put(2, 2)
		cache.Put(3, 3)
		cache.Get(1)
		cache.Put(4, 4)
		assertGet(t, &cache, 1, 1)
		assertGet(t, &cache, 2, -1)
		assertGet(t, &cache, 3, 3)
		assertGet(t, &cache, 4, 4)
	})

	t.Run("RepeatedPutSameKey", func(t *testing.T) {
		cache := Constructor(1)
		cache.Put(1, 1)
		cache.Put(1, 2)
		cache.Put(1, 3)
		assertGet(t, &cache, 1, 3)
	})

	t.Run("PutGetInterleavedSequence", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(2, 1)
		cache.Put(1, 1)
		cache.Put(2, 3)
		cache.Put(4, 1)
		assertGet(t, &cache, 1, -1)
		assertGet(t, &cache, 2, 3)
		assertGet(t, &cache, 4, 1)
	})

	t.Run("LargeCapacity", func(t *testing.T) {
		cache := Constructor(100)
		for i := 0; i < 100; i++ {
			cache.Put(i, i*10)
		}
		for i := 0; i < 100; i++ {
			assertGet(t, &cache, i, i*10)
		}
		cache.Put(100, 1000)
		assertGet(t, &cache, 0, -1)
		assertGet(t, &cache, 100, 1000)
	})

	t.Run("ZeroValueKey", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(0, 42)
		assertGet(t, &cache, 0, 42)
		cache.Put(1, 1)
		cache.Put(2, 2)
		assertGet(t, &cache, 0, -1)
	})

	t.Run("ZeroValue", func(t *testing.T) {
		cache := Constructor(2)
		cache.Put(1, 0)
		assertGet(t, &cache, 1, 0)
	})
}

func assertGet(t *testing.T, cache *LRUCache, key int, expected int) {
	t.Helper()
	got := cache.Get(key)
	if got != expected {
		t.Errorf("Get(%d) = %d, want %d", key, got, expected)
	}
}
