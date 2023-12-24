package helper

// x = 451  return 154

func Palindrome(x int) int {
	var reversedNum int
	tmp := x
	for tmp > 0 {
		reversedNum = reversedNum*10 + tmp%10
		tmp = tmp / 10
	}

	return reversedNum
}
