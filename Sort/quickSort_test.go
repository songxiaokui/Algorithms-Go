package Sort

import (
	"reflect"
	"sort"
	"testing"
)

/*
@Time    : 2021/3/12 09:50
@Author  : austsxk
@Email   : austsxk@163.com
@File    : quickSort_test.go.go
@Software: GoLand
*/

func TestQuickSort(t *testing.T) {
	type args struct {
		array []int
		left  int
		right int
	}
	tests := []struct {
		name string
		args args
	}{
		{"sort", struct {
			array []int
			left  int
			right int
		}{array: []int{1, 4, 2, 7, 9, 5, 8}, left: 0, right: 6}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			d := tt.args.array
			QuickSort(tt.args.array, tt.args.left, tt.args.right)
			sort.Slice(tt.args.array, func(i, j int) bool {
				return tt.args.array[i] < tt.args.array[j]
			})
			if !reflect.DeepEqual(d, tt.args.array) {
				t.Errorf("error")
			}
		})
	}
}
