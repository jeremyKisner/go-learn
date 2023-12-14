package leetcode

func DistributeCandies(n int, limit int) int64 {
	var ans int64
	for i := 0; i <= min(limit, n); i++ {
		if float64(n-i)/2. > float64(limit) {
			continue
		}
		e := min(limit, n-i)
		s := n - i - e
		ans += int64(e - s + 1)
	}
	return ans
}
