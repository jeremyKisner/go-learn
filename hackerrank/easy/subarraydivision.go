package hackerrank

import "fmt"

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

func Birthday(input []int32, maxTotal int32, numOfSquares int32) int32 {
	var res int32

	for i := 0; i < len(input); i++ {
		temp := input[i]
		fmt.Println(temp)
		var count int
		var curSum int
		var curNums []int32
		for j := int(temp); j < len(input); j++ {
			if count < int(numOfSquares) {
				newSum := curSum + int(input[j])
				if newSum == int(numOfSquares) {
					curNums = append(curNums, input[j])
				}
			} else {
				break
			}
			count++
		}
	}
	return res
}
