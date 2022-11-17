/*
* @lc app=leetcode id=76 lang=golang
*
* [76] Minimum Window Substring
  - Input: Given two strings s and t of lengths m and n respectively
  - Output: Minimum window substring of s such that every character in t (including duplicates) is included in the window.
  - 解法:
    參考 https://leetcode.com/problems/minimum-window-substring/solution/
    時間複雜度: O(S+T)
    空間複雜度: O(S+T)
*/
package main

import "math"

// @lc code=start
func minWindow(s string, t string) string {

	sLen, tLen := len(s), len(t)

	if sLen < tLen {
		return ""
	}

	left, right, index := 0, 0, 0
	counter, minSubStrLength := tLen, math.MaxInt
	tm := map[byte]int{}

	for _, val := range []byte(t) {
		tm[val]++
	}

	for right < sLen {
		if tm[s[right]] >= 1 {
			counter--
		}

		tm[s[right]]--
		right++

		for counter == 0 {

			if minSubStrLength > right-left {
				index = left
				minSubStrLength = right - left
			}

			tm[s[left]]++
			if tm[s[left]] == 1 {
				counter++
			}
			left++
		}

	}

	if minSubStrLength == math.MaxInt {
		return ""
	}

	return s[index : index+minSubStrLength]
}

// @lc code=end
