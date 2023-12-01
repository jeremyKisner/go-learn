package leetcode

func MinNumberOfHours(initialEnergy int, initialExperience int, energy []int, experience []int) int {
	var totalEnergy int
	var expToGain int
	curExp := initialExperience
	for i := 0; i < len(energy); i++ {
		totalEnergy += energy[i]
		winningExp := experience[i] - curExp + 1
		if winningExp > expToGain {
			expToGain = winningExp
		}
		curExp += experience[i]
	}
	requiredEnergy := totalEnergy - initialEnergy + 1
	if requiredEnergy < 0 {
		requiredEnergy = 0
	}

	return expToGain + requiredEnergy
}
