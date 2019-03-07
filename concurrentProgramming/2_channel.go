package main

//通过channel 实现同步
import (
	"fmt"
	"time"
)

//全局变量channel
var ch = make(chan int)

func main() {
	//这里不分先后，因为要同步：
	/**
	go person2()
	go person1()
	结果一样
	 */
	go person2()
	go person1()

	//主协程10s后结束
	time.Sleep(time.Second * 10)
}

func Printer(str string) {
	for _, data := range str {
		fmt.Printf("%c", data)
		time.Sleep(time.Second)
	}
	fmt.Printf("\n")
}

//person1执行完才执行person2
func person1() {
	Printer("Hello")
	ch <- 1 //给管道写数据
}

func person2() {
	<-ch //从管道取数据,接收，如果管道没有数据，则堵塞
	Printer("World")
}
