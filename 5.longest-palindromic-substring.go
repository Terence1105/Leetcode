/*
 * @lc app=leetcode id=5 lang=golang
 *
 * [5] Longest Palindromic Substring
 * Input: string
 * Output: longest palindromic substring
 * 解法:
   透過Dynamic Programing
   將問題切成多個小問題
   先檢查小範圍是否是迴文
   再擴大範圍並查看之前的小範圍結果
   來判斷是否為迴文

   時間複雜度: O(n^2)
   空間複雜度: O(n^2)
*/

// @lc code=start
func longestPalindrome(s string) string {
	strLen := len(s)
	dp := make([][]bool, strLen)

	longPalind := s[0:1]

	for i := 0; i < strLen; i++ {
		dp[i] = make([]bool, strLen)
	}

	for i := 0; i < strLen; i++ {
		dp[i][i] = true
		for j := i - 1; j >= 0; j-- {
			if i-j == 1 && s[i] == s[j] {
				dp[i][j] = true
			} else if s[i] == s[j] && dp[i-1][j+1] {
				dp[i][j] = true
			}
			if dp[i][j] && i-j+1 > len(longPalind) {
				longPalind = s[j : i+1]
			}
		}
	}

	return longPalind
}

// @lc code=end

