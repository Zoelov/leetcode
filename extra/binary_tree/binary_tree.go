package main

import (
	"fmt"
	"leetcode/extra/queue"
	"leetcode/extra/stack"
	"math"
)

type BinaryTree struct {
	val   int
	left  *BinaryTree
	right *BinaryTree
}

// 前序：{1,2,4,7,3,5,6,8} 中序：{4,7,2,1,5,3,8,6}

func Construct(pre []int, mid []int, length int) *BinaryTree {
	if pre == nil || mid == nil || length <= 0 {
		return nil
	}

	root := new(BinaryTree)
	root.val = pre[0]

	if length == 1 && pre[0] == mid[0] {
		return root
	}

	leftNum := 0

	for i := 0; i < length; i++ {
		if mid[i] != pre[0] {
			leftNum++
		} else {
			break
		}
	}

	if leftNum > 0 {
		root.left = Construct(pre[1:], mid[0:leftNum], leftNum)
	}

	if length-leftNum > 0 {
		root.right = Construct(pre[leftNum+1:], mid[leftNum+1:], length-leftNum-1)
	}

	return root

}

func (root *BinaryTree) PreorderTravel() []int {
	ret := make([]int, 0)
	ret = append(ret, root.val)

	if root.left != nil {
		t := root.left.PreorderTravel()
		ret = append(ret, t...)
	}

	if root.right != nil {
		t := root.right.PreorderTravel()
		ret = append(ret, t...)
	}

	return ret
}

func (root *BinaryTree) PreorderTravelIteration() []int {
	ret := make([]int, 0)

	stack := new(stack.Stack)
	stack.Push(root)

	for !stack.IsEmpty() {
		n, _ := stack.Pop()
		ret = append(ret, n.(*BinaryTree).val)

		if n.(*BinaryTree).right != nil {
			stack.Push(n.(*BinaryTree).right)
		}

		if n.(*BinaryTree).left != nil {
			stack.Push(n.(*BinaryTree).left)
		}
	}

	return ret
}

func (root *BinaryTree) InOrderTravel() []int {
	ret := make([]int, 0)
	if root.left != nil {
		left := root.left.InOrderTravel()
		ret = append(ret, left...)
	}

	ret = append(ret, root.val)

	if root.right != nil {
		right := root.right.InOrderTravel()
		ret = append(ret, right...)
	}

	return ret
}

func (root *BinaryTree) InOrderTravelIteration() []int {
	ret := make([]int, 0)

	stack := new(stack.Stack)

	cur := root
	for !stack.IsEmpty() || cur != nil {
		for cur != nil {
			stack.Push(cur)

			cur = cur.left
		}

		n, _ := stack.Pop()
		ret = append(ret, n.(*BinaryTree).val)
		cur = n.(*BinaryTree).right
	}

	return ret
}

func (root *BinaryTree) PostOrderTravel() []int {
	ret := make([]int, 0)

	if root.left != nil {
		l := root.left.PostOrderTravel()
		ret = append(ret, l...)
	}

	if root.right != nil {
		r := root.right.PostOrderTravel()
		ret = append(ret, r...)
	}

	ret = append(ret, root.val)

	return ret
}

func (root *BinaryTree) PostOrderTravelIteration() []int {
	ret := make([]int, 0)
	stack := new(stack.Stack)
	cur := root
	var pre *BinaryTree

	stack.Push(root)
	for !stack.IsEmpty() {
		n, _ := stack.Top()
		cur = n.(*BinaryTree)
		if (cur.left == nil && cur.right == nil) || (pre != nil && (pre == cur.left || pre == cur.right)) {
			ret = append(ret, cur.val)
			stack.Pop()
			pre = cur
		} else {
			if cur.right != nil {
				stack.Push(cur.right)
			}

			if cur.left != nil {
				stack.Push(cur.left)
			}
		}
	}

	return ret
}

