package hackerrank

func DivisibleSumPairs(n int32, k int32, ar []int32) int32 {
	var res int32
	for i := 0; i < len(ar); i++ {
		for j := i + 1; j < len(ar); j++ {
			if (ar[j]+ar[i])%k == 0 {
				res++
			}
		}
	}
	return res
}
