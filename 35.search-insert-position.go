/*
 * @lc app=leetcode id=35 lang=golang
 *
 * [35] Search Insert Position
 */

// @lc code=start
func searchInsert(nums []int, target int) int {

	right := len(nums) - 1
	left := 0
	middle := right / 2

	for left < right {
		if target == nums[middle] {
			return middle
		} else if target < nums[middle] {
			right = middle - 1
		} else {
			left = middle + 1
		}
		middle = (left + right) / 2
	}

	if target > nums[middle] {
		return middle + 1
	}

	return middle
}

// @lc code=end

