package main

type NodeMap struct {
	Key  int
	Val  int
	Next *NodeMap
	Prev *NodeMap
}

type LRUCacheMap struct {
	capacity int
	size     int
	cache    map[int]*NodeMap
	head     *NodeMap
	tail     *NodeMap
}

func ConstructorMap(capacity int) LRUCacheMap {
	return LRUCacheMap{
		capacity: capacity,
		cache:    make(map[int]*NodeMap),
		head:     nil,
		tail:     nil,
	}
}

func (this *LRUCacheMap) Get(key int) int {
	node, exists := this.cache[key]
	if !exists {
		return -1
	}
	this.moveToTail(node)
	return node.Val
}

func (this *LRUCacheMap) Put(key int, value int) {
	if node, exists := this.cache[key]; exists {
		node.Val = value
		this.moveToTail(node)
	} else {
		if this.size == this.capacity {
			delete(this.cache, this.head.Key)
			this.removeNode(this.head)
		} else {
			this.size++
		}

		node := &NodeMap{
			Key:  key,
			Val:  value,
			Next: nil,
			Prev: nil,
		}
		this.addNode(node)
		this.cache[key] = node
	}
}

func (this *LRUCacheMap) moveToTail(node *NodeMap) {
	if node == this.tail {
		return
	}
	this.removeNode(node)
	this.addNode(node)
}

func (this *LRUCacheMap) removeNode(node *NodeMap) {
	if node == this.head {
		this.head = node.Next
	}
	if node == this.tail {
		this.tail = node.Prev
	}
	if node.Prev != nil {
		node.Prev.Next = node.Next
	}
	if node.Next != nil {
		node.Next.Prev = node.Prev
	}
	node.Prev = nil
	node.Next = nil
}

func (this *LRUCacheMap) addNode(node *NodeMap) {
	if this.tail != nil {
		this.tail.Next = node
		node.Prev = this.tail
	}
	this.tail = node
	if this.head == nil {
		this.head = node
	}
}
