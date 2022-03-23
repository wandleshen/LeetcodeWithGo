/*
 * @lc app=leetcode.cn id=19 lang=golang
 *
 * [19] 删除链表的倒数第 N 个结点
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
	ans := &ListNode{Next: head}
	prev, node1, node2 := ans, head, head
	for node2 != nil {
		if n <= 0 {
			prev = node1
			node1 = node1.Next
		}
		n--
		node2 = node2.Next
	}
	prev.Next = node1.Next
	return ans.Next
}

// @lc code=end

