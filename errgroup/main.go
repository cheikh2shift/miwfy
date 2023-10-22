package main

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"golang.org/x/sync/errgroup"
)

type errGroupFn func() error

func main() {

	ctx := context.Background()

	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(2)
	for i := 0; i < 7; i++ {
		g.Go(heavyTask(i))
	}

	if err := g.Wait(); err != nil {
		log.Println(err)
	}

}

func heavyTask(index int) errGroupFn {
	return func() error {
		log.Println("Processing", index)
		time.Sleep(2 * time.Second)
		if index == 3 {
			return errors.New(
				fmt.Sprintf(
					"Random error from index %v",
					index,
				))
		}
		return nil
	}
}
