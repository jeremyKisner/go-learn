// You are given two integer arrays nums1 and nums2, sorted in non-decreasing order, and two integers m and n,
// representing the number of elements in nums1 and nums2 respectively.

// Merge nums1 and nums2 into a single array sorted in non-decreasing order.
// The final sorted array should not be returned by the function, but instead be stored inside the array nums1.
// To accommodate this, nums1 has a length of m + n, where the first m elements denote the elements that should be merged,
// and the last n elements are set to 0 and should be ignored. nums2 has a length of n.

package leetcode

import "fmt"

func Merge(nums1 []int, m int, nums2 []int, n int) {
	var first_half []int
	first_half = nums1[:m]
	fmt.Println("first_half ", first_half)
	for i := 0; i < n; i++ {
		if nums2[i] < first_half[i] {
			fmt.Println(nums2[i])
		}
	}
}
