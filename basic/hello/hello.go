package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	_, _, nickname := getName()
	fmt.Println(nickname)
}

/*
获取名字
 */
func getName() (firstName, secondName, nickName string) {
	firstName = "张"
	secondName = "三"
	nickName = "二狗"
	return
}
