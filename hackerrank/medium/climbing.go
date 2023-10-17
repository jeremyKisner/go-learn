package medium

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
	ranked = BuildLeaderBoard(player, ranked)
	fmt.Println(res)
	return res
}

func BuildLeaderBoard(player []int32, ranked []int32) []int32 {
	for i := len(player) - 1; i >= 0; i-- {
		start := 0
		end := len(ranked) - 1
		firstPlace := ranked[start]
		lastPlace := ranked[end]
		playerScore := player[i]
		if playerScore >= firstPlace {
			fmt.Println("new high score!")
			ranked = append([]int32{playerScore}, ranked...)
		} else if playerScore <= lastPlace {
			fmt.Println("tough luck, do better next time!")
			ranked = append(ranked, lastPlace)
		} else {
			mid := end / 2
			midPlace := ranked[mid]
			if playerScore == midPlace {
				fmt.Println("playerscore equals midplace")
				ranked = append(ranked[:mid], append([]int32{(playerScore)}, ranked[mid:]...)...)
			} else if playerScore > midPlace {
				fmt.Println("playerscore >= than midplace")
				for j := mid; j >= 0; j-- {
					temp := ranked[j]
					if playerScore > temp {
						ranked = append(ranked[:j], append([]int32{playerScore}, ranked[j:]...)...)
						break
					}
				}
			} else {
				fmt.Println("playerscore less than midplace")
				for j := mid; j < len(ranked); j++ {
					temp := ranked[j]
					if playerScore > temp {
						ranked = append(ranked[:j], append([]int32{playerScore}, ranked[j:]...)...)
						break
					}
				}
			}
		}
	}
	fmt.Println(ranked)
	return ranked
}
