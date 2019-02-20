package main

import "fmt"

func main() {
	//匿名函数
	f := func(x, y int) int {
		return x + y
	}
	fmt.Println(f(8, 9))
	fmt.Println("------------------------------------------------------------------------")
	//实例化一个接口IFly
	var fly IFly = new(Bird)
	fly.Fly()
	fmt.Println("------------------------------------------------------------------------")
	//并发编程
	values := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	resultChan := make(chan int, 2)
	go sum(values[:len(values)/2], resultChan)
	go sum(values[len(values)/2:], resultChan)
	sum1, sum2 := <-resultChan, <-resultChan //接收结果
	fmt.Println("Result:", sum1, sum2, sum1+sum2)
}

func sum(values [] int, resultChan chan int) {
	sum := 0
	for _, value := range values {
		sum += value
	}

	resultChan <- sum
}

//定义一个接口
type IFly interface {
	Fly()
}

//定义一个Bird类
type Bird struct {
}

//给这个类增加一个FLy()方法，实现了Fly方法就实现了IFly接口
func (b *Bird) Fly() {
	fmt.Println("Bird fly")
}
