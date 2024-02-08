package leetcode

func MaxArea(height []int) int {
	firstPtr := 0
	lastPtr := len(height) - 1
	maxArea := 0
	for firstPtr < lastPtr {
		firstHeight := height[firstPtr]
		lastHeight := height[lastPtr]
		var area int
		if lastHeight > firstHeight {
			area = firstHeight * (lastPtr - firstPtr)
		} else {
			area = lastHeight * (lastPtr - firstPtr)
		}

		if area > maxArea {
			maxArea = area
		}

		if firstHeight < lastHeight {
			firstPtr++
		} else {
			lastPtr--
		}
	}
	return maxArea
}
