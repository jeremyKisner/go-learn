package leetcode

func NumberOfDifferentIntegers(word string) int {
	// map for storing the number
	m := make(map[string]struct{})
	// stack for storing the number digit by digit
	stack := ""

	// iterates the word
	for i, w := range word {
		// if digit
		if isDigit(w) {
			// if stack is 0, it's mean the number is leading zero
			// e.g. 001 then remove the 0
			if stack == "0" {
				stack = ""
			}

			// push digit on the stack
			stack += string(w)

			// if current index isn't last index, continue iterates
			if i != len(word)-1 {
				continue
			}
		}

		// starting this will execute if non digit number found
		// if stack is empty, then no number need to store
		if stack == "" {
			continue
		}

		// store to the map, reset the stack
		m[stack] = struct{}{}
		stack = ""
	}

	// count number of digit
	return len(m)
}

// isDigit checks the rune is digit or not
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}
