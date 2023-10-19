package leetcode

// TwoSum will interate an array of ints to figure out what two equal target
func TwoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		for j := i + 1; j < len(nums); j++ {
			if nums[i]+nums[j] == target {
				result = append(result, i, j)
				break
			}
		}
		if len(result) != 0 && i != 0 {
			break
		}
	}
	return result
}
