/*
 * @lc app=leetcode id=55 lang=golang
 *
 * [55] Jump Game
  - Input: An integer array nums.
  - Output: Return true if you can reach the last index, or false otherwise.
  - 第一種解法:
    Used Greedy Algorithm
    step1: 從0開始迭代 每次都更新當前能跳到的最遠位置
	step2: 如果能走到最後一個元素 回傳true 不行則回傳false
    時間複雜度: O(n)
    空間複雜度: O(1)
  - 第二種解法:
    Used Dynamic promagging
	step1: 將問題切割成每一個點能不能跳到終點
	step2: 從第一個元素開始迭代 接著不斷沿用之前的結果
	step3: 最後一個元素若大於等於 len(nums)-1 回傳 true
	時間複雜度: O(n)
    空間複雜度: O(n)
*/

// @lc code=start
func canJump(nums []int) bool {
	//return GreedyAlgorithm(nums)
	return DynamicPromagging(nums)
}

func DynamicPromagging(nums []int) bool {
	numsLen := len(nums)
	dp := make([]int, numsLen)

	for i := 1; i < numsLen; i++ {
		if dp[i-1] > nums[i-1] {
			dp[i] = dp[i-1] - 1
		} else {
			dp[i] = nums[i-1] - 1
		}

		if dp[i] < 0 {
			return false
		}
	}

	return true
}

func GreedyAlgorithm(nums []int) bool {
	numsLen := len(nums)
	i := 0
	for reach := 0; i < numsLen && i <= reach; i++ {
		if nums[i]+i > reach {
			reach = nums[i] + i
		}
	}

	if i == numsLen {
		return true
	} else {
		return false
	}
}

// @lc code=end

