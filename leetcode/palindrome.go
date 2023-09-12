package leetcode

import "fmt"

func IsPalindrome(x int) bool {
	digits := []int{}
	if x == 0 {
		return true
	}
	if x > 0 {
		for x > 0 {
			digit := x % 10
			digits = append(digits, digit)
			x = x / 10
		}
		reverseSlice(digits)
		midpoint := len(digits) / 2

		if digits[0] == digits[len(digits)-1] {

			for i := 0; i <= midpoint; i++ {
				for j := len(digits) - i - 1; j >= midpoint; j-- {
					fmt.Printf("checking %v - %v\n", digits[i], digits[j])
					if digits[i] != digits[j] {
						fmt.Println("not a palindrome")
						return false
					}
					break
				}
			}
			return true
		}
	}

	return false
}

func reverseSlice(slice []int) {
	for i, j := 0, len(slice)-1; i < j; i, j = i+1, j-1 {
		slice[i], slice[j] = slice[j], slice[i]
	}
}
