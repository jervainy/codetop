package test

import "testing"

type Node struct {
	Key  int
	Val  int
	Next *Node
	Pre  *Node
}

type LRUCache struct {
	Head *Node
	Tail *Node
	Map  map[int]*Node
	Cap  int
	Len  int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		Head: nil,
		Tail: nil,
		Map:  make(map[int]*Node),
		Cap:  capacity,
		Len:  0,
	}
}

func (this *LRUCache) Del(node *Node) {
	delete(this.Map, node.Key)
	if this.Len <= 1 {
		this.Head = nil
		this.Tail = nil
		this.Len = 0
	} else {
		pre, next := node.Pre, node.Next
		if pre == nil {
			this.Head = next
		} else {
			pre.Next = next
		}
		if next == nil {
			this.Tail = pre
		} else {
			next.Pre = pre
		}
		this.Len -= 1
		node.Pre = nil
		node.Next = nil
	}
}

func (this *LRUCache) Append(node *Node) {
	if this.Len == 0 {
		this.Head = node
		this.Tail = node
	} else {
		this.Tail.Next = node
		node.Pre = this.Tail
		this.Tail = node
	}
	this.Map[node.Key] = node
	this.Len += 1
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.Map[key]; ok {
		if node != this.Tail {
			this.Del(node)
			this.Append(node)
		}
		return node.Val
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	node, ok := this.Map[key]
	if ok { // 命中 或 容量已满
		this.Del(node)
	} else if this.Len >= this.Cap {
		this.Del(this.Head)
	}
	this.Append(&Node{Key: key, Val: value})
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */

func TestDo146(t *testing.T) {
	obj := Constructor(2)
	obj.Put(1, 1)
	obj.Put(2, 2)
	para1 := obj.Get(1)
	obj.Put(3, 3)
	obj.Get(2)
	obj.Put(4, 4)
	obj.Get(1)
	obj.Get(3)
	obj.Get(4)
	println(para1)
}
