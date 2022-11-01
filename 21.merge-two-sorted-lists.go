/*
 * @lc app=leetcode id=21 lang=golang
 *
 * [21] Merge Two Sorted Lists
  - Input: The heads of two sorted linked lists
  - Output: Merge the two lists in a one sorted list
  - 第一種解法:
    step1: Declare *ListNode slice
    step2: 將list1 & list2每個節點依序存進slice
    step3: 將slice裡元素的Next指向下一個節點
    時間複雜度: O(n)
    空間複雜度: O(n)
 - 第二種解法:
 	step1: 使用recursive找下一個節點
	step2: 若其中一個ListNode為nil 那就return
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
func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	// nodes := []ListNode{}

	// if list1 == nil && list2 == nil {
	// 	return nil
	// }

	// for {

	// 	if list1 == nil && list2 == nil {
	// 		break
	// 	} else if list1 == nil {
	// 		nodes = append(nodes, *list2)
	// 		list2 = list2.Next
	// 	} else if list2 == nil {
	// 		nodes = append(nodes, *list1)
	// 		list1 = list1.Next
	// 	} else if list1.Val > list2.Val {
	// 		nodes = append(nodes, *list2)
	// 		list2 = list2.Next
	// 	} else {
	// 		nodes = append(nodes, *list1)
	// 		list1 = list1.Next
	// 	}
	// }

	// for i := 0; i < len(nodes)-1; i++ {
	// 	nodes[i].Next = &nodes[i+1]
	// }

	// return &nodes[0]

	if list1 == nil {
		return list2
	}
	if list2 == nil {
		return list1
	}

	if list1.Val > list2.Val {
		list2.Next = mergeTwoLists(list1, list2.Next)
		return list2
	} else {
		list1.Next = mergeTwoLists(list1.Next, list2)
		return list1
	}
}

// @lc code=end

