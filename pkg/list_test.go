package pkg

import (
	"fmt"
	"reflect"
	"testing"
)

/*
@Time    : 2021/3/8 19:12
@Author  : austsxk
@Email   : austsxk@163.com
@File    : list_test.go.go
@Software: GoLand
*/

func TestArrayConvertLinkListAndLinkListConvertArray(t *testing.T) {
	tests := []struct {
		name string
		args []int
		want []int
	}{
		{"case1", []int{1, 2, 3, 4, 5}, []int{1, 2, 3, 4, 5}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := ArrayConvertLinkList(tt.args)
			fmt.Println(*got)
			data := LinkListConvertArray(got)
			fmt.Println(data)
			if !reflect.DeepEqual(data, tt.want) {
				t.Errorf("ArrayConvertLinkList() = %v, want %v", got, tt.want)
			}
		})
	}
}
