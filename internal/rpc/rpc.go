package rpc

import (
	"net"
	"net/rpc"
)

type Cache interface {
	Get(key string) (interface{}, bool)
	Put(key string, value string)
}

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

//initialising and running the RPC server

func startServer(cache Cache, address string) error {
	service := &CacheService{cache: cache}
	rpc.Register(service)

	listener, err := net.Listen("tcp", address)
	if err != nil {
		return err
	}
	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}

		go rpc.ServeConn(conn)
	}

}
