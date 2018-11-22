package add_two_num

import "testing"

func TestAddTwoNum(t *testing.T) {
	s1 := []int{2, 4, 3}
	s2 := []int{5, 6, 4}

	expectSlice := []int{7, 0, 8}

	node1 := genListNode(s1)
	node2 := genListNode(s2)

	expectNode := genListNode(expectSlice)

	actualSum := addTwoNumbers(node1, node2)

	for expectNode != nil && actualSum != nil {
		if expectNode.Val != actualSum.Val {
			t.Fatal("failed")
			break
		}
		expectNode = expectNode.Next
		actualSum = actualSum.Next
	}

	t.Log("success")
}

func genListNode(nums []int) *ListNode {
	head := new(ListNode)
	var lastNode *ListNode

	for _, v := range nums {
		node := new(ListNode)
		node.Val = v

		if lastNode != nil {
			lastNode.Next = node
		}

		if head == nil {
			head = node
		}

		lastNode = node
	}

	return head
}
