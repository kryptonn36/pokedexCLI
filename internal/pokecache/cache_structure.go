package pokecache

import (
	"sync"
	"time"
)

type Cache struct{
	structure   map[string]cacheEntry
	sync.Mutex
	interval    time.Duration
}


type cacheEntry struct{
	createdAt   time.Time
	val         []byte
}

func NewCache(interval time.Duration) *Cache{
	var c = &Cache{
		structure: make(map[string]cacheEntry),
		// mu       : &sync.Mutex{},
		interval:  interval,
	}
	go c.reapLoop()
	return c
}

func (cache *Cache) Add(key string, val []byte){
	cache.Lock()

	cache.structure[key] = cacheEntry{
		createdAt: time.Now(),
		val : val,
	}
	// entrypoint.val = val
	// cache.structure[key] = entrypoint
	// cache.structure.createdAt = time.Now()
	cache.Unlock()
}

func (cache *Cache) Get(key string) ([]byte, bool){
	val,ok:= cache.structure[key]
	if !ok {
		return nil,false
	}
	return val.val,true
}

func (cache *Cache) reapLoop(){
	ticker := time.NewTicker(cache.interval)
	defer ticker.Stop()

	for range ticker.C{
		cache.Lock()
		for key, entry := range cache.structure{
			if time.Since(entry.createdAt) > cache.interval{
				delete(cache.structure,key)
			}
		}
		cache.Unlock()
	}
}