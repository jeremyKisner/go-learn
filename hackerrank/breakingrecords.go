package hackerrank

func BreakingRecords(scores []int32) []int32 {
	var res []int32
	var numOfBrokenRecords int32
	var numOfFailedAttempts int32
	var minimum int32
	var maximum int32

	for i := 0; i < len(scores); i++ {
		newScore := scores[i]
		if minimum == 0 && maximum == 0 {
			minimum = newScore
			maximum = newScore
			continue
		}

		if newScore > maximum {
			maximum = newScore
			numOfBrokenRecords++
		} else if newScore < minimum {
			minimum = newScore
			numOfFailedAttempts++
		}
	}

	res = append(res, numOfBrokenRecords, numOfFailedAttempts)
	return res
}
