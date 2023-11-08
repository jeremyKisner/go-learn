package hackerrank

import "fmt"

/*
 * Complete the 'climbingLeaderboard' function below.
 *
 * The function is expected to return an INTEGER_ARRAY.
 * The function accepts following parameters:
 *  1. INTEGER_ARRAY ranked
 *  2. INTEGER_ARRAY player

 	we could do a binary search by splitting the scores
	for each player score,
		compare the player score to the ranks


*/

func ClimbingLeaderboard(ranked []int32, player []int32) []int32 {
	var res []int32

	for i := 0; i < len(player); i++ {
		playerScore := player[i]
		scoreFound := false
		prevScore := int32(0)
		rank := int32(0)
		for j := 0; j < len(ranked); j++ {
			r := ranked[j]
			if prevScore == 0 {
				prevScore = r
				rank += 1
			} else if r != prevScore {
				rank += 1
				prevScore = r
			}
			if playerScore >= r {
				fmt.Println("index ", j, ":", playerScore, " greater or equal ", r)
				res = append(res, rank)
				ranked = append(ranked[:j], append([]int32{(playerScore)}, ranked[j:]...)...)
				scoreFound = true
				break
			} else {
				fmt.Println("index ", j, ":", playerScore, " less than ", r)
			}

		}

		if !scoreFound {
			rank++
			ranked = append(ranked, rank)
			res = append(res, rank)
		}
	}
	return res
}
