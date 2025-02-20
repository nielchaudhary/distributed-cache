package cache

import "container/list"

type LFUCache struct {
	capacity int
	cache    map[string]*list.Element
	list     *list.List
}

type en struct {
	key   string
	value interface{}
}

func newLFUCache(capacity int) *LFUCache {
	return &LFUCache{
		capacity: capacity,
		cache:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

func (l *LFUCache) Get(key string) (interface{}, bool) {
	if elem, ok := l.cache[key]; ok {
		l.list.MoveToFront(elem)
		return elem.Value.(*en).value, true

		//elem.Value            → gets the en struct: {key: "user1", value: "John"}
		//elem.Value.(*en)      → type asserts it's an 'en' struct
		//elem.Value.(*en).value → gets the value field ("John")s
		//Returns: "John", true
	}

	return nil, false
}

func (l *LFUCache) Put(key string, value interface{}) {
	if elem, ok := l.cache[key]; ok {
		l.list.PushFront(elem)
		elem.Value.(*en).value = value
	}

	if l.list.Len() > l.capacity {
		oldest := l.list.Back()
		delete(l.cache, oldest.Value.(*en).key)
		l.list.Remove(oldest)

	}

	elem := l.list.PushFront(&en{key, value})
	l.cache[key] = elem

}
