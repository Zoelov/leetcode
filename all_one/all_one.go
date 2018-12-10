package all_one

/**************************************************************

	https://leetcode-cn.com/problems/all-oone-data-structure/

	实现一个数据结构支持以下操作：

	Inc(key) - 插入一个新的值为 1 的 key。或者使一个存在的 key 增加一，保证 key 不为空字符串。
	Dec(key) - 如果这个 key 的值是 1，那么把他从数据结构中移除掉。否者使一个存在的 key 值减一。如果这个 key 不存在，这个函数不做任何事情。key 保证不为空字符串。
	GetMaxKey() - 返回 key 中值最大的任意一个。如果没有元素存在，返回一个空字符串""。
	GetMinKey() - 返回 key 中值最小的任意一个。如果没有元素存在，返回一个空字符串""。
	挑战：以 O(1) 的时间复杂度实现所有操作。

***************************************************************/

import "fmt"

/***************************************
 解题思路：
 O(1)的数据结构肯定是用map,但为了O(1)的找到最大和
 最小的值，可以使用双向链表

 下面结构体中，data存储了所有的key和对应的值
 doublelinkedList是个双向链表，它的值为key对应的
 值和相同值的key的一个map,并且保持这个双向链表是从
 小到大排序的。
 posMap 存储了每一个值在双向链表中的位置。
*****************************************/

type AllOne struct {
	data   map[string]int
	posMap map[int]*doubleLinkedNode
	doublelinkedList
}

type doubleLinkedNode struct {
	val  int
	keys map[string]string
	pre  *doubleLinkedNode
	next *doubleLinkedNode
}

type doublelinkedList struct {
	head *doubleLinkedNode
	tail *doubleLinkedNode
	size int
}

/** Initialize your data structure here. */
func Constructor() AllOne {
	a := AllOne{
		data:   make(map[string]int),
		posMap: make(map[int]*doubleLinkedNode),
		doublelinkedList: doublelinkedList{
			head: nil,
			tail: nil,
			size: 0,
		},
	}

	return a
}

/** Inserts a new key <Key> with value 1. Or increments an existing key by 1. */
func (this *AllOne) Inc(key string) {
	if this.head == nil && this.tail == nil {
		newNode := new(doubleLinkedNode)
		newNode.val = 1
		newNode.next = nil
		newNode.pre = nil
		newNode.keys = make(map[string]string)
		newNode.keys[key] = key
		this.head = newNode
		this.tail = newNode
	}

	oldValue, ok := this.data[key]
	if ok {
		this.data[key] = oldValue + 1
		oldPosNode := this.posMap[oldValue]

		_, has := this.posMap[oldValue+1]
		if has {
			(this.posMap[oldValue+1].keys)[key] = key
		} else {
			// 插入到oldPosNode后面
			newNode := new(doubleLinkedNode)
			newNode.val = oldValue + 1
			newNode.keys = make(map[string]string)
			newNode.keys[key] = key

			if oldPosNode == this.head && oldPosNode == this.tail {
				newNode.next = nil
				newNode.pre = this.head
				this.head.next = newNode
				this.tail = newNode
			} else if oldPosNode == this.head {
				newNode.next = this.head.next
				newNode.pre = this.head
				this.head.next.pre = newNode
				this.head.next = newNode
			} else if oldPosNode == this.tail {
				newNode.next = nil
				newNode.pre = this.tail
				this.tail.next = newNode
				this.tail = newNode
			} else {
				newNode.next = oldPosNode.next
				newNode.pre = oldPosNode.next.pre
				oldPosNode.next.pre = newNode
				oldPosNode.next = newNode
			}

			this.posMap[oldValue+1] = newNode
		}

		// 处理oldNode
		delete(oldPosNode.keys, key)
		if len(oldPosNode.keys) == 0 {
			delete(this.posMap, oldValue)
			// 删除old node
			if oldPosNode == this.head {
				this.head.next.pre = nil
				this.head = this.head.next
			} else if oldPosNode == this.tail {
				this.tail.pre.next = nil
				this.tail = this.tail.pre
			} else {
				oldPosNode.pre.next = oldPosNode.next
				oldPosNode.next.pre = oldPosNode.pre
			}
		}

	} else {
		this.data[key] = 1
		if this.head.val == 1 {
			this.head.keys[key] = key
			this.posMap[1] = this.head
		} else {
			newNode := new(doubleLinkedNode)
			newNode.val = 1
			newNode.keys = make(map[string]string)
			newNode.keys[key] = key
			newNode.next = this.head
			newNode.pre = nil
			this.head.pre = newNode
			this.head = newNode

			this.posMap[1] = newNode
		}
	}

}

