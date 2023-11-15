package main

import (
	"log"
	"sync"
	"time"
)

func loadDb() string {
	time.Sleep(5 * time.Second)
	return "DB-DATA"
}

func main() {

	cLoadDb := sync.OnceValue(loadDb)

	normalTime := time.Now()

	for i := 0; i < 2; i++ {
		go func(i int) {
			loadDb()
			log.Println("Complete", i)

			if i == 1 {
				log.Println(time.Since(normalTime))
			}
		}(i)
		time.Sleep(2 * time.Second)
	}

	syncTime := time.Now()
	for i := 0; i < 2; i++ {

		go func(i int) {
			cLoadDb()
			log.Println("Complete Sync", i)

			if i == 1 {
				log.Println(time.Since(syncTime))
			}
		}(i)
		time.Sleep(2 * time.Second)
	}

	// Wait for code to complete
	c := make(chan int)
	<-c
}
