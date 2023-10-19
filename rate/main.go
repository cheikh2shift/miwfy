package main

import (
	"context"
	"log"
	"time"

	"golang.org/x/time/rate"
)

func main() {

	limiter := rate.NewLimiter(10, 5)
	ctx, _ := context.WithDeadline(
		context.Background(),
		time.Now().Add(10*time.Second),
	)
	for i := 0; i < 20; i++ {

		go func(i int) {

			log.Printf(
				"ETA %v to process %v\n",
				limiter.Reserve().Delay(),
				i,
			)

			err := limiter.Wait(ctx)

			if err != nil {
				log.Println("Error processing", i, err)
				return
			}

			log.Println("n", i)
			time.Sleep(2 * time.Second)

		}(i)

	}

	time.Sleep(10 * time.Second)
}
