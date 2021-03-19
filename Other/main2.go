package main

import (
	"golang.org/x/time/rate"
	"log"
	"net/http"
	"time"
)

/*
@Time    : 2021/3/15 16:52
@Author  : austsxk
@Email   : austsxk@163.com
@File    : main2.go
@Software: GoLand
*/

var Limit = rate.NewLimiter(1, 5)

func RateLimiter(limit *rate.Limiter, next http.HandlerFunc) http.HandlerFunc {
	return func(writer http.ResponseWriter, request *http.Request) {
		if !limit.AllowN(time.Now(), 5) {
			http.Error(writer, "request too many", http.StatusTooManyRequests)
			return
		}
		next(writer, request)
	}
}

func RootHandler(writer http.ResponseWriter, request *http.Request) {
	_, err := writer.Write([]byte("hello golang"))

	if err != nil {
		log.Fatal("system error")
	}
}

func main() {
	c := http.NewServeMux()
	c.HandleFunc("/", RateLimiter(Limit, RootHandler))
	if err := http.ListenAndServe(":9999", c); err != nil {
		log.Fatal(err)
	}
}
