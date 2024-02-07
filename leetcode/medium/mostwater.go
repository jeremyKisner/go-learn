package leetcode

import "fmt"

func MaxArea(height []int) int {
	var first, last int
	for i := range height {
		first = height[i]
		for j := len(height) - i - 1; j > 0; j-- {
			last = height[j]
			fmt.Println(height[i], height[j])
			if first == last {
				if first == 1 {
					return first
				} else {
					return (first - 1) * (last - 1)
				}
			}
		}
	}
	return 0
}
