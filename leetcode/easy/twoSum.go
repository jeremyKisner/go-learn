package leetcode

// TwoSum returns the ints that equal the target.
func TwoSum(nums []int, target int) []int {
	result := []int{}
	for i := 0; i < len(nums); i++ {
		result = calculateTarget(i, nums, target, result)
		if len(result) != 0 && i != 0 {
			break
		}
	}
	return result
}

func calculateTarget(i int, nums []int, target int, result []int) []int {
	for j := i + 1; j < len(nums); j++ {
		if nums[i]+nums[j] == target {
			result = append(result, i, j)
			break
		}
	}
	return result
}
