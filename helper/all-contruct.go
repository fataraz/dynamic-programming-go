package helper

// write a function that accepts a target string and an array of strings.
// the function should return a 2D array containing all of the ways that the `target` can be constructed by concatenating
// elements of the `wordBank` array. Each element of the 2D array should represent one combination that constructs the `target`

// allConstruct(hello, [cat, dog, mouse]) -> []
// allConstruct('', [cat,dog,mouse]) -> [[]]

// m = target length
// n = wordBank length
// time: O(n^m)
// space: O(m)
func allConstruct(target string, wordBank []string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	var result [][]string
	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			suffix := target[len(word):]
			suffixWays := allConstruct(suffix, wordBank)

			for _, way := range suffixWays {
				result = append(result, append([]string{word}, way...))
			}
		}
	}

	return result
}

func AllConstructMemorized(target string, wordBank []string) [][]string {
	memo := make(map[string][][]string)
	return construct(target, wordBank, memo)
}

func construct(target string, wordBank []string, memo map[string][][]string) [][]string {
	if target == "" {
		return [][]string{{}}
	}

	if val, found := memo[target]; found {
		return val
	}

	var result [][]string

	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			suffix := target[len(word):]
			suffixWays := construct(suffix, wordBank, memo)

			for _, way := range suffixWays {
				result = append(result, append([]string{word}, way...))
			}
		}
	}
	memo[target] = result

	return result
}
