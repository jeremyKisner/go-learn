package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var deadlineSecs = flag.Int("deadline", 5, "deadline in seconds")

func main() {
	flag.Parse()
	fmt.Println("deadline in seconds", *deadlineSecs)
	deadline := time.Now().Add(time.Duration(*deadlineSecs) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()
	var wg sync.WaitGroup
	wg.Add(1)
	go func() {
		defer wg.Done()
		fmt.Println("start polling")
		Roll(ctx)
	}()
	wg.Wait()
}

func Roll(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("polling process complete within the deadline")
			return
		default:
			die, die2 := GetDie(), GetDie()
			fmt.Println("rolled", die+die2)
			time.Sleep(2 * time.Second)
		}
	}
}

func GetDie() int {
	min := 1
	max := 6
	randNum := rand.Intn(max-min+1) + min
	return randNum
}
