package cache

import "container/heap"

type lfuItem struct {
	key      string
	value    interface{}
	priority int
	index    int
}

type lfuQueue []*lfuItem

type LFUCache struct {
	capacity int
	hashMap  map[string]*lfuItem
	queue    *lfuQueue
}

func (fq lfuQueue) Len() int { return len(fq) }

func (fq lfuQueue) Less(i, j int) bool {
	return fq[i].priority < fq[j].priority
}

func (fq lfuQueue) Swap(i, j int) {
	fq[i], fq[j] = fq[j], fq[i]
	fq[i].index = i
	fq[j].index = j
}

func (fq *lfuQueue) Push(x interface{}) {
	n := len(*fq)
	item := x.(*lfuItem)
	item.index = n
	*fq = append(*fq, item)
}

func (fq *lfuQueue) Pop() interface{} {
	old := *fq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*fq = old[0 : n-1]
	return item
}

// update modifies the priority and value of an Item in the queue.
func (fq *lfuQueue) update(item *lfuItem, value interface{}, priority int) {
	item.value = value
	item.priority = priority
	heap.Fix(fq, item.index)
}

func LFUNew(capacity int) *LFUCache {
	return &LFUCache{capacity, map[string]*lfuItem{}, &lfuQueue{}}
}

func (cache *LFUCache) Len() int {
	return len(*cache.queue)
}

func (cache *LFUCache) Get(key string) (interface{}, bool) {
	if item, ok := cache.hashMap[key]; ok {
		cache.queue.update(item, item.value, item.priority+1)
		return item.value, ok
	}

	return nil, false
}

func (cache *LFUCache) Set(key string, value interface{}) {
	if item, ok := cache.hashMap[key]; ok {
		cache.queue.update(item, value, item.priority+1)
		return
	}

	if cache.Len() >= cache.capacity {
		item := heap.Pop(cache.queue).(*lfuItem)
		delete(cache.hashMap, item.key)
	}

	item := &lfuItem{
		key, value, 1, 0,
	}

	heap.Push(cache.queue, item)
	cache.hashMap[key] = item
}
