package main

import "fmt"

func main() {
	ch := make(chan string)

	go func() {
		defer fmt.Println("子协程调用完毕")
		for i := 0; i < 2; i++ {
			fmt.Println("子协程 i = ", i)
		}

		ch <- "我是子协程，工作完毕"
	}()

	str := <-ch
	fmt.Println("str = ", str)
}
