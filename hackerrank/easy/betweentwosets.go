package hackerrank

func gcd(a, b int32) int32 {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func lcm(a, b int32) int32 {
	return (a * b) / gcd(a, b)
}

func BetweenTwoSets(a []int32, b []int32) int32 {
	var lcmValue int32 = a[0]
	var gcdValue int32 = b[0]

	for _, num := range a[1:] {
		lcmValue = lcm(lcmValue, num)
	}

	for _, num := range b[1:] {
		gcdValue = gcd(gcdValue, num)
	}

	count := int32(0)
	for i := lcmValue; i <= gcdValue; i += lcmValue {
		if gcdValue%i == 0 {
			count++
		}
	}

	return count
}
