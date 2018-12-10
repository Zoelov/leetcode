package intersection_node

/********************************************************************

	https://leetcode-cn.com/problems/intersection-of-two-linked-lists/

	编写一个程序，找到两个单链表相交的起始节点。



	例如，下面的两个链表：

	A:          a1 → a2
					↘
						c1 → c2 → c3
					↗
	B:     b1 → b2 → b3
	在节点 c1 开始相交。



	注意：

	如果两个链表没有交点，返回 null.
	在返回结果后，两个链表仍须保持原有的结构。
	可假定整个链表结构中没有循环。
	程序尽量满足 O(n) 时间复杂度，且仅用 O(1) 内存。

********************************************************************/

type ListNode struct {
	Val  int
	Next *ListNode
}

func GetIntersectionNode(headA, headB *ListNode) *ListNode {
	var aLen, bLen int
	a := headA
	b := headB

	for a != nil {
		aLen++
		a = a.Next
	}

	for b != nil {
		bLen++
		b = b.Next
	}

	a = headA
	b = headB

	for aLen > bLen {
		a = a.Next
		aLen--
	}

	for bLen > aLen {
		b = b.Next
		bLen--
	}

	for a != b {
		a = a.Next
		b = b.Next
	}

	return a
}
