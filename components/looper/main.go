package main

import (
	"context"
	"flag"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	deadlineSecs := flag.Int("deadline", 5, "deadline in seconds")
	numberOfDie := flag.Int("dieCount", 1, "number of dice (default 1 die)")
	flag.Parse()

	fmt.Println("deadline in seconds", *deadlineSecs)
	deadline := time.Now().Add(time.Duration(*deadlineSecs) * time.Second)
	ctx, cancel := context.WithDeadline(context.Background(), deadline)
	defer cancel()

	var wg sync.WaitGroup
	wg.Add(1)
	boardConfig := RollScoreBoard{
		Game:     &wg,
		totalDie: *numberOfDie,
	}
	go Roll(ctx, &boardConfig)
	wg.Wait()

	fmt.Println("complete!")
}

type RollScoreBoard struct {
	totalDie   int
	totalRolls int
	wins       int
	losses     int
	draws      int
	lock       sync.Mutex
	Game       *sync.WaitGroup
}

func Roll(ctx context.Context, sb *RollScoreBoard) {
	defer sb.Game.Done()
	fmt.Println("start rolling")
	numberOfDie := sb.totalDie
	for {
		select {
		case <-ctx.Done():
			fmt.Printf(fmt.Sprintf("total rolls: %d | wins: %d - loses: %d - draws: %d\n", sb.totalRolls, sb.wins, sb.losses, sb.draws))
			return
		default:
			// roll player dice
			pTotal := 0
			for i := 0; i < numberOfDie; i++ {
				pTotal += GetDie()
			}
			// roll computer dice
			compTotal := 0
			for i := 0; i < numberOfDie; i++ {
				compTotal += GetDie()
			}
			msg := fmt.Sprintf("you rolled %d  -- comp rolled %d", pTotal, compTotal)
			var res string
			if pTotal > compTotal {
				res = "you win"
				Increment(&sb.lock, &sb.wins)
			} else if pTotal == compTotal {
				res = "draw"
				Increment(&sb.lock, &sb.draws)
			} else {
				res = "you lose"
				Increment(&sb.lock, &sb.losses)
			}
			Increment(&sb.lock, &sb.totalRolls)
			fmt.Println("round", sb.totalRolls, ": ", res, msg)
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
