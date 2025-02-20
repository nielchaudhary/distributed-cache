package rpc

type CacheService struct {
	cache Cache
}

func (c *CacheService) Get(key string, reply *interface{}) error {
	value, ok := c.cache.Get(key)
	if !ok {
		return nil
	}
	*reply = value
	return nil
}

func (c *CacheService) Put(args [2]string, reply *bool) error {
	c.cache.Put(args[0], args[1])
	*reply = true
	return nil
}
