/*
* @lc app=leetcode id=39 lang=golang
*
* [39] Combination Sum

  - Input: an array of distinct integers(nonsorted)

  - Output: Combinations(sum equal target)

  - 解法:
	Use backtracking algorithm.
	step1: 路徑要填的即為選取的 candidates
	step2: 選擇清單為candidates當前和之後的元素
	step3: 終止條件即路徑內的總和大於等於 target 
    時間複雜度: O(2^n)
    空間複雜度: O(2^n)
*/

// @lc code=start
func combinationSum(candidates []int, target int) [][]int {

	results := [][]int{}

	if len(candidates) == 0 {
		return results
	}

	backtrack(&results, candidates, []int{}, 0, target)

	return results
}
 result []int
func backtrack(results *[][]int, nums []int,, start_index int, target int) {
	if target < 0 {
		return
	}

	if target == 0 {
		temp := make([]int, len(result))
		copy(temp, result)
		*results = append(*results, temp)
	}

	for i := start_index; i < len(nums); i++ {
		result = append(result, nums[i])
		backtrack(results, nums, result, i, target-nums[i])
		result = result[:len(result)-1]
	}
}

// @lc code=end
