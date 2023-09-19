package leetcode

import (
	"strings"
)

type TokenManager struct {
	startPos int
	endPose  int
}

func evaluate(s string, knowledge [][]string) string {
	result := ""
	separators := "()"

	// Split the string using the custom separator
	substrings := strings.FieldsFunc(s, func(r rune) bool {
		return strings.ContainsRune(separators, r)
	})
	for i := 0; i < len(substrings); i++ {
		eval := substrings[i]
		if eval == "" {
			continue
		}
		isFound := false
		for j := 0; j < len(knowledge); j++ {
			if knowledge[j][0] == eval {
				result += knowledge[j][1]
				isFound = true
				break
			}
		}
		if !isFound {
			result += eval
		}
	}
	return result
}
