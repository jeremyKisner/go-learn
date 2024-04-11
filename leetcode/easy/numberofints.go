package leetcode

import (
	"strconv"
	"strings"
)

func NumberOfDifferentIntegers(word string) int {
	uniqueints := make(map[int]bool)
	var placeholder string
	for _, k := range strings.Split(word, "") {
		_, err := strconv.Atoi(k)
		if err != nil {
			if placeholder != "" {
				key, _ := strconv.Atoi(placeholder)
				uniqueints[key] = true
				placeholder = ""
			}
			continue
		} else {
			placeholder += k
		}
	}
	if placeholder != "" {
		key, _ := strconv.Atoi(placeholder)
		uniqueints[key] = true
		placeholder = ""
	}
	return len(uniqueints)
}
