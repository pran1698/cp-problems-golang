package main

type Node struct {
	Key  int
	Val  int
	Next *Node
	Back *Node
}

type LRUCache struct {
	cap  int
	size int
	head *Node
	tail *Node
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		cap:  capacity,
		size: 0,
		head: nil,
		tail: nil,
	}
}

func (this *LRUCache) Get(key int) int {
	if this.tail != nil && this.tail.Key == key {
		return this.tail.Val
	}

	t := this.head
	for t != nil {
		if t.Key == key {
			//fmt.Println(t)
			this.UpdateTail(t)
			return t.Val
		}
		t = t.Next
	}

	return -1
}

func (this *LRUCache) Put(key int, value int) {
	v := this.Get(key)
	if v != -1 {
		this.tail.Val = value
		return
	}

	if this.size == this.cap {
		if this.cap != 1 {
			//fmt.Println("head",this.head)
			//fmt.Println("head.next", this.head.Next)
			this.head = this.head.Next
			this.head.Back = nil
		}
		this.size--
	}

	n := &Node{
		Key:  key,
		Val:  value,
		Next: nil,
		Back: nil,
	}

	if this.cap == 1 {
		this.head = n
		this.tail = n
	} else {
		if this.head == nil {
			this.head = n
			this.tail = n
		} else {
			this.tail.Next = n
			n.Back = this.tail
			this.tail = n
		}
	}
	this.size++
	//fmt.Println(this.head, this.tail)
}

func (this *LRUCache) UpdateTail(t *Node) {
	if this.head == t {
		this.head = t.Next
	}

	for t.Next != nil {
		temp1 := t.Next
		temp2 := t.Back
		//fmt.Println(temp1,temp2)

		t.Next = temp1.Next
		t.Back = temp1

		temp1.Back = temp2
		if temp2 != nil {
			temp2.Next = temp1
		}

		temp1.Next = t
	}
	this.tail = t
	//fmt.Println(this)
	//fmt.Println("head", "tail",this.head, this.tail)
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
