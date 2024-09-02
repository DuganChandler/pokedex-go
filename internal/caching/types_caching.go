package caching

import (
	"sync"
	"time"
)

type CacheEntry struct {
	createdAt time.Time // time  created
	val       []byte    // raw data cache
}

type Cache struct {
	cacheData map[string]CacheEntry
	mu        *sync.Mutex
}

func NewCache(interval time.Duration) Cache {
	ca := Cache{
		cacheData: make(map[string]CacheEntry),
		mu:        &sync.Mutex{},
	}

	go ca.reapLoop(interval)

	return ca
}

func (ca *Cache) Add(key string, val []byte) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	ca.cacheData[key] = CacheEntry{
		createdAt: time.Now().UTC(),
		val:       val,
	}
}

func (ca *Cache) Get(key string) ([]byte, bool) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	entry, ok := ca.cacheData[key]
	return entry.val, ok

}

func (ca *Cache) reapLoop(interval time.Duration) {
	ticker := time.NewTicker(interval)
	for range ticker.C {
		ca.reap(time.Now().UTC(), interval)
	}
}

func (ca *Cache) reap(now time.Time, last time.Duration) {
	ca.mu.Lock()
	defer ca.mu.Unlock()

	for key, val := range ca.cacheData {
		if val.createdAt.Before(now.Add(-last)) {
			delete(ca.cacheData, key)
		}
	}
}
