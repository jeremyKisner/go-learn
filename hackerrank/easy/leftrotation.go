package hackerrank

func RotateLeft(d int32, arr []int32) []int32 {
	var counter int32
	for counter = 0; counter < d; counter++ {
		first := arr[0]
		arr = append(arr[1:], first)
	}
	return arr
}
