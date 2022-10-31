/*
 * @lc app=leetcode id=19 lang=golang
 *
 * [19] Remove Nth Node From End of List
 * Input: Linked list
 * Output: Linked list which remove nth node from the end
 * 解法:
   step1: Declare Listnode[] 記錄走過的節點
   step2: nth-1.next set to nth.Next
   時間複雜度: O(n)
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
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	nodes := []*ListNode{}
	index := 0
	copy := head
	for true == true {
		nodes = append(nodes, copy)
		index++
		if copy.Next == nil {
			break
		}
		copy = copy.Next
	}
	nodesLen := len(nodes)

	if nodesLen == n {
		return head.Next
	}

	nodes[nodesLen-n-1].Next = nodes[nodesLen-n].Next

	return head
}

// @lc code=end

