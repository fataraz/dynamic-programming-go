package helper

// write a function that takes in a targetSum and an array of numbers as arguments.
// the function should return an array containing the shortest combination of numbers that add up to exactly the targetSum
// if there is a tie for the shortest combination, you may return any one of the shortest

// bestSum(7, [5,3,4,7]) -> [7]
// bestSum(8, [2,3,5]) -> [3,5]

// m = target sum
// n = numbers.length
// Brute Force
// time: O(n^m*m)
// space: O(m)
func BestSum(targetSum int, numbers []int) []int {
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int
	for _, number := range numbers {
		remainder := targetSum - number
		remainderCombination := BestSum(remainder, numbers)
		if remainderCombination != nil {
			remainderCombination = append(remainderCombination, number)
			// if the combination is shorter than the current "shortest", update it
			if shortestCombination == nil || len(remainderCombination) < len(shortestCombination) {
				shortestCombination = remainderCombination
			}
		}
	}

	return shortestCombination
}

// Memorized
// time: O(m * n)
// space: O(m)
func BestSumMemorization(targetSum int, numbers []int, memo ...*map[int][]int) []int {
	memoBestSum := make(map[int][]int)
	if len(memo) != 0 {
		memoBestSum = *memo[0]
	}
	if val, ok := memoBestSum[targetSum]; ok {
		return val
	}
	if targetSum == 0 {
		return []int{}
	}
	if targetSum < 0 {
		return nil
	}

	var shortestCombination []int
	for _, number := range numbers {
		remainder := targetSum - number
		remainderCombination := BestSumMemorization(remainder, numbers, &memoBestSum)
		if remainderCombination != nil {
			remainderCombination = append(remainderCombination, number)
			// if the combination is shorter than the current "shortest", update it
			if shortestCombination == nil || len(remainderCombination) < len(shortestCombination) {
				shortestCombination = remainderCombination
			}
		}
	}
	memoBestSum[targetSum] = shortestCombination

	return memoBestSum[targetSum]
}
