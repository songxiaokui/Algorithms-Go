package main

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func worker(ch chan bool, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		case <-ch:
			return
		default:
			time.Sleep(time.Second)
			fmt.Println("doing something")
		}
	}
}

func workers2(i int, ctx context.Context, wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		select {
		default:
			time.Sleep(time.Second * 1)
			fmt.Println(i, "do something")
		case <-ctx.Done():
			fmt.Println(ctx.Err())
			return
		}
	}
}

func Run() {
	var wg sync.WaitGroup
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*3)
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go workers2(i, ctx, &wg)
	}

	time.Sleep(time.Second * 1)
	cancel()
	wg.Wait()
	fmt.Println("shot down")
}
func main() {
	//var wg sync.WaitGroup
	//ch := make(chan bool)
	//for i:=0; i<5; i++ {
	//	wg.Add(1)
	//	go worker(ch, &wg)
	//}
	//
	//time.Sleep(time.Second * 5)
	//close(ch)
	//wg.Wait()
	//fmt.Println("shutdown")

	Run()
}
