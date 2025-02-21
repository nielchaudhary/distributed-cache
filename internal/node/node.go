package node

import (
	"distributed-cache/internal/cache"
	"distributed-cache/internal/consistent_hashing"
	"distributed-cache/internal/rpc"
)

type Node struct {
	address string
	cache   cache.Cache
	ring    *consistent_hashing.ConsistentHashing
}

func NewNode(address string, capacity int, policy string) *Node {
	var c cache.Cache
	if policy == "LRU" {
		c = cache.NewLRUCache(capacity)
	}
	return &Node{
		address: address,
		cache:   c,
		ring:    consistent_hashing.NewConsistentHashing(),
	}
}

func (n *Node) Start() error {
	return rpc.StartServer(n.cache, n.address)
}
