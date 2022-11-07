/*
 * @lc app=leetcode id=53 lang=golang
 *
 * [53] Maximum Subarray
  - Input: An integer array nums

  - Output: The subarray which has the largest sum

  - 解法:
	step1: 使用Dynamic Programing Algorithms
	step2: 將問題切成多個小問題 -> 起點到當前位置的總和
	step3: 因此每一個點可以拿前一個點的總和來計算
    時間複雜度: O(n)
    空間複雜度: O(n)
*/

// @lc code=start
func maxSubArray(nums []int) int {

	numsLen := len(nums)
	dp := make([]int, numsLen)
	dp[0] = nums[0]
	maxSum := dp[0]

	for i := 1; i < numsLen; i++ {
		if dp[i-1] < 0 {
			dp[i] = nums[i]
		} else {
			dp[i] = dp[i-1] + nums[i]
		}

		if maxSum < dp[i] {
			maxSum = dp[i]
		}
	}

	return maxSum
}

// @lc code=end

