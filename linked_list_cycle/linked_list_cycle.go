package linked_list_cycle

/*********************************************************

	https://leetcode-cn.com/problems/linked-list-cycle-ii/submissions/

	给定一个链表，返回链表开始入环的第一个节点。 如果链表无环，则返回 null。

	说明：不允许修改给定的链表。

	进阶：
	你是否可以不用额外空间解决此题？


*********************************************************************/

type ListNode struct {
	Val  int
	Next *ListNode
}

func DetectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	var has bool
	slow := head.Next
	if slow == nil {
		return nil
	}

	fast := slow.Next
	if fast == nil {
		return nil
	}

	for fast != nil && slow != nil {
		if fast == slow {
			has = true
			break
		}

		slow = slow.Next
		fast = fast.Next
		if fast != nil {
			fast = fast.Next
		}
	}

	// 当存在环时，slow无论走多少步都是在环内，因此当从head开始，与slow相等时，就是环开始的节点
	if has {
		q := head
		for slow != q {
			slow = slow.Next
			q = q.Next
		}

		return q
	}

	return nil
}

func DetectCycle2(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}

	m := make(map[*ListNode]int)
	q := head
	for q != nil {
		m[q]++
		v := m[q]
		if v == 2 {
			return q
		}

		q = q.Next
	}

	return nil
}
