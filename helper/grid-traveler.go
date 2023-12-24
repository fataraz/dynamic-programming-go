package helper

import "strconv"

// Jadi, big O nya adalah
// O(2**n+m) time
// O(n + m) space
func GridTraveler(m, n int) int {
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}

	return GridTraveler(m-1, n) + GridTraveler(m, n-1)
}

// memoized
// O(m*n) time
// O(n+m) space
func GridTravelerMemorization(m, n int, memo map[string]int) int {
	key := strconv.Itoa(m) + ":" + strconv.Itoa(n)
	if memo[key] != 0 {
		return memo[key]
	}
	if m == 0 || n == 0 {
		return 0
	}
	if m == 1 && n == 1 {
		return 1
	}
	memo[key] = GridTravelerMemorization(m-1, n, memo) + GridTravelerMemorization(m, n-1, memo)

	return memo[key]
}
