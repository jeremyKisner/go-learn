package leetcode

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var sorted []int

	sorted = append(sorted, nums1...)
	sorted = append(sorted, nums2...)

	for i := len(sorted) - 1; i >= 0; i-- {
		for j := 0; j < i; j++ {
			first := sorted[j]
			second := sorted[j+1]
			if first > second {
				sorted[j] = second
				sorted[j+1] = first
			}
		}

	}

	mid := len(sorted) / 2
	if len(sorted)%2 != 0 {
		return float64(sorted[mid])
	} else {
		return float64(sorted[mid]) + float64(.5)
	}
}
