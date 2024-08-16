package main

import (
	"context"
	"fmt"
	"time"
)

func doWork(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			fmt.Println("Operation canceled")
		case <-time.After(2 * time.Second):
			fmt.Println("Operation completed")
		default:
			fmt.Println("Processing...")
		}
	}
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 3*time.Second)
	defer cancel()

	go doWork(ctx)

	time.Sleep(5 * time.Second) // Longer sleep duration
	cancel()
}
