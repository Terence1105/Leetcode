/*
 * @lc app=leetcode id=11 lang=golang
 *
 * [11] Container With Most Water
 * Input: one dimension int array
 * Output: the max area of water
 * 解法:
   透過Two Pointer(start, end)
   從頭尾開始計算水量
   小的指標往大的移動
   直到兩個指標碰到
   在此期間檢查是否有最大值即是答案

   時間複雜度: O(n)
   空間複雜度: O(1)
*/

// @lc code=start
func maxArea(height []int) int {
	heightLen := len(height)
	start, end := 0, heightLen-1
	maxContainer := 0

	for start < end {
		nowContainer := ChooseSmaller(height[start], height[end]) * (end - start)
		if nowContainer > maxContainer {
			maxContainer = nowContainer
		}

		if height[start] > height[end] {
			end--
		} else {
			start++
		}
	}
	return maxContainer
}

func ChooseSmaller(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

// @lc code=end

