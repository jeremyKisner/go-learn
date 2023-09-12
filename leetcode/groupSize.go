// There are n people that are split into some unknown number of groups. Each person is labeled with a unique ID from 0 to n - 1.
// You are given an integer array groupSizes, where groupSizes[i] is the size of the group that person i is in.
// For example, if groupSizes[1] = 3, then person 1 must be in a group of size 3.
// Return a list of groups such that each person i is in a group of size groupSizes[i].
// Each person should appear in exactly one group, and every person must be in a group. If there are multiple answers, return any of them. It is guaranteed that there will be at least one valid solution for the given input.

package leetcode

import (
	"fmt"
	"sort"
)

func GroupThePeople(groupSizes []int) [][]int {
	result := [][]int{}
	corral := map[int][]int{}
	for i := 0; i < len(groupSizes); i++ {
		fmt.Println(groupSizes[i])
		size := groupSizes[i]
		_, exists := corral[size]
		if !exists {
			corral[size] = []int{}
		}
		corral[size] = append(corral[size], i)
	}
	var keys []int
	for key := range corral {
		keys = append(keys, key)
	}

	// Sort the slice of keys
	sort.Ints(keys)

	for _, group := range keys {
		groupSize := len(corral[group])
		if group == groupSize {
			result = append(result, corral[group])
		} else if group > groupSize {
			result = append(result, corral[group])
		} else {
			for i := 0; i < groupSize; i += group {
				end := i + group
				if end > groupSize {
					end = groupSize
				}
				result = append(result, corral[group][i:end])
			}

		}
	}
	return result
}
