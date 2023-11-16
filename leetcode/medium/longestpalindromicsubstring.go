package leetcode

import "strings"

func LongestPalindrome(s string) string {
	if len(s) == 0 {
		return ""
	}
	var (
		maxLen int
		maxStr string
	)
	for i := 0; i < len(s); i++ {
		for j := i + 1; j <= len(s); j++ {
			subStr := s[i:j]
			if len(subStr) > maxLen && isPalindrome(subStr) {
				maxLen = len(subStr)
				maxStr = subStr
			}
		}
	}
	return maxStr
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return false
	}
	if len(s) == 1 {
		return true
	}
	s = strings.ToLower(s)
	for i := 0; i < len(s)/2; i++ {
		if s[i] != s[len(s)-1-i] {
			return false
		}
	}
	return true
}
