/*
 * @lc app=leetcode id=33 lang=golang
 *
 * [33] Search in Rotated Sorted Array
  - Input: An integer array nums sorted in ascending order
  - Output: Target position
  - 解法:
	使用二分法，並依照情境判斷如何移動左右指標
    時間複雜度: O(log n)
    空間複雜度: O(1)
*/

// @lc code=start
func search(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left <= right {
		middle := (left + right) / 2
		if target == nums[middle] {
			return middle
		} else if nums[middle] < nums[right] {
			if target > nums[right] {
				right = middle - 1
			} else if target < nums[middle] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		} else {
			if target < nums[middle] && target > nums[right] {
				right = middle - 1
			} else {
				left = middle + 1
			}
		}
	}

	return -1
}

// @lc code=end

