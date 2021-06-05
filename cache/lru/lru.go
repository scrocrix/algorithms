package lru

import (
	"container/list"
)

// CacheItem represents the item our LRU Cache is caching.
type CacheItem struct {
	Key   string
	Value string
}

// LRU or Least Recently Used Cache (a.k.a. Page Replacement Algorithm) is a caching mechanism
// that allows us to replace a cache that hasn't been used for a long amount of time.
type LRU struct {
	// limit represents the limit of frames we're able to store in memory.
	limit int

	// items represent a Doubly Linked List structure which automatically stores the cache items in it.
	items *list.List

	// cache keeps record of the elements we're caching on runtime.
	cache map[string]*CacheItem
}

// NewLRU creates an instance of LRU cache algorithm.
func NewLRU(limit int) *LRU {
	return &LRU{
		limit: limit,
		items: list.New(),
		cache: map[string]*CacheItem{},
	}
}

// Cache automatically stores an integer into our *list.List object.
func (lru *LRU) Cache(item *CacheItem) {
	if _, ok := lru.cache[item.Key]; ok {
		return
	}

	if len(lru.cache) >= lru.limit {
		leastUsed := lru.items.Back()
		lru.items.Remove(leastUsed)
		delete(lru.cache, leastUsed.Value.(*CacheItem).Key)
	}

	lru.items.PushFront(item)

	lru.cache[item.Key] = item
}

// GetItems return items that are currently cached
func (lru *LRU) GetItems() *list.List {
	return lru.items
}
