package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"strconv"
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
	go Start(ctx, &wg)
	wg.Wait()
}

func Start(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start rolling")
	Roll(ctx)
	fmt.Println("rolling complete")
}

func Roll(ctx context.Context) {
	var totalRolls int
	var wins int
	var losses int
	var draws int
	var lock sync.Mutex
	for {
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("Total Rolls: %s | Wins: %s - Loses: %s - Draws: %s",
				strconv.Itoa(totalRolls), strconv.Itoa(wins), strconv.Itoa(losses), strconv.Itoa(draws)))
			return
		default:
			pTotal := GetDie() + GetDie()
			compTotal := GetDie() + GetDie()
			msg := fmt.Sprintf("you rolled %d  -- comp rolled %d", pTotal, compTotal)
			var res string
			if pTotal > compTotal {
				res = fmt.Sprintf("you win! ")
				Increment(&lock, &wins)
			} else if pTotal == compTotal {
				res = fmt.Sprintf("draw! ")
				Increment(&lock, &draws)
			} else {
				res = fmt.Sprintf("you lose! ")
				Increment(&lock, &losses)
			}
			Increment(&lock, &totalRolls)
			fmt.Println("round", totalRolls, res, msg)
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

func Increment(lock *sync.Mutex, num *int) {
	lock.Lock()
	*num++
	lock.Unlock()
}
