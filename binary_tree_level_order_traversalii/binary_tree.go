package binary_tree_level_order_traversalii

/***********************************************

	https://leetcode-cn.com/problems/binary-tree-level-order-traversal-ii/
	给定一个二叉树，返回其节点值自底向上的层次遍历。 （即按从叶子节点所在层到根节点所在的层，逐层从左向右遍历）

	例如：
	给定二叉树 [3,9,20,null,null,15,7],

		3
	/ \
	9  20
		/  \
	15   7
	返回其自底向上的层次遍历为：

	[
	[15,7],
	[9,20],
	[3]
	]

***********************************************/
type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	q := queue{}

	ret := [][]int{}

	d := Data{root, 0}
	q.EnQueue(d)

	for !q.IsEmpty() {
		v := q.DeQueue()

		if len(ret) == v.level {
			ret = append(ret, []int{})
		}

		ret[v.level] = append(ret[v.level], v.node.Val)

		if v.node.Left != nil {
			d := Data{
				v.node.Left,
				v.level + 1,
			}
			q.EnQueue(d)
		}

		if v.node.Right != nil {
			d := Data{
				v.node.Right,
				v.level + 1,
			}
			q.EnQueue(d)
		}

	}

	i, j := 0, len(ret)-1

	for i < j {
		ret[i], ret[j] = ret[j], ret[i]
		i++
		j--
	}

	return ret
}

type Data struct {
	node  *TreeNode
	level int
}

type queue []Data

func (q *queue) EnQueue(n Data) {
	*q = append(*q, n)
}

func (q *queue) DeQueue() Data {
	v := (*q)[0]

	*q = (*q)[1:]
	return v
}

func (q queue) IsEmpty() bool {
	return len(q) == 0
}
