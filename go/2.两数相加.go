/*
 * @lc app=leetcode.cn id=2 lang=golang
 *
 * [2] 两数相加
 */

// @lc code=start
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	dummy := new(ListNode)
// 	curr := dummy
// 	carry := 0

// 	for(l1 != nil || l2 != nil || carry > 0){
// 		curr.Next = new(ListNode)
// 		curr = curr.Next
// 		if l1 != nil {
// 			carry += l1.Val
// 			l1 = l1.Next
// 		}
// 		if l2 != nil{
// 			carry += l2.Val
// 			l2 = l2.Next
// 		}
// 		curr.Val = carry %10
// 		carry /= 10
// 	}
// 	return dummy.Next
// }
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var newList *ListNode = new(ListNode)
	temp := 0
	curr := newList
	for l1 != nil || l2 != nil || temp != 0 {
		curr.Next = new(ListNode)
		curr = curr.Next
		if l1 != nil {
			temp += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			temp += +l2.Val
			l2 = l2.Next
		}
		curr.Val = temp % 10
		temp = temp / 10
	}
	return newList.Next
}

// @lc code=end

