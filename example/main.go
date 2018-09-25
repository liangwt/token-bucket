package main

import (
	"github.com/liangwt/token-bucket"
	"log"
	"time"
)

func main() {
	bucket := bucket.New(10, time.Second, 100, true)
	isLimit := bucket.Consume(100)
	log.Println(isLimit)

	time.Sleep(2 * time.Second)
	isLimit = bucket.Consume(10)
	log.Println(isLimit)
}
