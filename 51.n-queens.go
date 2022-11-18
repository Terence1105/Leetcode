/*
* @lc app=leetcode id=51 lang=golang
*
* [51] N-Queens
  - Input: Given an integer n
  - Output: Return all distinct solutions to the n-queens puzzle
  - 解法:
    Use backtracking
    step1: 由於每一row只能有一個皇后，所以path為由[curRow][0] -> [curRow][n]
    step2: 宣告一個二維Slice紀錄皇后已走過的區域，如果還沒走過就將皇后放下
    step3: 並更新皇后走過的區域，接著持續向下一Row遞迴
    step3: 如果當前放置的皇后個數=n 將此結果加到結果集中 並等全部可能都走完就回傳
    時間複雜度: O(?)
    空間複雜度: O(?)
*/
package main

// @lc code=start
func solveNQueens(n int) [][]string {
	results := [][]string{}

	quennsRange := make([][]bool, n)

	for i := 0; i < n; i++ {
		quennsRange[i] = make([]bool, n)
	}

	quennsPosition := map[int]int{}

	backtracking(&results, n, 0, quennsRange, &quennsPosition)

	return results
}

func backtracking(results *[][]string, n, culRow int, isQuennsRange [][]bool, queensPosition *map[int]int) {
	if len(*queensPosition) == n {
		result := AddResult(queensPosition)
		*results = append(*results, result)
	}

	for j := 0; j < n && culRow < n; j++ {
		if isQuennsRange[culRow][j] == false {
			copyQR := DeepCopyQuennsRange(isQuennsRange, culRow, j, n)
			copyQP := DeepCopyQueensPosition(queensPosition, culRow, j)
			backtracking(results, n, culRow+1, copyQR, &copyQP)
		}
	}
}

func DeepCopyQueensPosition(queensPosition *map[int]int, row, column int) map[int]int {
	copyQP := make(map[int]int, len(*queensPosition)+1)

	for k, v := range *queensPosition {
		copyQP[k] = v
	}

	copyQP[row] = column

	return copyQP
}

func DeepCopyQuennsRange(isQuennsRange [][]bool, row, column, n int) [][]bool {

	copyQR := make([][]bool, n)

	for i := 0; i < n; i++ {
		copyQR[i] = make([]bool, n)
		for j := 0; j < n; j++ {
			copyQR[i][j] = isQuennsRange[i][j]
		}
	}

	for i := row; i < n; i++ {
		copyQR[i][column] = true
	}

	for i, j := row+1, column+1; i < n && j < n; {
		copyQR[i][j] = true
		i++
		j++
	}

	for i, j := row+1, column-1; i < n && j >= 0; {
		copyQR[i][j] = true
		i++
		j--
	}

	return copyQR
}

func AddResult(queensPosition *map[int]int) []string {
	queensCount := len(*queensPosition)
	result := make([]string, queensCount)

	for i := 0; i < queensCount; i++ {
		for j := 0; j < queensCount; j++ {
			result[i] += "."
		}
	}

	for k, v := range *queensPosition {
		str := ""
		for i := 0; i < queensCount; i++ {
			if i == v {
				str += "Q"
			} else {
				str += "."
			}
		}
		result[k] = str
	}

	return result
}

// @lc code=end