/** Decrements an existing key by 1. If Key's value is 1, remove it from the data structure. */
func (this *AllOne) Dec(key string) {
	oldValue, ok := this.data[key]
	if !ok {
		return
	}

	if oldValue == 1 {
		delete(this.data, key)
		oldPosNode := this.posMap[oldValue]

		delete(oldPosNode.keys, key)
		if len(oldPosNode.keys) == 0 {
			if oldPosNode == this.head {
				oldPosNode.pre = nil
				this.head = oldPosNode.next
			} else if oldPosNode == this.tail {
				oldPosNode.next = nil
				this.tail = oldPosNode.pre
			} else {
				oldPosNode.pre.next = oldPosNode.next
				oldPosNode.next.pre = oldPosNode.pre
			}

			delete(this.posMap, 1)
		}
	} else {
		this.data[key] = oldValue - 1
		oldPosNode := this.posMap[oldValue]

		newPosNode, ok := this.posMap[oldValue-1]

		if ok {
			newPosNode.keys[key] = key
		} else {
			newNode := new(doubleLinkedNode)
			newNode.val = oldValue - 1
			newNode.keys = make(map[string]string)
			newNode.keys[key] = key

			if oldPosNode == this.head {
				newNode.next = this.head
				newNode.pre = nil
				this.head.pre = newNode
				this.head = newNode
			} else if oldPosNode == this.tail {
				newNode.next = this.tail
				newNode.pre = this.tail.pre
				this.tail.pre = newNode
			} else {
				newNode.pre = oldPosNode.pre
				newNode.next = oldPosNode
				oldPosNode.pre.next = newNode
			}

			this.posMap[oldValue-1] = newNode
		}

		delete(oldPosNode.keys, key)
		if len(oldPosNode.keys) == 0 {
			delete(this.posMap, oldValue)
			// 删除old node
			if oldPosNode == this.head {
				this.head.next.pre = nil
				this.head = this.head.next
			} else if oldPosNode == this.tail {
				this.tail.pre.next = nil
				this.tail = this.tail.pre
			} else {
				oldPosNode.pre.next = oldPosNode.next
				oldPosNode.next.pre = oldPosNode.pre
			}
		}
	}
}

/** Returns one of the keys with maximal value. */
func (this *AllOne) GetMaxKey() string {
	if this.tail == nil {
		return ""
	}

	for k := range this.tail.keys {
		return k
	}

	return ""
}

/** Returns one of the keys with Minimal value. */
func (this *AllOne) GetMinKey() string {
	if this.head == nil {
		return ""
	}

	for k := range this.head.keys {
		return k
	}

	return ""
}

func (this *AllOne) AllMax() {
	fmt.Printf("max: ")
	for k := range this.tail.keys {
		fmt.Printf("%v:%v ", k, this.data[k])
	}
	fmt.Println()
}

func (this *AllOne) ALlMin() {
	fmt.Printf("min: ")
	for k := range this.head.keys {
		fmt.Printf("%v:%v ", k, this.data[k])
	}
	fmt.Println()
}

/**
 * Your AllOne object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Inc(key);
 * obj.Dec(key);
 * param_3 := obj.GetMaxKey();
 * param_4 := obj.GetMinKey();
 */
