/*
 * @lc app=leetcode id=48 lang=golang
 *
 * [48] Rotate Image
  - Input: An n x n 2D matrix representing an image.

  - Output: rotate the image by 90 degrees (clockwise).
  - Limited: in-place

  - 解法:
	step1: reverse up to down
	step2: swap the symmetry
    時間複雜度: O(n^2)
    空間複雜度: O(1)
*/

// @lc code=start
func rotate(matrix [][]int) {

	matrixLen := len(matrix)

	for i := 0; i < matrixLen/2; i++ {
		matrix[i], matrix[matrixLen-1-i] = matrix[matrixLen-1-i], matrix[i]
	}

	for i := 0; i < matrixLen; i++ {
		for j := i + 1; j < matrixLen; j++ {
			matrix[i][j], matrix[j][i] = matrix[j][i], matrix[i][j]
		}
	}
}

// @lc code=end

