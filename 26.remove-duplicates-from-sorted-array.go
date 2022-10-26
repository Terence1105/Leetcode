/*
 * @lc app=leetcode id=26 lang=golang
 *
 * [26] Remove Duplicates from Sorted Array
 */

// @lc code=start
func removeDuplicates(nums []int) int {
	arrCount := len(nums)
	startPointer := 0

	for i := 1; i < arrCount; i++ {
		if nums[startPointer] != nums[i] {
			startPointer++
			nums[startPointer] = nums[i]
		}
	}

	return startPointer + 1
}

// @lc code=end

