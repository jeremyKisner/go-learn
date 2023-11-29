package leetcode

import "fmt"

func MergeAlternatively(word1 string, word2 string) string {
	var counter int
	w1Len := len(word1)
	w2Len := len(word2)

	if w2Len == 0 {
		return word1
	} else if w1Len == 0 {
		return word2
	} else if w1Len < w2Len {
		counter = w1Len
	} else {
		counter = w2Len
	}

	var resArr []byte
	for i := 0; i < counter; i++ {
		resArr = append(resArr, word1[i], word2[i])
	}

	res := string(resArr)
	if counter < w1Len {
		res += word1[counter:]
	} else {
		fmt.Println(word2[counter:])
		res += word2[counter:]
	}

	return res
}
