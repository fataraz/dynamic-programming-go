package helper

// Input: nums = [2,7,11,15], target = 9
// Output: [0,1]
// Explanation: Because nums[0] + nums[1] == 9, we return [0, 1].

// Input: nums = [3,2,4], target = 6
// Output: [1,2]
func TwoSum(nums []int, target int) []int {
	// Space: O(n)
	twoSum := make(map[int]int)

	// Time: O(n)
	for i, num := range nums {
		if val, found := twoSum[target-num]; found {
			return []int{val, i}
		}

		twoSum[num] = i
	}

	return []int{}
}
