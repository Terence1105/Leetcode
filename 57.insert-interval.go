/*
 * @lc app=leetcode id=57 lang=golang
 *
 * [57] Insert Interval
  - Input: an array of non-overlapping intervals and sorted  by start
  - Output: Return intervals after the insertion(nonOverlapping).
  - 解法:
    step1: 從第一個元素開始迭代
    step2: 若元素的end<insert-start 加到結果集裡
    step3: 當end>insert-start 代表已加完前面的元素
	step4: 開始檢查剩餘元素 若start<insert-end代表還在需要Merge的範圍
    step5: 每次檢查都更新insert的start和end 不斷做合併的動作
	step6: 若合併完成 把insert元素加到結果集 接著把剩餘元素也加進去
	step7: 此時的結果集就是解答 return 結果集
    時間複雜度: O(n)
    空間複雜度: O(n*2)
*/

// @lc code=start
func insert(intervals [][]int, newInterval []int) [][]int {
	intervalsLen := len(intervals)
	results := [][]int{}
	i := 0

	for ; i < intervalsLen && intervals[i][1] < newInterval[0]; i++ {
		results = append(results, intervals[i])
	}

	for j := i; j < intervalsLen && intervals[j][0] <= newInterval[1]; j++ {
		newInterval[0] = Min(intervals[j][0], newInterval[0])
		newInterval[1] = Max(intervals[j][1], newInterval[1])
		i++
	}

	results = append(results, newInterval)

	for ; i < intervalsLen; i++ {
		results = append(results, intervals[i])
	}

	return results
}

func Max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func Min(num1, num2 int) int {
	if num1 > num2 {
		return num2
	} else {
		return num1
	}
}

// @lc code=end

