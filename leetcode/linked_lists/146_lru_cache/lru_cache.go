package lru_cache

import (
	"container/list"
)

type entry struct {
	key   int
	value int
}

type LRUCache struct {
	capacity int
	store    map[int]*list.Element
	list     *list.List
}

const maxCapacity = 3000
const minCapacity = 1

// Constructor initializes the LRU cache with positive size capacity.
func Constructor(capacity int) LRUCache {
	if capacity < minCapacity || capacity > maxCapacity {
		panic("capacity out of range")
	}

	return LRUCache{
		capacity: capacity,
		store:    make(map[int]*list.Element, capacity),
		list:     list.New(),
	}
}

// Get returns the value of the key if it exists, otherwise returns -1.
func (c *LRUCache) Get(key int) int {
	elem, ok := c.store[key]
	if !ok {
		return -1
	}

	c.list.MoveToFront(elem)
	return elem.Value.(*entry).value
}

// Put updates the value of the key if it exists. Otherwise, adds the key-value
// pair to the cache. If the number of keys exceeds capacity, evicts the least
// recently used key.
func (c *LRUCache) Put(key int, value int) {
	if elem, ok := c.store[key]; ok {
		elem.Value.(*entry).value = value
		c.list.MoveToFront(elem)
		return
	}

	if len(c.store) >= c.capacity {
		back := c.list.Back()
		if back == nil {
			return
		}

		evicted := back.Value.(*entry)
		delete(c.store, evicted.key)
		c.list.Remove(back)
	}

	elem := c.list.PushFront(&entry{key, value})
	c.store[key] = elem
}
