package helper

// write a function that accepts a target string and an array of strings
// The function should return a boolean indicating whether or not the `target` can be constructed
// by concatenating elements of the `wordBank` array.

// canConstruct("abcdef", ["ab","abc","cd","def","abcd"]) // true
// canConstruct("skateboard", ["bo","rd","ate","t","ska","sk","boar"]) // false
// canConstruct("enterapotentpot", ["a","p","ent","enter","ot","o","t"]) // true

// m = target length
// n = wordBank length
// time = O(n^m*m)
// space = O(m^2)
func CanConstruct(target string, wordBank []string, memo ...*map[string]bool) bool {
	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			if CanConstruct(target[len(word):], wordBank) {
				return true
			}
		}
	}

	return false
}

//lorem := "enterapotentpot"
//fmt.Println(lorem[1:]) // "nterapotentpot"
//fmt.Println(lorem[5:]) // "apotentpot"
//fmt.Println(lorem[:1]) // "e"

// Memoized
// time: O(n*m^2)
// space: O(m^2)
func CanConstructMemorization(target string, wordBank []string, memo ...*map[string]bool) bool {
	canConstruct := make(map[string]bool)
	if len(memo) != 0 {
		canConstruct = *memo[0]
	}
	// check memoization table
	if val, ok := canConstruct[target]; ok {
		return val
	}

	if target == "" {
		return true
	}
	for _, word := range wordBank {
		if len(target) >= len(word) && target[:len(word)] == word {
			if CanConstructMemorization(target[len(word):], wordBank, &canConstruct) {
				canConstruct[target] = true
				return canConstruct[target]
			}
		}
	}

	canConstruct[target] = false
	return canConstruct[target]
}
