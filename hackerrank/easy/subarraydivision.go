package hackerrank

// s = number of chocolates
// d = birthday
// m = birth month

func Birthday(s []int32, d int32, m int32) int32 {
	var res int32

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
