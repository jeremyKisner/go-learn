package leetcode

import "math"

// Special Instructions:
// Assume the environment does not allow you to store 64-bit integers (signed or unsigned).

func ReverseInt(x int) int {
	// Handle the case when x is 0 separately
	if x == 0 {
		return 0
	}

	// Keep track of the sign
	sign := 1
	if x < 0 {
		sign = -1
		x = -x
	}

	// Reverse the digits without using strings
	reversed := 0
	for x > 0 {
		digit := x % 10
		// Check for overflow before updating reversed
		if reversed > (math.MaxInt32-digit)/10 {
			return 0 // Overflow will occur, return 0
		}
		reversed = reversed*10 + digit
		x /= 10
	}

	return sign * reversed
}
