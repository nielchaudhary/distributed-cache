package main

import (
	"log"
	"net/rpc"
)

func main() {
	client, err := rpc.Dial("tcp", "localhost:8080")
	if err != nil {
		log.Fatal(err)
	}

	var reply interface{}
	err = client.Call("CacheService.Get", "key", &reply)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(reply)
}
