package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	fmt.Println("Server starting on :8090...")

	log.Fatal(http.ListenAndServe(":8090", nil))
}
