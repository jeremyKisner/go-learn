package main

import (
	"fmt"
	"sync"
)

var (
	numOfAttempts int
	lock          sync.Mutex
)

type Ninja struct {
	Ordinal int
	Name    string
}

func incrementAttempts() int {
	var curAttempt int
	lock.Lock()
	numOfAttempts++
	curAttempt = numOfAttempts
	lock.Unlock()
	return curAttempt
}

func attack(ninja Ninja, wg *sync.WaitGroup) {
	attempt := incrementAttempts()
	fmt.Println("attacking", ninja, attempt)
	wg.Done()
}

func main() {
	ninjas := []string{"Jeremy", "George", "Joe"}
	var beeper sync.WaitGroup

	for k, v := range ninjas {
		beeper.Add(1)
		n := Ninja{
			Ordinal: k,
			Name:    v,
		}
		go attack(n, &beeper)
	}

	beeper.Wait()
	fmt.Println("total attempts: ", numOfAttempts)
}
