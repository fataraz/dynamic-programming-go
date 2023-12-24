package helper

// write function that takes in a number as an argument.
// the function should return the n-th number of the Fibonacci sequence.
// the 0th number of the sequence is 0.
// the 1st number of the sequence is 1.
// to generate the next number of the sequence, we sum the previous two.

func FibTabulation(n int) int {
	table := make([]int, n+1)
	if n <= 1 {
		return n
	}

	table[0] = 0
	table[1] = 1
	for i := 2; i <= n; i++ {
		table[i] += table[i-1] + table[i-2]
	}

	return table[n]
}
