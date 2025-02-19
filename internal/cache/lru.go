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

// to get the string from cache, move that element to most recently used in cache if found, else return false
func (l *LRUCache) Get(key string) (interface{}, bool) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false

}

func (l *LRUCache) Put(key string, value interface{}) {
	if elem, ok := l.cache[key]; ok {
		//if found, move it to front and update the current value.
		l.list.MoveToFront(elem)
		elem.Value.(*entry).value = value
	}

	if l.list.Len() >= l.capacity {
		oldest := l.list.Back()
		delete(l.cache, oldest.Value.(*entry).key)
		l.list.Remove(oldest)
	}

	elem := l.list.PushFront(&entry{key, value})
	l.cache[key] = elem
}
