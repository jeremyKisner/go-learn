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
	go Roll(ctx, &wg)
	wg.Wait()
	fmt.Print("complete!")
}

type RollScoreBoard struct {
	totalRolls int
	wins       int
	losses     int
	draws      int
	lock       sync.Mutex
}

func Roll(ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("start rolling")
	sb := RollScoreBoard{}
	for {
		select {
		case <-ctx.Done():
			fmt.Println(fmt.Sprintf("total tolls: %s | wins: %s - loses: %s - draws: %s",
				strconv.Itoa(sb.totalRolls), strconv.Itoa(sb.wins), strconv.Itoa(sb.losses), strconv.Itoa(sb.draws)))
			return
		default:
			pTotal := GetDie() + GetDie()
			compTotal := GetDie() + GetDie()
			msg := fmt.Sprintf("you rolled %d  -- comp rolled %d", pTotal, compTotal)
			var res string
			if pTotal > compTotal {
				res = fmt.Sprintf("you win! ")
				Increment(&sb.lock, &sb.wins)
			} else if pTotal == compTotal {
				res = fmt.Sprintf("draw! ")
				Increment(&sb.lock, &sb.draws)
			} else {
				res = fmt.Sprintf("you lose! ")
				Increment(&sb.lock, &sb.losses)
			}
			Increment(&sb.lock, &sb.totalRolls)
			fmt.Println("round", sb.totalRolls, res, msg)
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
