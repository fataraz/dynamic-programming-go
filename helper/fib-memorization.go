package helper

var memo = make(map[int]int)

func FibMemorization(n int) int {
	if n <= 1 {
		return n
	}
	// check value in map
	if v, ok := memo[n]; ok {
		return v
	}
	memo[n] = FibMemorization(n-2) + FibMemorization(n-1)

	return memo[n]
}
