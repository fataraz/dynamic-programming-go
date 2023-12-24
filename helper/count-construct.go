package helper

// write a function, that accepts a target string and an array of strings.
// the function should return the number of ways that the `target` can
// be constructed by concatenating element of the `wordBank` array.
// you may reuse element of `wordBank` as many times as needed.

// countConstruct(abcdef, [ab, abc, cd, def, abcd]) -> 1
// countConstruct(purple, [purp, p, ur, le, purpl]) -> 2
// countConstruct(enterapotentpot, [a, p, ent, enter,ot,o,t]) -> 4

// brute force
// time: O(n^m * m)
// space: O(m^2)
func CountConstruct(target string, wordBank []string) int {
	if target == "" {
		return 1
	}

	totalCount := 0
	for _, word := range wordBank {
		if len(word) <= len(target) && target[:len(word)] == word {
			suffix := target[len(word):]
			ways := CountConstruct(suffix, wordBank)
			totalCount += ways
		}
	}

	return totalCount
}

//lorem := "enterapotentpot"
//fmt.Println(lorem[1:]) // "nterapotentpot"
//fmt.Println(lorem[5:]) // "apotentpot"
//fmt.Println(lorem[:1]) // "e"

// memoized
// time: O(n * m^2)
// space: O(m^2)
func CountConstructMemorization(target string, wordBank []string) int {
	memo := make(map[string]int)

	return CountConstructHelper(target, wordBank, memo)
}

func CountConstructHelper(target string, wordBank []string, memo map[string]int) int {
	if val, ok := memo[target]; ok {
		return val
	}
	if target == "" {
		return 1
	}

	totalCount := 0
	for _, word := range wordBank {
		if len(word) <= len(target) && target[:len(word)] == word {
			suffix := target[len(word):]
			ways := CountConstructHelper(suffix, wordBank, memo)
			totalCount += ways
		}
	}
	memo[target] = totalCount

	return totalCount
}
