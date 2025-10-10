package main

import (
	"container/list"
)

type KeyType = string
type ValueType = string

type CacheEntry struct {
	Key KeyType
	Val ValueType
}

type LRU struct {
	maxSize  int                       // cache is limited in size
	items    map[KeyType]*list.Element // map items store pointers to Elements in tracking list; fast lookup by key
	tracking *list.List                // tracking list Elements hold instances of *CacheEntry; fast lookup by who was least recently used
}

func NewLRU(size int) *LRU {
	return &LRU{
		maxSize:  size,
		items:    make(map[KeyType]*list.Element),
		tracking: list.New(),
	}
}

func (lru *LRU) Get(key KeyType) *ValueType {
	element, ok := lru.items[key]
	if !ok {
		return nil
	}
	lru.tracking.MoveToFront(element)
	entry := element.Value.(*CacheEntry) // list element holds instance of *CacheEntry
	return &entry.Val
}

func (lru *LRU) Put(key KeyType, val ValueType) {
	// such key already exists
	if el, ok := lru.items[key]; ok {
		el.Value.(*CacheEntry).Val = val
		lru.tracking.MoveToFront(el)
		return
	}

	// cache size limit reached; we need to kick out someone first
	if len(lru.items) == lru.maxSize {
		lruElement := lru.tracking.Back()          // find least recently used element
		lruEntry := lruElement.Value.(*CacheEntry) // list element holds instance of *CacheEntry
		delete(lru.items, lruEntry.Key)            // remove it from the item map
		lru.tracking.Remove(lruElement)            // remove it from the tracking list
	}

	newEntry := &CacheEntry{Key: key, Val: val}
	newElement := lru.tracking.PushFront(newEntry) // insert new element into the tracking list
	lru.items[key] = newElement                    // insert new element into the items map
}

func (lru *LRU) Size() int {
	return len(lru.items)
}
