package pkg

/*
@Time    : 2021/3/9 22:24
@Author  : austsxk
@Email   : austsxk@163.com
@File    : heap_test.go
@Software: GoLand
*/

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestAllHeapOperation(t *testing.T) {
	data := &IntHeap{2, 1, 7, 4, 8}
	heap.Init(data)
	heap.Push(data, 9)
	heap.Push(data, 3)
	heap.Push(data, 22)
	fmt.Println("min heap value:", (*data)[0])
	for data.Len() > 0 {
		fmt.Println("  弹出元素: ", heap.Pop(data))
	}
}
