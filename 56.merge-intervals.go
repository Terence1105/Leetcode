/*
  - @lc app=leetcode id=56 lang=golang
    *
  - [56] Merge Intervals
  - Input: An array of intervals where intervals[i] = [starti, endi]
  - Output: An array of the non-overlapping intervals
  - 解法:
    step1: 宣告一個map和slice元素放置intervals每個元素的第一個值
    step2: 將slice依照由小到大重新排序
    step3: 再透過map對應回intervals一個一個檢查是否有overlapping
    step3: 回傳檢查好的 2-D slice 即是解答
    時間複雜度: O(nlogn)
    空間複雜度: O(n)
*/
package main

import (
	"fmt"
	"sort"
)

func main() {
	test := [][]int{
		[]int{2, 3},
		[]int{5, 5},
		[]int{2, 2},
		[]int{3, 4},
		[]int{3, 4},
	}
	fmt.Println(merge(test))
}

// @lc code=start
func merge(intervals [][]int) [][]int {
	//return DIYSort(intervals)
	return SortSlice(intervals)
}

func SortSlice(intervals [][]int) [][]int {

	nonOverlapping := [][]int{}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	nonOverlapping = append(nonOverlapping, intervals[0])
	for i, last := 1, 0; i < len(intervals); i++ {
		if nonOverlapping[last][1] < intervals[i][0] {
			nonOverlapping = append(nonOverlapping, intervals[i])
			last++
		} else if nonOverlapping[last][1] <= intervals[i][1] {
			nonOverlapping[last][1] = intervals[i][1]
		}
	}
	return nonOverlapping
}

func DIYSort(intervals [][]int) [][]int {

	intervalsLen := len(intervals)

	if intervalsLen <= 1 {
		return intervals
	}

	sorted := make([]int, intervalsLen)
	mapping := make(map[int]int, intervalsLen)

	nonOverlapping := [][]int{}

	for i := 0; i < intervalsLen; i++ {
		sorted[i] = intervals[i][0]
		if _, isExists := mapping[intervals[i][0]]; isExists {
			if intervals[mapping[intervals[i][0]]][1] < intervals[i][1] {
				mapping[intervals[i][0]] = i
			}
		} else {
			mapping[intervals[i][0]] = i
		}
	}

	sort.Ints(sorted)

	nonOverlapping = append(nonOverlapping, intervals[mapping[sorted[0]]])

	for i, last := 1, 0; i < intervalsLen; i++ {
		if nonOverlapping[last][1] >= intervals[mapping[sorted[i]]][1] {
			continue
		} else if nonOverlapping[last][1] >= intervals[mapping[sorted[i]]][0] {
			nonOverlapping[last][1] = intervals[mapping[sorted[i]]][1]
		} else {
			nonOverlapping = append(nonOverlapping, intervals[mapping[sorted[i]]])
			last++
		}
	}

	return nonOverlapping
}

// @lc code=end
