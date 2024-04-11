package leetcode

import (
	"strconv"
	"strings"
)

func NumberOfDifferentIntegers(word string) int {
	uniqueints := make(map[string]bool)
	var placeholder string
	for _, k := range strings.Split(word, "") {
		_, err := strconv.Atoi(k)
		if err != nil {
			if placeholder != "" {
				uniqueints[placeholder] = true
				placeholder = ""
			}
			continue
		} else {
			placeholder += k
		}
	}
	if placeholder != "" {
		uniqueints[placeholder] = true
		placeholder = ""
	}
	return len(uniqueints)
}
