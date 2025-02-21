package main

import (
	"distributed-cache/internal/node"
	"log"
)

func main() {
	node := node.NewNode(":8080", 100, "LRU")
	log.Fatal(node.Start())
}
