package hackerrank

func findMax(input map[int32]int32) int32 {
	var key, max int32
	for k, v := range input {
		if key == 0 && max == 0 {
			key = k
			max = v
			continue
		}
		if v > max && k < key {
			key = k
			max = v
		}
	}
	return key
}

func MigratoryBirds(arr []int32) int32 {
	bt := make(map[int32]int32)
	for _, v := range arr {
		value, ok := bt[v]
		if !ok {
			bt[v] = 1
		} else {
			bt[v] = value + 1
		}
	}
	return findMax(bt)
}
