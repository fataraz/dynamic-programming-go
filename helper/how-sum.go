package helper

// write a function that takes in a targetSum and an array of numbers as arguments.
// the function should return an array containing any combination of elements that add up to the targetSum.
// If there is no combinations that adds up tp the targetSum, then return null
// If there are multiple combinations possible, you may return any single one.

// howSum(7, [5,3,4,7]) -> [3,4]

// m = target sum
// n = numbers.length
// time = O(n^m*m)
// space = O(m)
func HowSum(targetSum int, numbers []int) []int {
	howSum := make(map[int][]int)
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, number := range numbers {
		remainder := targetSum - number
		remainderResult := HowSum(remainder, numbers)
		if remainderResult != nil {
			remainderResult = append(remainderResult, number)
			howSum[targetSum] = remainderResult
			return howSum[targetSum]
		}
	}

	return howSum[targetSum]
}

// time: O(n*m^2)
// space: O(m^2)
func HowSumMemorization(targetSum int, numbers []int, memo ...*map[int][]int) []int {
	howSum := make(map[int][]int)
	if len(memo) != 0 {
		howSum = *memo[0]
	}
	if val, ok := howSum[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	for _, number := range numbers {
		remainder := targetSum - number
		remainderResult := HowSumMemorization(remainder, numbers, &howSum)
		if remainderResult != nil {
			remainderResult = append(remainderResult, number)
			howSum[targetSum] = remainderResult
			return howSum[targetSum]
		}
	}

	howSum[targetSum] = nil
	return howSum[targetSum]
}
