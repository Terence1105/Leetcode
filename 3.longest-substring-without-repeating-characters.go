/*
 * @lc app=leetcode id=3 lang=golang
 *
 * [3] Longest Substring Without Repeating Characters
 * Input: string which includes substring
 * Output: longest substring without repeating characters.
 * 解法:
   宣告nonDuplicated map用來儲存走過的元素
   接著透過start, end兩個Pointer
   一開始都指向第一個元素
   end依序檢查指向的元素是否有在map裡
   有的話就將start向前移並從map裡移除start指向的元素
   在每次迴圈執行時取end-start最大值 就是答案了
   時間複雜度: O(n^2)
   空間複雜度: O(1)
*/

// @lc code=start
func lengthOfLongestSubstring(s string) int {
	strCount := len(s)
	maxSubStringLen := 0
	nonDuplicated := make(map[byte]int)

	for start, end := 0, 0; end < strCount; {
		if _, isExist := nonDuplicated[s[end]]; isExist {
			delete(nonDuplicated, s[start])
			start++
		} else {
			nonDuplicated[s[end]] = 0
			end++
		}

		if end-start > maxSubStringLen {
			maxSubStringLen = end - start
		}
	}

	return maxSubStringLen
}

// @lc code=end

