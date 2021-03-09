package pkg

import "container/heap"

/*
@Time    : 2021/3/9 22:33
@Author  : austsxk
@Email   : austsxk@163.com
@File    : priorityQueue.go
@Software: GoLand
*/

type AustItme struct {
	value    string // item value
	priority int    // item priority of heap
	index    int    // item in queue location
}

type PriorityQueue []*AustItme

// impl sort.Interface interface
func (p PriorityQueue) Len() int {
	return len(p)
}

func (p PriorityQueue) Less(i, k int) bool {
	return p[i].priority > p[k].priority
}

func (p PriorityQueue) Swap(i, k int) {
	p[i], p[k] = p[k], p[i]
	p[i].index = k
	p[k].index = i
}

// impl other methods
func (p *PriorityQueue) Push(e interface{}) {
	old := *p
	n := len(old)
	item := e.(*AustItme)
	// record pushed e current index
	item.index = n
	*p = append(*p, item)
}

// pop last element
func (p *PriorityQueue) Pop() interface{} {
	old := *p
	n := len(old)
	item := old[n-1]

	// avoid memory leak
	old[n-1] = nil
	// data safe
	item.index = -1

	*p = old[:n-1]
	return item
}

func (p *PriorityQueue) update(item *AustItme, value string, d int) {
	item.value = value
	item.priority = d
	// TODO Fix how to work
	heap.Fix(p, item.index)
}
