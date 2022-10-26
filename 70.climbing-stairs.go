/*
 * @lc app=leetcode id=70 lang=golang
 *
 * [70] Climbing Stairs
 */

// @lc code=start
func climbStairs(n int) int {
	lastSecondStair := 1
	lastOneStair := 2
	nowStair := 0

	if n == 1 {
		return 1
	} else if n == 2 {
		return 2
	}

	for i := 3; i <= n; i++ {
		nowStair = lastOneStair + lastSecondStair
		lastSecondStair = lastOneStair
		lastOneStair = nowStair
	}
	return nowStair
}

// @lc code=end

// 1 : 1
// 2 : 2
// 3 : 3
// 4 : 5
// 5 : 8
// 6 : 13