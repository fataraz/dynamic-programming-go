package helper

// write a function `canSum(targetSum, numbers)` that takes in a
// targetSum and an array of numbers as arguments.
// the function should return a boolean indicating whether or not it
// is possible to generate the targetSum using numbers from the array.
// canSum(7, [5,3,4,7]) -> true
// canSum(7, [2,4]) -> false
// canSum(7, [2,3,5) -> true

// m = target sum
// n = array length
// O(n**m) time
// O(m) space
func CanSum(targetSum int, numbers []int) bool {
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, number := range numbers {
		remainder := targetSum - number
		if CanSum(remainder, numbers) == true {
			return true
		}
	}

	return false
}

// m = target sum
// n = array length
// O(m*n) time
// O(m) space
func CanSumMemorization(targetSum int, numbers []int, memo map[int]bool) bool {
	if memo[targetSum] == true {
		return memo[targetSum]
	}
	if targetSum == 0 {
		return true
	}
	if targetSum < 0 {
		return false
	}

	for _, number := range numbers {
		remainder := targetSum - number
		if CanSumMemorization(remainder, numbers, memo) == true {
			memo[targetSum] = true
			return true
		}
	}

	memo[targetSum] = false
	return false
}
