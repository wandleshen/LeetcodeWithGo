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
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
   head := &ListNode{Val: 0}
 
  n1, n2, carry, anchor := 0, 0, 0, head
    for l1 != nil || l2 != nil{
        if l1 == nil {
            n1 = 0
        } else {
            n1, l1 = l1.Val, l1.Next
        }
        if l2 == nil {
            n2 = 0
        } else {
            n2, l2 = l2.Val, l2.Next
        }
        head.Next = &ListNode{Val: (n1 + n2 + carry) % 10}
        carry = (n1 + n2 + carry) / 10
        head = head.Next
    }
    if carry > 0 {
        head.Next = &ListNode{Val: carry, Next: nil}
    }
    return anchor.Next
}
// @lc code=end

