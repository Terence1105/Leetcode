/*
* @lc app=leetcode id=79 lang=golang
*
  - [79] Word Search
  - Input: Given an m x n grid of characters board and a string word
  - Output: Return true if word exists in the grid
  - 解法:
    Use backtracking and pruning
    step1: 從每個位置向前後左右出發
    step2: 若不符合 word 就 purning
    step3: 若有找到回傳true
    時間複雜度: O(?)
    空間複雜度: O(?)
*/
package main

// @lc code=start

func exist(board [][]byte, word string) bool {

	row := len(board)
	col := len(board[0])
	isFind := false

	if row*col < len(word) {
		return false
	}

	for i := 0; i < row; i++ {
		for j := 0; j < col; j++ {
			backtracking(board, word, i, j, 0, &isFind)
		}
	}

	return isFind
}

func backtracking(board [][]byte, word string, row, col, wordPosition int, isFind *bool) {

	br := len(board)
	bc := len(board[0])

	if wordPosition == len(word) {
		*isFind = true
		return
	}

	if row < 0 || col < 0 || row == br || col == bc {
		return
	}

	if board[row][col] != word[wordPosition] {
		return
	}

	c := board[row][col]
	board[row][col] = '*'
	backtracking(board, word, row-1, col, wordPosition+1, isFind)
	backtracking(board, word, row, col-1, wordPosition+1, isFind)
	backtracking(board, word, row+1, col, wordPosition+1, isFind)
	backtracking(board, word, row, col+1, wordPosition+1, isFind)
	board[row][col] = c
}

// @lc code=end
