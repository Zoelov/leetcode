package reverse_list

/************************************************

https://leetcode-cn.com/problems/reverse-linked-list/

反转一个单链表。

示例:

输入: 1->2->3->4->5->NULL
输出: 5->4->3->2->1->NULL
进阶:
你可以迭代或递归地反转链表。你能否用两种方法解决这道题？

************************************************/

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseListBest(head *ListNode) *ListNode {
	cur := head
	var pre *ListNode

	for cur != nil {
		cur.Next, pre, cur = pre, cur, cur.Next
	}

	return pre
}

func reverseList(head *ListNode) *ListNode {
	var lastNode *ListNode
	var tail *ListNode
	var last bool

	for head != nil {
		var tmp *ListNode
		if lastNode != nil {
			tmp = head.Next
			if tmp == nil {
				tail.Next = nil
				last = true
			}
			head.Next = lastNode
		}

		lastNode = head

		if tmp != nil {
			head = tmp
		} else {
			if tail == nil {
				tail = head
			}
			if last {
				head = nil
			} else {
				head = head.Next
			}

		}

	}

	return lastNode
}
