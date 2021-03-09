package pkg

/*
@Time    : 2021/3/9 22:11
@Author  : austsxk
@Email   : austsxk@163.com
@File    : heap.go
@Software: GoLand
*/

// int type heap, impl Less Pop Push methods
type IntHeap []int

// get IntHeap length, numbers of element of list
func (h IntHeap) Len() int {
	return len(h)
}

// judge list[i] < list[k]
func (h IntHeap) Less(i, k int) bool {
	return h[i] < h[k]
}

// swap i, k value
func (h IntHeap) Swap(i, k int) {
	h[i], h[k] = h[k], h[i]
}

// up is impl sort.interface interface
// down is impl Push Pop methods

// add a element to IntHeap
func (h *IntHeap) Push(e interface{}) {
	*h = append(*h, e.(int))
}

// pop min element from heap
func (h *IntHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[:n-1]
	return x
}

// combination impl heap.Interface interface

// that can use heap init all methods

/*
push methods --> append last --> up

func up(h Interface, j int) {
	for {
		i := (j - 1) / 2 // parent
		if i == j || !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		j = i
	}
}

delete a element: last swap first --> down --> return
func Pop(h Interface) interface{} {
	n := h.Len() - 1
	h.Swap(0, n)
	down(h, 0, n)
	return h.Pop()
}

heap use a list to save, index can find a element child element if use complete binary tree

func down(h Interface, i0, n int) bool {
	i := i0
	for {
		j1 := 2*i + 1
		if j1 >= n || j1 < 0 { // j1 < 0 after int overflow
			break
		}
		j := j1 // left child
		if j2 := j1 + 1; j2 < n && h.Less(j2, j1) {
			j = j2 // = 2*i + 2  // right child
		}
		if !h.Less(j, i) {
			break
		}
		h.Swap(i, j)
		i = j
	}
	return i > i0
}
*/
