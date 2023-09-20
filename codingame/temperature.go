package main

import (
	"bufio"
	"fmt"
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

	result := GetLowestTemperature(n, inputs)

	fmt.Println("result", result)
}

func GetLowestTemperature(n int, inputs []string) string {
	var lowestTemp int64

	for i := 0; i < n; i++ {
		// t: a temperature expressed as an integer ranging from -273 to 5526
		t, _ := strconv.ParseInt(inputs[i], 10, 32)
		if lowestTemp == 0 {
			lowestTemp = t
		} else if t < lowestTemp {
			// TODO
		}
	}
	return strconv.FormatInt(lowestTemp, 10)
}
