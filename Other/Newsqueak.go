package main

import (
	"fmt"
)

/*
@Time    : 2021/3/17 00:27
@Author  : austsxk
@Email   : austsxk@163.com
@File    : Newsqueak.go
@Software: GoLand
*/

func GenerateNatural() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

func PrimeFilter(x int, in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if i := <-in; i%prime != 0 {
				out <- i
			}
		}
	}()
	return out
}

func main() {
	//ch := GenerateNatural()
	//for i := 0; i < 1000; i++ {
	//	prime := <-ch //只做存储操作
	//	fmt.Println("generate natural number is :", prime)
	//	// ch 等价一个过滤器，每一次过滤上一个素数的值，线过滤2的倍数，将其赋值到新的chan中，然后过率新的素数
	//	ch = PrimeFilter(i, ch, prime)
	//}
H:
	{
		fmt.Println("hello label")
	}

	fmt.Println("kkk")

	goto H
}
