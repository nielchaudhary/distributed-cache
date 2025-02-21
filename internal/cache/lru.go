package cache

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
}

type entry struct {
	key   string
	value interface{}
}
type Cache interface {
	Get(key string) (interface{}, bool)
	Put(key string, value interface{})
}

func NewLRUCache(capacity int) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (l *LRUCache) Get(key string) (interface{}, bool) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

func (l *LRUCache) Put(key string, value interface{}) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
		return
	}

	if l.list.Len() >= l.capacity {
		oldest := l.list.Back()
		delete(l.cache, oldest.Value.(*entry).key)
		l.list.Remove(oldest)
	}

	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}
