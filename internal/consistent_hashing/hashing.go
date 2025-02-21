package consistent_hashing

import (
	"hash/crc32"
	"sort"
)

type HashRing []uint32

func (h HashRing) Len() int           { return len(h) }
func (h HashRing) Less(i, j int) bool { return h[i] < h[j] }
func (h HashRing) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }

type ConsistentHashing struct {
	nodes      map[uint32]string
	sortedKeys HashRing
}

func NewConsistentHashing() *ConsistentHashing {
	return &ConsistentHashing{
		nodes:      make(map[uint32]string),
		sortedKeys: HashRing{},
	}
}

func (c *ConsistentHashing) AddNode(node string) {
	hash := c.hash(node)
	c.nodes[hash] = node
	c.sortedKeys = append(c.sortedKeys, hash)
	sort.Sort(c.sortedKeys)
}

func (c *ConsistentHashing) GetNode(key string) string {
	hash := c.hash(key)
	idx := sort.Search(len(c.sortedKeys), func(i int) bool { return c.sortedKeys[i] >= hash })
	if idx == len(c.sortedKeys) {
		idx = 0
	}
	return c.nodes[c.sortedKeys[idx]]
}

func (c *ConsistentHashing) hash(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
