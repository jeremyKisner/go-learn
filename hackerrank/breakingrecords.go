package hackerrank

func BreakingRecords(scores []int32) []int32 {
	var res []int32
	var numOfBrokenRecords int32
	var numOfFailedAttempts int32
	var minimum int32
	var maximum int32

	if len(scores) > 1 {
		minimum = scores[0]
		maximum = scores[0]
	}
	for i := 1; i < len(scores); i++ {
		newScore := scores[i]

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
