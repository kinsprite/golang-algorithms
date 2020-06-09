package store

import (
	"container/list"
)

type LRUElement struct {
	Key   string
	Value interface{}
}

type LRUCache struct {
	Max  int
	Map  map[string]*list.Element
	List *list.List
}

func LRUNew(maxSize int) *LRUCache {
	return &LRUCache{maxSize, map[string]*list.Element{}, list.New()}
}

func (cache *LRUCache) Len() int {
	return cache.List.Len()
}

func (cache *LRUCache) Get(key string) (interface{}, bool) {
	if e, ok := cache.Map[key]; ok {
		cache.List.MoveToBack(e)
		return e.Value.(*LRUElement).Value, ok
	}

	return nil, false
}

func (cache *LRUCache) Set(key string, value interface{}) {
	l := cache.List

	if l.Len() >= cache.Max {
		e := l.Front()
		l.Remove(e)

		delete(cache.Map, e.Value.(*LRUElement).Key)
	}

	e := &LRUElement{key, value}
	le := l.PushBack(e)
	cache.Map[key] = le
}
