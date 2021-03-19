package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"net/http"
)

/*
@Time    : 2021/3/18 21:20
@Author  : austsxk
@Email   : austsxk@163.com
@File    : errgroup.go
@Software: GoLand
*/
func UserErrGroup() {
	var g = new(errgroup.Group)
	urls := []string{
		"http://www.austsxk.com",
		"https://www.baidu.com/",
		"https://www.google.com",
	}
	for _, url := range urls {
		u := url
		g.Go(func() error {
			resp, err := http.Get(u)
			if err == nil {

				resp.Body.Close()
			}
			return err
		})
	}
	if err := g.Wait(); err == nil {
		fmt.Println("success")
	} else {
		fmt.Println(err)
	}
}
func main() {
	UserErrGroup()
}
