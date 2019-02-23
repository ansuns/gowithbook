package main

import "fmt"

func main() {
	fmt.Println("Hello World!")
	_, _, nickname := getName()
	fmt.Println(nickname)
	info, err := orderInfo(158)
	fmt.Println(info, err)
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

/**
没有明确赋值时，返回默认值，并不是所遇的返回值都必须赋值
 */
func orderInfo(order_id int) (info []int, err error) {
	return
}
