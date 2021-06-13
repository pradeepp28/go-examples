package main

import (
	"hash/fnv"
	"log"
)

func main() {
	key := "5bab13391c9d44000052c8df"
	hasher := fnv.New32a()
	_, err := hasher.Write([]byte(key))
	if err != nil {
		log.Fatalf("failed to write key, %v", err)
	}
	hash := int32(hasher.Sum32())
	if hash < 0 {
		hash = -hash
	}
	log.Printf("hash: %v", hash)
	log.Printf("partition: %v", hash%10)
}
