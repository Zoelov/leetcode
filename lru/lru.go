package lru

/**********************************************************************************

	https://leetcode-cn.com/explore/interview/card/bytedance/245/data-structure/1032/
	运用你所掌握的数据结构，设计和实现一个  LRU (最近最少使用) 缓存机制。它应该支持以下操作： 获取数据 get 和 写入数据 put 。

	获取数据 get(key) - 如果密钥 (key) 存在于缓存中，则获取密钥的值（总是正数），否则返回 -1。
	写入数据 put(key, value) - 如果密钥不存在，则写入其数据值。当缓存容量达到上限时，它应该在写入新数据之前删除最近最少使用的数据值，从而为新的数据值留出空间。

	进阶:

	你是否可以在 O(1) 时间复杂度内完成这两种操作？

	示例:

	LRUCache cache = new LRUCache( 2  );

	cache.put(1, 1);
	cache.put(2, 2);
	cache.get(1);       // 返回  1
	cache.put(3, 3);    // 该操作会使得密钥 2 作废
	cache.get(2);       // 返回 -1 (未找到)
	cache.put(4, 4);    // 该操作会使得密钥 1 作废
	cache.get(1);       // 返回 -1 (未找到)
	cache.get(3);       // 返回  3
	cache.get(4);       // 返回  4

***********************************************************************************/

type LRUCache struct {
	capacity int
	used     int
	data     map[int]int
	nodeMap  map[int]*doubleLinkedNode
	doubleLinkedList
}

type doubleLinkedNode struct {
	key  int
	pre  *doubleLinkedNode
	next *doubleLinkedNode
}

type doubleLinkedList struct {
	head *doubleLinkedNode
	tail *doubleLinkedNode
}

func (d *doubleLinkedList) delete(node *doubleLinkedNode) {
	if d.head == node && d.tail == node {
		d.head = nil
		d.tail = nil
	} else if d.head == node {
		d.head.next.pre = nil
		d.head = d.head.next
	} else if d.tail == node {
		d.tail.pre.next = nil
		d.tail = d.tail.pre
	} else {
		node.pre.next = node.next
		node.next.pre = node.pre
	}
}

func (d *doubleLinkedList) deleteTail() {
	if d.head == nil && d.tail == nil {
		return
	}

	if d.head == d.tail {
		d.head = nil
		d.tail = nil
	} else {
		d.tail.pre.next = nil
		d.tail = d.tail.pre
	}
}

func (d *doubleLinkedList) insertBeforHead(node *doubleLinkedNode) {
	if d.head == nil && d.tail == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.pre = node
		d.head = node
	}
}

func Constructor(capacity int) LRUCache {
	lru := LRUCache{
		capacity: capacity,
		used:     0,
		data:     make(map[int]int),
		nodeMap:  make(map[int]*doubleLinkedNode),
		doubleLinkedList: doubleLinkedList{
			head: nil,
			tail: nil,
		},
	}

	return lru
}

func (this *LRUCache) Get(key int) int {
	v, ok := this.data[key]
	if !ok {
		return -1
	}

	oldNode := this.nodeMap[key]
	if oldNode == this.head {
		return v
	}

	newNode := new(doubleLinkedNode)
	newNode.key = key
	newNode.next = nil
	newNode.pre = nil
	if oldNode == this.tail && this.tail.pre != nil {
		this.tail.pre.next = nil
		this.tail = this.tail.pre
		newNode.next = this.head
		this.head.pre = newNode
		this.head = newNode
	} else {
		oldNode.pre.next = oldNode.next
		oldNode.next.pre = oldNode.pre
		newNode.next = this.head
		this.head.pre = newNode
		this.head = newNode
	}

	this.nodeMap[key] = newNode
	return v
}

func (this *LRUCache) Put(key int, value int) {
	this.data[key] = value

	newNode := new(doubleLinkedNode)
	newNode.key = key
	if this.head == nil && this.tail == nil {
		this.used = 1
		this.doubleLinkedList.insertBeforHead(newNode)
		this.nodeMap[key] = newNode
	} else {
		_, ok := this.nodeMap[key]
		if ok {
			this.Update(key, value)
		} else {
			if this.used < this.capacity {
				this.used++
				this.doubleLinkedList.insertBeforHead(newNode)
				this.nodeMap[key] = newNode
			} else {
				tailKey := this.tail.key
				delete(this.data, tailKey)
				this.doubleLinkedList.deleteTail()
				delete(this.nodeMap, tailKey)
				this.insertBeforHead(newNode)

				this.nodeMap[key] = newNode
			}
		}
	}
}

func (this *LRUCache) Update(key, value int) {
	oldNode := this.nodeMap[key]
	this.data[key] = value
	if oldNode == this.head {
		return
	}

	this.doubleLinkedList.delete(oldNode)

	newNode := new(doubleLinkedNode)
	newNode.key = key
	newNode.next = nil
	newNode.pre = nil
	this.doubleLinkedList.insertBeforHead(newNode)

	this.nodeMap[key] = newNode
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
