/*
 * @lc app=leetcode id=23 lang=golang
 *
 * [23] Merge k Sorted Lists
  - Input: An array of k linked-lists lists, each linked-list is sorted in ascending order.
  - Output: One sorted linked-list
  - 解法:
    step1: Declare *ListNode slice
    step2: 將所有*ListNode節點存進slice
	step3: 將slice裡的元素重新排序
    step3: 將slice裡元素的Next指向下一個節點
	step4: 回傳slice[0]即是答案
    時間複雜度: O(n^2)
    空間複雜度: O(n)
*/

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func mergeKLists(lists []*ListNode) *ListNode {
	sortedLists := []*ListNode{}
	index := 0

	for _, ls := range lists {
		for {
			if ls == nil {
				break
			} else {
				sortedLists = append(sortedLists, ls)
				index++
				ls = ls.Next
			}

		}
	}

	if len(sortedLists) == 0 {
		return nil
	}

	sort.Slice(sortedLists, func(i, j int) bool {
		return sortedLists[i].Val < sortedLists[j].Val
	})

	for i := 0; i < len(sortedLists)-1; i++ {
		sortedLists[i].Next = sortedLists[i+1]
	}

	sortedLists[len(sortedLists)-1].Next = nil

	return sortedLists[0]
}

// @lc code=end

