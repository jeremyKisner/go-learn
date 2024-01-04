package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	numOfAttempts int
	lock          sync.Mutex
)

func incrementAttempts() {
	lock.Lock()
	numOfAttempts++
	lock.Unlock()
}

func attack(item string) {
	incrementAttempts()
	fmt.Println("attacking ", item)
}

func main() {
	ninja := "Jeremy"
	go attack(ninja)
	time.Sleep(1 * time.Second)
	fmt.Println("total attempts: ", numOfAttempts)
}