func (root *BinaryTree) BreadthTravel() []int {
	ret := make([]int, 0)
	q := new(queue.Queue)

	q.EnQueue(root)
	for !q.IsEmpty() {
		v, _ := q.DeQueue()
		ret = append(ret, v.(*BinaryTree).val)

		if v.(*BinaryTree).left != nil {
			q.EnQueue(v.(*BinaryTree).left)
		}

		if v.(*BinaryTree).right != nil {
			q.EnQueue(v.(*BinaryTree).right)
		}
	}

	return ret
}

func (root *BinaryTree) NodeNum() int {
	if root == nil {
		return 0
	}

	return root.left.NodeNum() + root.right.NodeNum() + 1
}

func (root *BinaryTree) Depth() int {
	if root == nil {
		return 0
	}

	leftDepth := root.left.Depth()
	rightDepth := root.right.Depth()
	max := math.Max(float64(leftDepth), float64(rightDepth))

	return int(max) + 1
}

func GetNodeNumKthLevel(root *BinaryTree, k int) int {
	if root == nil || k < 1 {
		return 0
	}

	if k == 1 {
		return 1
	}

	leftNum := GetNodeNumKthLevel(root.left, k-1)
	rightNum := GetNodeNumKthLevel(root.right, k-1)

	return leftNum + rightNum
}

func GetLeafNodeNum(root *BinaryTree) int {
	if root == nil {
		return 0
	}

	if root.left == nil && root.right == nil {
		return 1
	}

	leftNum := GetLeafNodeNum(root.left)
	rightNum := GetLeafNodeNum(root.right)

	return leftNum + rightNum
}

func FindNode(root, node *BinaryTree) bool {
	if root == nil || node == nil {
		return false
	}

	if root == node {
		return true
	}

	found := FindNode(root.left, node)
	if !found {
		found = FindNode(root.right, node)
	}

	return found
}

func GetLastCommonParent(root, node1, node2 *BinaryTree) *BinaryTree {
	if FindNode(root.left, node1) {
		if FindNode(root.right, node2) {
			return root
		}

		return GetLastCommonParent(root.left, node1, node2)
	}

	if FindNode(root.left, node2) {
		return root
	}

	// l := list.New()
	return GetLastCommonParent(root.right, node1, node2)

}

func GetPath(root, node *BinaryTree, q *queue.Queue) bool {
	if root == nil {
		return false
	}

	q.EnQueue(root)

	var found bool
	if root == node {
		return true
	}

	if !found && root.left != nil {
		found = GetPath(root.left, node, q)
	}

	if !found && root.right != nil {
		found = GetPath(root.right, node, q)
	}

	if !found {
		q.DeQueueBack()
	}

	return found
}

func main() {
	pre := []int{1, 2, 4, 7, 3, 5, 6, 8}
	mid := []int{4, 7, 2, 1, 5, 3, 8, 6}
	// post := []int{7, 4, 2, 5, 8, 6, 3, 1}

	r := Construct(pre, mid, len(pre))

	v := r.PreorderTravel()
	v1 := r.PreorderTravelIteration()
	fmt.Println(v)
	fmt.Println(v1)
	fmt.Println("...................")
	m := r.InOrderTravel()
	m1 := r.InOrderTravelIteration()
	fmt.Println(m)
	fmt.Println(m1)
	fmt.Println("...................")

	p := r.PostOrderTravel()
	p1 := r.PostOrderTravelIteration()
	fmt.Println(p)
	fmt.Println(p1)
	fmt.Println("...................")

	w := r.BreadthTravel()
	fmt.Println(w)

	n := r.NodeNum()
	fmt.Println("num:", n)

	d := r.Depth()
	fmt.Println("depth:", d)

	n = GetNodeNumKthLevel(r, 2)
	fmt.Println("n:", n)

	n = GetLeafNodeNum(r)
	fmt.Println("leaf num:", n)

	node := GetLastCommonParent(r, r.right.left, r.right.right.left)
	fmt.Println("common:", node.val)

	q := new(queue.Queue)
	ok := GetPath(r, r.right.right.left, q)

	for ok && !q.IsEmpty() {
		v, _ := q.DeQueue()
		fmt.Println(v.(*BinaryTree).val)
	}

}
