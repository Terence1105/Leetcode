/*
 * @lc app=leetcode id=61 lang=golang
 *
 * [61] Rotate List
  - Input: head of a linked list

  - Output: rotate the list to the right by k places.
  - 第一種解法:
	step1: 將所有ListNode加到slice裡
	step2: 設定 k % len(slice) 的next value等於nil
	step3: 設定 len(slice) -1的next value等於slice[0]
	step4: 回傳slice[k % len(slice) + 1]
    時間複雜度: O(n)
    空間複雜度: O(n)

	第二種解法:
	step1: 計算len(head)，並獲得最後一個元素的pointer(tail)
	step2: 將最後一個元素的Next等於head
	step3: 將head和tail進行rotate len-k次
	step4: 設定tail.next = nil 並回傳head即是答案
	時間複雜度: O(n)
	空間複雜度: O(1)
*/

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {

	if head == nil {
		return nil
	}

	// listnodes := []*ListNode{}

	// for temp := head; temp != nil; temp = temp.Next {
	// 	listnodes = append(listnodes, temp)
	// }

	// headLen := len(listnodes)
	// rotatePoint := k % headLen

	// if rotatePoint != 0 {
	// 	listnodes[headLen-rotatePoint-1].Next = nil
	// 	listnodes[headLen-1].Next = listnodes[0]
	// 	return listnodes[headLen-rotatePoint]
	// }

	// return head

	headLen := 1
	tail := head

	for ; tail.Next != nil; tail = tail.Next {
		headLen++
	}

	tail.Next = head
	rotateTimes := k % headLen

	if rotateTimes != 0 {
		for i := 0; i < headLen-rotateTimes; i++ {
			head = head.Next
			tail = tail.Next
		}
	}

	tail.Next = nil

	return head
}

// @lc code=end

