package leetcode

func SortArrayByParity(nums []int) []int {
	var res []int
	for i := 0; i < len(nums); i++ {
		if nums[i]%2 == 0 {
			temp := []int{nums[i]}
			res = append(temp, res...)
		} else {
			res = append(res, nums[i])
		}
	}
	return res
}
