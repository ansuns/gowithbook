package main

import "fmt"

type Interger int //定义一个新的类型Interger(使用 type 关键字可以定义出新的自定义类型)

/**
给Interger类型添加 Less 方法，int原来的方法被Integer继承
增加一个Less方法
 */
func (a Interger) Less(b Interger) bool {
	return a < b
}

func (a *Interger) Add(b Interger) {
	*a += b
}

func main() {
	var a Interger = 1 //a 是一个int对象
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}

	a.Add(2)
	fmt.Println(a)
}
