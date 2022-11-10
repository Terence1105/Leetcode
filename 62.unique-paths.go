/*
 * @lc app=leetcode id=62 lang=golang
 *
 * [62] Unique Paths
  - Input: An m x n grid.
  - Output: the number of possible unique paths.
  - 第一種解法:(TimeOut)
    Use backtracking algorithm
	step1: 從起點出發 每一個位置都遞迴向左或向下走
    step2: 當走到左下角 就把當前元素加到結果集裡
    step3: 回傳所有可能的路徑數量
    時間複雜度: O(2^n)
    空間複雜度: O(1)
  - 第二種解法
    Use dynamic programming
	step1: 每一格的走法會等於把上一格和左邊一格的走法加起來
	step2: 一開始先向右走並同時記錄有多少種走法
	step2: 向右走到底就回到下一列第一個元素繼續向右走
	step3: 最後 dp[m-1][n-1]就是答案了
	時間複雜度: O(n^2)
	空間複雜度: O(n^2)
*/

// @lc code=start
func uniquePaths(m int, n int) int {
	uniPathCount := 0
	//backtracking(m, n, 0, 0, &uniPathCount)
	uniPathCount = DynamicProgramming(m, n)
	return uniPathCount
}

func DynamicProgramming(m, n int) int {

	dp := make([][]int, m)

	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			currentPathCount := 0

			if i == 0 && j == 0 {
				dp[i][j] = 1
				continue
			}

			if i-1 >= 0 {
				currentPathCount += dp[i-1][j]
			}

			if j-1 >= 0 {
				currentPathCount += dp[i][j-1]
			}
			dp[i][j] = currentPathCount
		}
	}

	return dp[m-1][n-1]
}

//It will be time limit exceeded
// func BackTracking(m, n, curRow, curCol int, upc *int) {
// 	if curRow == m-1 && curCol == n-1 {
// 		*upc += 1
// 	}

// 	if curRow < m {
// 		backtracking(m, n, curRow+1, curCol, upc)
// 	}

// 	if curCol < n {
// 		backtracking(m, n, curRow, curCol+1, upc)
// 	}
// }

// @lc code=end

