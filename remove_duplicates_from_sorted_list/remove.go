package remove_duplicates_from_sorted_list

/****************************************************************
	https://leetcode-cn.com/problems/remove-duplicates-from-sorted-list/

	给定一个排序链表，删除所有重复的元素，使得每个元素只出现一次。

	示例 1:

	输入: 1->1->2
	输出: 1->2
	示例 2:

	输入: 1->1->2->3->3
	输出: 1->2->3
****************************************************************/

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteDuplicates(head *ListNode) *ListNode {
	h := head

	for h != nil {
		if h.Next != nil {
			if h.Val == h.Next.Val {
				h.Next = h.Next.Next
				continue
			}
		}

		h = h.Next
	}

	return head
}
