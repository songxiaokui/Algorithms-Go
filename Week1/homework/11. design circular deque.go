package homework

import "sync"

/*
@Time    : 2021/3/11 11:12
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 11. design circular deque.go
@Software: GoLand
*/

/*
641. 设计循环双端队列
设计实现双端队列。
你的实现需要支持以下操作：

MyCircularDeque(k)：构造函数,双端队列的大小为k。
insertFront()：将一个元素添加到双端队列头部。 如果操作成功返回 true。
insertLast()：将一个元素添加到双端队列尾部。如果操作成功返回 true。
deleteFront()：从双端队列头部删除一个元素。 如果操作成功返回 true。
deleteLast()：从双端队列尾部删除一个元素。如果操作成功返回 true。
getFront()：从双端队列头部获得一个元素。如果双端队列为空，返回 -1。
getRear()：获得双端队列的最后一个元素。 如果双端队列为空，返回 -1。
isEmpty()：检查双端队列是否为空。
isFull()：检查双端队列是否满了。
示例：

MyCircularDeque circularDeque = new MycircularDeque(3); // 设置容量大小为3
circularDeque.insertLast(1);			        // 返回 true
circularDeque.insertLast(2);			        // 返回 true
circularDeque.insertFront(3);			        // 返回 true
circularDeque.insertFront(4);			        // 已经满了，返回 false
circularDeque.getRear();  				// 返回 2
circularDeque.isFull();				        // 返回 true
circularDeque.deleteLast();			        // 返回 true
circularDeque.insertFront(4);			        // 返回 true
circularDeque.getFront();				// 返回 4

提示：

所有值的范围为 [1, 1000]
操作次数的范围为 [1, 1000]
请不要使用内置的双端队列库。
*/

type MyCircularDeque struct {
	mux    sync.RWMutex
	queue  []int
	length int
	c      int
}

/** Initialize your data structure here. Set the size of the deque to be k. */
func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{queue: make([]int, 0, k), c: k}
}

/** Adds an item at the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertFront(value int) bool {
	this.mux.Lock()
	defer this.mux.Unlock()
	if this.IsFull() {
		return false
	}
	this.queue = append([]int{value}, this.queue[:this.length]...)
	this.length++
	return true
}

/** Adds an item at the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) InsertLast(value int) bool {
	this.mux.Lock()
	defer this.mux.Unlock()
	if this.IsFull() {
		return false
	}
	this.queue = append(this.queue[:this.length], value)
	this.length++
	return true
}

/** Deletes an item from the front of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteFront() bool {
	this.mux.Lock()
	defer this.mux.Unlock()
	if this.IsEmpty() {
		return false
	}

	this.queue = this.queue[1:this.length]
	this.length--
	return true
}

/** Deletes an item from the rear of Deque. Return true if the operation is successful. */
func (this *MyCircularDeque) DeleteLast() bool {
	this.mux.Lock()
	defer this.mux.Unlock()
	if this.IsEmpty() {
		return false
	}

	this.queue = this.queue[:this.length-1]
	this.length--
	return true
}

/** Get the front item from the deque. */
func (this *MyCircularDeque) GetFront() int {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if this.IsEmpty() {
		return -1
	}
	return this.queue[0]
}

/** Get the last item from the deque. */
func (this *MyCircularDeque) GetRear() int {
	this.mux.RLock()
	defer this.mux.RUnlock()
	if this.IsEmpty() {
		return -1
	}
	return this.queue[this.length-1]
}

/** Checks whether the circular deque is empty or not. */
func (this *MyCircularDeque) IsEmpty() bool {
	if this.length == 0 {
		return true
	}
	return false
}

/** Checks whether the circular deque is full or not. */
func (this *MyCircularDeque) IsFull() bool {
	if this.length == this.c {
		return true
	}
	return false
}
