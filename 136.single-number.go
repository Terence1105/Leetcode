/*
 * @lc app=leetcode id=136 lang=golang
 *
 * [136] Single Number
 */

// @lc code=start
func singleNumber(nums []int) int {

	// // METHOD1 :
	// // Step1 Map the given array's elements to their frequency.
	// // Step2 Traverse that map and return the key whose value =1.

	// numsCount := len(nums)
	// duplicated := map[int]int{}

	// for i := 0; i < numsCount; i++ {
	// 	_, isExist := duplicated[nums[i]]
	// 	if isExist == false {
	// 		duplicated[nums[i]] = 1
	// 	} else {
	// 		duplicated[nums[i]]++
	// 	}
	// }

	// for key, value := range duplicated {
	// 	if value == 1 {
	// 		return key
	// 	}
	// }

	// return 0

	//Method2 :
	//USING BITWISE XOR OPERATOR (USING CONSTANT SPACE)
	//The same number do XOP OPERATOR => 00000000
	ans := 0

	for i := 0; i < len(nums); i++ {
		ans ^= nums[i]
	}

	return ans
}

// @lc code=end

