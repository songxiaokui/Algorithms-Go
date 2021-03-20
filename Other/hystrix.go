package main

/*
@Time    : 2021/3/20 11:28
@Author  : austsxk
@Email   : austsxk@163.com
@File    : hystrix.go
@Software: GoLand
*/

// 熔断器
import (
	"errors"
	"fmt"
	"github.com/afex/hystrix-go/hystrix"
	"math/rand"
	"sync"
	"time"
)

type User struct {
	Name string
	Age  int
}

func GetUser() User {
	d := rand.Intn(10)
	if d < 7 {
		time.Sleep(time.Second * 3)
	}
	return User{"宋晓奎", 27}
}

func DefaultUser() User {
	return User{"rabbit", -1}
}

// 超时降级
func TimeOut() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout: 1000,
	})
	rand.Seed(time.Now().UnixNano())
	userChannel := make(chan User, 1)
	for {
		err := hystrix.Go("my_command", func() error {
			u := GetUser()
			userChannel <- u
			return nil
		}, func(err error) error {
			// 如果上面的logic出现错误，则可以在此处捕获
			// 反回兜底数据
			//u := DefaultUser()
			//userChannel <- u
			//return nil
			return errors.New("time out")

		})
		select {
		case u := <-userChannel:
			fmt.Println(u)
		case e := <-err:
			fmt.Println(e)
		}
		time.Sleep(time.Second * 1)
	}

}

// 并发降级
func Currency() {
	hystrix.ConfigureCommand("my_command", hystrix.CommandConfig{
		Timeout:               2000,
		MaxConcurrentRequests: 5,
	})
	rand.Seed(time.Now().UnixNano())
	userChannel := make(chan User, 1)
	wg := sync.WaitGroup{}

	for i := 0; i < 10; i++ {

		go (func() {
			wg.Add(1)
			defer wg.Done()
			err := hystrix.Go("my_command", func() error {
				u := GetUser()
				userChannel <- u
				fmt.Println(u)
				return nil
			}, func(err error) error {
				fmt.Println(err)
				u := DefaultUser()
				userChannel <- u
				return err

			})
			select {
			case u := <-userChannel:
				fmt.Println(u)
			case x := <-err:
				fmt.Println(x)
			}

		})()
	}
	wg.Wait()
}

// 错误熔断
func Hystrix() {
	hystrix.ConfigureCommand("hystrix", hystrix.CommandConfig{
		Timeout:                2000,
		RequestVolumeThreshold: 3,
		ErrorPercentThreshold:  10,
		// 5s后半开状态
		SleepWindow: 5000,
	})
	// 获取熔断器
	c, _, _ := hystrix.GetCircuit("hystrix")
	for {
		hystrix.Do("hystrix", func() error {
			u := GetUser()
			fmt.Println(u)
			return nil
		}, func(err error) error {
			fmt.Println(err)
			return err
		})
		fmt.Println("当前熔断器状态:", c.IsOpen())
		time.Sleep(time.Second)
	}
}

func main() {
	Hystrix()
}
