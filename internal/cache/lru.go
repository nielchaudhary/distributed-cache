package cache

import "container/list"

type LRUCache struct {
	capacity int
	cache    map[string]*list.Element //Hash map to store key-value pairs for O(1) lookups.
	list     *list.List
}

type entry struct {
	key   string
	value interface{}
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
