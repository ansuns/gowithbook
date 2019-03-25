package main

import (
	"fmt"
	"time"
)

var ch1 chan int

func main() {

	timeout := make(chan bool, 1)
	//超时匿名等待函数
	go func() {
		time.Sleep(time.Second) //等待一秒钟
		timeout <- true
	}()

	select {
	case <-ch1:
	case <-timeout: //已超时chan读出到了数据
		close(timeout) //关闭chan
		fmt.Println("请求超时")
	}

	//判断一个chan是否已关闭
	_, ok := <-timeout
	if ok == false {
		fmt.Println("timeout chan 已关闭")
	}
}
