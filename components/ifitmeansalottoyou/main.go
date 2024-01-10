package main

import (
	"context"
	"flag"
	"fmt"
	"strings"
	"sync"
	"time"
)

func main() {
	flag.Parse()
	deadline := time.Now().Add(time.Duration(40) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var wg sync.WaitGroup
	fmt.Println("... we knew it'd happen eventually ...")
	wg.Add(1)
	go Start(ctx, &wg, 1)
	time.Sleep(7 * time.Second)

	for i := 0; i < 3; i++ {
		fmt.Println("... now everybody's singing ...")
		wg.Add(3)
		go Start(ctx, &wg, 1)
		go Start(ctx, &wg, 2)
		go Start(ctx, &wg, 3)
		time.Sleep(12 * time.Second)
	}

	wg.Wait()
}

func Start(ctx context.Context, wg *sync.WaitGroup, num int) {
	defer wg.Done()
	Sing(ctx, num)
	return
}

func Sing(ctx context.Context, num int) {
	for {
		select {
		case <-ctx.Done():
			return
		default:
			fmt.Println(num, strings.Repeat("la ", 7))
			time.Sleep(1 * time.Second)
		}
	}
}
