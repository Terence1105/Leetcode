/*
 * @lc app=leetcode id=15 lang=golang
 *
 * [15] 3Sum
 * Input: Nonsorted integer array and allow duplicated
 * Output: triplets array
 * 解法:
   Step1: Sort input array
   Step2: Transfer all element at i = 0 and use two pointer left right
	      left = i+1, right = numsLen-1
   step3: nums[i] + nums[left] + nums[right] == 0
		  將結果先暫存到Output array
   step4: if i < arrLen-2 時回傳output array
   時間複雜度: O(n^2)
   空間複雜度: O(n)
*/

// @lc code=start
func threeSum(nums []int) [][]int {
	numsLen := len(nums)
	answer := [][]int{}

	sort.Ints(nums)

	for i := 0; i < numsLen-2; i++ {
		left, right := i+1, numsLen-1
		target := -nums[i]
		for left < right {

			if nums[left]+nums[right] > target {
				right--
			} else if nums[left]+nums[right] < target {
				left++
			} else {
				a := []int{nums[i], nums[left], nums[right]}
				answer = append(answer, a)

				for ; left < right && nums[left] == a[1]; left++ {
				}

				for ; left < right && nums[right] == a[2]; right-- {
				}
			}
		}

		for ; i < numsLen-1 && nums[i] == nums[i+1]; i++ {
		}
	}

	return answer
}

// @lc code=end

