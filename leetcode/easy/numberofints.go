package leetcode

import (
	"fmt"
	"strconv"
	"strings"
)

func NumberOfDifferentIntegers(word string) int {
	uniqueints := make(map[string]bool)
	var placeholder string
	for i, k := range strings.Split(word, "") {
		fmt.Println(i, k)
		parsed, err := strconv.Atoi(k)
		if err != nil {
			if placeholder != "" {
				uniqueints[placeholder] = true
				placeholder = ""
			}
			continue
		} else {
			fmt.Println("parsed", parsed)
			placeholder += k
		}
	}
	return len(uniqueints)
}
