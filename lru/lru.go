package main

import "time"

type KeyType string
type ValueType = string
type Value struct {
	Val  ValueType
	date time.Time
}

type LRU struct {
	maxSize int
	items   map[KeyType]Value
}

func NewLRU(size int) *LRU {
	return &LRU{maxSize: size, items: make(map[KeyType]Value)}
}

func (lru *LRU) Get(key KeyType) *ValueType {
	val, ok := lru.items[key]
	if !ok {
		return nil
	}
	newVal := Value{Val: val.Val, date: val.date}
	lru.items[key] = newVal
	return &val.Val
}

func (lru *LRU) Put(key KeyType, val ValueType) {
	if len(lru.items) == lru.maxSize {
		lru.DiscarLRU()
	}
	newVal := Value{Val: val, date: time.Now()}
	lru.items[key] = newVal
}

func (lru *LRU) Size() int {
	return len(lru.items)
}

func (lru *LRU) DiscarLRU() {
	discardCandidateKey := lru.FindDiscardCandiate()
	if discardCandidateKey != nil {
		delete(lru.items, *discardCandidateKey)
	}
}

func (lru *LRU) FindDiscardCandiate() *KeyType {
	oldest := time.Now()
	var key *KeyType
	for k, v := range lru.items {
		if v.date.Before(oldest) {
			oldest = v.date
			key = &k
		}
	}
	return key
}
