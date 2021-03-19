package main

import (
	"context"
	"fmt"
)

/*
@Time    : 2021/3/17 22:27
@Author  : austsxk
@Email   : austsxk@163.com
@File    : 素数.go
@Software: GoLand
*/

// 生成自然数
func GenerateNatural1() chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			ch <- i
		}
	}()
	return ch
}

// prime generate, return new channel
func FilterPrime(in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if x := <-in; x%prime != 0 {
				out <- x
			}
		}
	}()
	return out
}

func G() {
	rawChan := GenerateNatural1()
	for i := 0; i < 100; i++ {
		p := <-rawChan
		fmt.Println("生成的素数:", p)
		rawChan = FilterPrime(rawChan, p)
	}
}

/////////////////////////////////////////////////////////////////////////
// 上面会存在内存泄漏，所以使用context平滑关闭主函数

func CGenerateNatural(ctx context.Context) chan int {
	ch := make(chan int)
	go func() {
		for i := 2; ; i++ {
			select {
			default:
				ch <- i
			case <-ctx.Done():
				// 平滑关闭
				return
			}
		}
	}()
	return ch
}

func CFilterPrime(ctx context.Context, in chan int, prime int) chan int {
	out := make(chan int)
	go func() {
		for {
			if x := <-in; x%prime != 0 {
				select {
				default:
					out <- x
				case <-ctx.Done():
					return
				}
			}
		}
	}()
	return out
}

func CG() {
	ctx, cancel := context.WithCancel(context.Background())
	ch := CGenerateNatural(ctx)
	for i := 0; i < 100; i++ {
		p := <-ch
		fmt.Println("new 素数:", p)
		ch = CFilterPrime(ctx, ch, p)
	}
	// 主函数执行结束后，主动结束掉
	cancel()
}
func main() {
	// G()
	CG()
}
