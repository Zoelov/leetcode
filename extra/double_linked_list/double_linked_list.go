package double_linked_list

type DNode struct {
	val  interface{}
	pre  *DNode
	next *DNode
}

type DNodeList struct {
	head *DNode
	tail *DNode
	size uint64
}

func InitList() *DNodeList {
	s := new(DNodeList)
	s.size = 0
	s.head = nil
	s.tail = nil
	return s
}

func (d *DNodeList) Append(val interface{}) {
	node := new(DNode)
	node.val = val
	if d.GetSize() == 0 {
		d.head = node
		d.tail = node
		node.pre = nil
		node.next = nil
	} else {
		d.tail.next = node
		node.pre = d.tail
		node.next = nil
		d.tail = node
	}

	d.size++
}

func (d *DNodeList) InsertNext(node *DNode, val interface{}) bool {
	if node == nil {
		return false
	}

	if node == d.tail {
		d.Append(val)
	} else {
		newNode := new(DNode)
		newNode.val = val
		newNode.next = node.next
		newNode.pre = node
		node.next = newNode
		newNode.next.pre = newNode
		d.size++
	}
	return true
}

func (d *DNodeList) GetSize() uint64 {
	return d.size
}
