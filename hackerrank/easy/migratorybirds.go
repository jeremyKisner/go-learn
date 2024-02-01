package hackerrank

func MigratoryBirds(arr []int32) int32 {
	var keyId, max int32
	bt := make(map[int32]int32)
	for _, key := range arr {
		value, ok := bt[key]
		if !ok {
			bt[key] = 1
		} else {
			bt[key] = value + 1
		}

		value = bt[key]
		if value > max {
			keyId = key
			max = value
		} else if value == max && key < keyId {
			keyId = key
			max = value
		}
	}
	return keyId
}
