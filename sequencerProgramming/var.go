package main

import "fmt"

//纯粹变量声明用var
// 以下是声明变量
var v1 int
var (
	v2 string
	v3 int
)
var v4 [10] int       //数组
var v6 *int           //指针
var v7 map[string]int //map,key 为sring类型，value 为int类型
var v8 func(a int) int

const C1 = 0 // 常量const定义
//预定义常量
const ( //iota被重设为0
	cc0 = iota //0
	cc1 = iota //1
	cc2 = iota //2
)

//枚举
const (
	Sunday = iota
	Monday
	Tuesday
	Wednesday
	Thursday
	Friday
	Sturday
	numerOfDays
)

var b1 bool //布尔值声明

func main() {

	v1 = 1088
	v9 := 10
	fmt.Println(v1)
	fmt.Println(v9)
	v1, v9 = v9, v1 //灵活的多重赋值
	fmt.Println(v1)
	fmt.Println(v9)

	_, secondeName, _ := getName() // _匿名变量
	fmt.Println(secondeName)

	b1 = true
	b2 := (1 == 1)
	if b1 == b2 {
		fmt.Println("b1,b2都为真")
	}

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
