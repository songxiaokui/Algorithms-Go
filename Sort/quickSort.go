package Sort

/*
@Time    : 2021/3/12 09:44
@Author  : austsxk
@Email   : austsxk@163.com
@File    : quickSort.go
@Software: GoLand
*/

// 快速排序
// 关键就是partition 以一个分割值，将数组分成两部分

func QuickSort(array []int, left, right int) {
	if left < right {
		i, j := left, right
		mid := array[(left+right)/2]
		for array[i] < mid {
			i++
		}
		for array[j] > mid {
			j--
		}
		if i <= j {
			array[i], array[j] = array[j], array[i]
			i++
			j--
		}
		if left < j {
			QuickSort(array, left, j)
		}
		if right > i {
			QuickSort(array, i, right)
		}
	}
}
