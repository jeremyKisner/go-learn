package main

import (
	"fmt"

	"github.com/go-learn/learnpub/algebra"
	"github.com/go-learn/learnpub/mather"
)

type Message struct {
	*algebra.OrderedPair
}

func main() {
	op := &algebra.OrderedPair{X: 0, Y: 0}
	fmt.Println(op.Quadrant())

	m := Message{op}
	fmt.Println(m.Quadrant())
	if m.IsPlotted {
		fmt.Println("should not work")
	}
	m.IsPlotted = true
	if m.IsPlotted {
		fmt.Println("should work")
	}

	athlete := mather.Athlete{
		Speed: 5,
		Steps: 10,
	}

	var track mather.Track = athlete

	fmt.Println("Running:", track.Run())
	fmt.Println("Walking:", track.Walk())
}
