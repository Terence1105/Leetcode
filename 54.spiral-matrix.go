/*
* @lc app=leetcode id=54 lang=golang
*
* [54] Spiral Matrix
  - Input: An m x n matrix
  - Output: All elements of the matrix in spiral order
  - 解法:
    step1: 實作四個方法 分別是向右/向左/向下/向上走
    step2: 一開始向右走 走到底後再向下 向下到底後再往左走 以此類推
    step3: 走完全部元素即是解答
    時間複雜度: O(n*m)
    空間複雜度: O(n*m)
*/

// @lc code=start
func spiralOrder(matrix [][]int) []int {
	rowLen := len(matrix)
	columnLen := len(matrix[0])

	spiralOrder := make([]int, 0, rowLen*columnLen)
	for i := 0; i < (rowLen+1)/2; i++ {
		GoRight(matrix, &spiralOrder, 0+i, 0+i, columnLen-i)
		if len(spiralOrder) == rowLen*columnLen {
			break
		}
		GoDown(matrix, &spiralOrder, columnLen-1-i, 0+i+1, rowLen-i)
		GoLeft(matrix, &spiralOrder, rowLen-1-i, columnLen-2-i, 0+i+1)
		if len(spiralOrder) == rowLen*columnLen {
			break
		}
		GoUp(matrix, &spiralOrder, 0+i, rowLen-1-i, 0+i)
	}

	return spiralOrder
}

func GoRight(matrix [][]int, so *[]int, index, start, end int) {
	for i := start; i < end; i++ {
		*so = append(*so, matrix[index][i])
	}
}

func GoLeft(matrix [][]int, so *[]int, index, start, end int) {
	for i := start; i >= end; i-- {
		*so = append(*so, matrix[index][i])
	}
}

func GoUp(matrix [][]int, so *[]int, index, start, end int) {
	for i := start; i > end; i-- {
		*so = append(*so, matrix[i][index])
	}
}

func GoDown(matrix [][]int, so *[]int, index, start, end int) {
	for i := start; i < end; i++ {
		*so = append(*so, matrix[i][index])
	}
}

// @lc code=end
