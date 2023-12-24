package main

import (
	"dynamic-programming-go/helper"
	"fmt"
)

func main() {
	//fmt.Println(helper.FibMemorization(50))
	//fmt.Println(helper.GridTraveler(2, 3))
	//
	//memo := make(map[string]int)
	//fmt.Println(helper.GridTravelerMemorization(2, 3, memo)) // 3
	//
	//memoCanSum := make(map[int]bool)
	////fmt.Println(helper.CanSum(300, []int{2, 14}))
	//fmt.Println(helper.CanSumMemorization(300, []int{2, 14}, memoCanSum))

	//fmt.Println(helper.HowSum(7, []int{2, 3}))
	//fmt.Println(helper.HowSum(7, []int{2, 4}))
	//fmt.Println(helper.HowSum(7, []int{3, 5, 2}))
	//fmt.Println(helper.HowSumMemorization(1000, []int{3, 5, 2}))

	//fmt.Println(helper.BestSum(7, []int{5, 3, 4, 7}))
	//fmt.Println(helper.BestSum(8, []int{2, 3, 5}))
	//
	//fmt.Println("Best Sum Memorization")
	//fmt.Println(helper.BestSumMemorization(7, []int{5, 3, 4, 7}))
	//fmt.Println(helper.BestSumMemorization(8, []int{2, 3, 5}))
	//
	//fmt.Println(helper.CanConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	////lorem := "enterapotentpot"
	////fmt.Println(lorem[1:]) // "nterapotentpot
	////fmt.Println(lorem[5:]) // "apotentpot"
	//
	////fmt.Println(lorem[:1]) // "e"
	//fmt.Println("CanConstruct Memorization")
	//fmt.Println(helper.CanConstructMemorization("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	//fmt.Println(helper.CanConstructMemorization("abcdef", []string{"ab", "abc", "cd", "def", "abcd"}))

	//fmt.Println(helper.TwoSum([]int{2, 7, 11, 15}, 9))
	//fmt.Println(helper.TwoSum([]int{3, 2, 4}, 6))

	//fmt.Println(helper.Palindrome(321))

	//fmt.Println(helper.CountConstruct("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))
	//fmt.Println(helper.CountConstructMemorization("enterapotentpot", []string{"a", "p", "ent", "enter", "ot", "o", "t"}))

	//fmt.Println(helper.AllConstructMemorized("skateboard", []string{"bo", "rd", "ate", "t", "ska", "sk", "boar"}))
	//fmt.Println(helper.AllConstructMemorized("abcdef", []string{"ab", "abc", "cd", "def", "abcd", "ef", "c"}))

	fmt.Println(helper.FibTabulation(6))
}
