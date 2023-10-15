package hackerrank

/*
 * Complete the 'birthday' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY s
 *  2. INTEGER d
 *  3. INTEGER m
 */

// s = number of chocolates
// d = birthday
// m = birth month

/*
 iterate the number of chocolates
 add up the current number
 if sum of current numbers equal d
 then count all the summed numbers equal m

 if sum of numbers is too low continue

 if sum is too much reset
*/

func Birthday(s []int32, d int32, m int32) int32 {
	var res int32

	// guard clause for a single item
	if len(s) == 1 {
		if s[0] == d && m == 1 {
			res++
			return res
		}
	}

	for i := 0; i < len(s); i++ {
		var sum int32
		var nums int32
		for j := i; j < len(s); j++ {
			sum += s[j]
			nums++
			if sum == d && nums == m {
				res++
			} else if sum < d {
				continue
			} else {
				break
			}
		}
	}
	return res
}
