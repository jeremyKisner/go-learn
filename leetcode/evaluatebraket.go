package leetcode

import (
	"fmt"
)

func evaluate(s string, knowledge [][]string) string {
	result := ""
	start := false
	startPos := 0
	for i, tok := range s {
		token := string(tok)
		if token == "(" { // Use single quotes for runes
			fmt.Println("i ", i)
			start = true
			startPos = i + 1
			continue
		}
		if token == ")" {
			start = false
			curSubStr := s[startPos:i] // Use i instead of i-1
			isFound := false
			for j := 0; j < len(knowledge); j++ {
				if knowledge[j][0] == curSubStr {
					result += knowledge[j][1]
					isFound = true
					break
				}
			}
			if !isFound {
				result += "?"
			}
			continue
		} else if !start {
			result += string(token)
		}
	}
	return result
}
