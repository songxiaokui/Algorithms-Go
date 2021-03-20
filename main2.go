package main

import (
	"fmt"
	"time"
)

/*
@Time    : 2021/3/19 19:56
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main2.go
@Software: GoLand
*/

func rangeChannel() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2
	time.AfterFunc(time.Second, func() {
		close(ch)
	})
	fmt.Println("xxxx")
	for x := range ch {
		fmt.Println(x)
	}
}

func main() {
	rangeChannel()
}
