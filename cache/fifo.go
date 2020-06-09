package cache

import (
	"container/list"
)

type FIFOCache struct {
	capacity int
	list     *list.List
	hashMap  map[string]*list.Element
}

type fifoElement struct {
	key   string
	value interface{}
}

func FIFONew(capacity int) *FIFOCache {
	return &FIFOCache{
		capacity,
		list.New(),
		map[string]*list.Element{},
	}
}

func (cache *FIFOCache) Len() int {
	return cache.list.Len()
}

func (cache *FIFOCache) Get(key string) (interface{}, bool) {
	if v, ok := cache.hashMap[key]; ok {
		return v.Value.(*fifoElement).value, ok
	}
	return nil, false
}

func (cache *FIFOCache) Set(key string, value interface{}) {
	if e, ok := cache.hashMap[key]; ok {
		e.Value.(*fifoElement).value = value
		return
	}

	if cache.list.Len() >= cache.capacity {
		e := cache.list.Front()
		cache.list.Remove(e)

		delete(cache.hashMap, e.Value.(*fifoElement).key)
	}

	e := cache.list.PushBack(&fifoElement{key, value})
	cache.hashMap[key] = e
}
