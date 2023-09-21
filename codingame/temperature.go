package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
	"strings"
)

/**
 * Auto-generated code below aims at helping you parse
 * the standard input according to the problem statement.
 **/

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Buffer(make([]byte, 1000000), 1000000)
	var inputs []string

	// n: the number of temperatures to analyze
	var n int
	scanner.Scan()
	fmt.Sscan(scanner.Text(), &n)

	scanner.Scan()
	inputs = strings.Split(scanner.Text(), " ")

	result := GetClosestTemperatureToZero(n, inputs)

	fmt.Println("result: ", result)
}

func GetClosestTemperatureToZero(n int, inputs []string) string {
	var belowZero int64
	var aboveZero int64

	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		if t > 0 {
			if aboveZero == 0 {
				aboveZero = t
			} else if t < aboveZero {
				aboveZero = t
			}
		} else if t < 0 {
			if belowZero == 0 {
				belowZero = t
			} else if t > belowZero {
				belowZero = t
			}
		}
	}

	if aboveZero == 0 {
		return strconv.FormatInt(belowZero, 10)
	}

	if belowZero == 0 {
		return strconv.FormatInt(aboveZero, 10)
	}

	if math.Abs(float64(belowZero)) < float64(aboveZero) && aboveZero != 0 {
		return strconv.FormatInt(belowZero, 10)
	}

	return strconv.FormatInt(aboveZero, 10)
}
