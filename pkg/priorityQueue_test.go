package pkg

/*
@Time    : 2021/3/10 00:20
@Author  : austsxk
@Email   : austsxk@163.com
@File    : priorityQueue_test.go
@Software: GoLand
*/

import (
	"container/heap"
	"fmt"
	"testing"
)

func TestPriorityQueueAll(t *testing.T) {
	maps := map[string]int{
		"Math":    88,
		"Chinese": 100,
		"English": 77,
		"Other":   81,
	}
	pq := make(PriorityQueue, len(maps))

	i := 0
	for key, value := range maps {
		pq[i] = &AustItme{
			value:    key,
			priority: value,
			index:    i,
		}
		i++
	}
	// init my heap
	heap.Init(&pq)

	it := &AustItme{
		value:    "HH",
		priority: 9,
	}
	heap.Push(&pq, it)

	// update item priority
	pq.update(it, it.value, 999)

	for pq.Len() > 0 {
		d := heap.Pop(&pq).(*AustItme)
		fmt.Printf("current element value:%s, priority:%d \n", d.value, d.priority)
	}

}
