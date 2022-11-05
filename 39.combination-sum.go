/*
* @lc app=leetcode id=39 lang=golang
*
* [39] Combination Sum

  - Input: an array of distinct integers(nonsorted)

  - Output: Combinations(sum equal target)

  - 解法:

    時間複雜度: O(log n)
    空間複雜度: O(1)
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

func backtrack(results *[][]int, nums []int, result []int, start_index int, target int) {
	if target < 0 {
		return
	}

	if target == 0 {
		// temp := make([]int, len(result))
		// copy(temp, result)
		*results = append(*results, result)
	}

	for i := start_index; i < len(nums); i++ {
		result = append(result, nums[i])
		backtrack(results, nums, result, i, target-nums[i])
		result = result[:len(result)-1]
	}
}

// @lc code=end
