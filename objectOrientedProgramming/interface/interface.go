package main

import "fmt"

//定义一个接口
type Human interface {
	Say() //只有声明，不实现（用其他类实现）
}

//定义一个类
type Student struct {
}

//实现Human方法
func (s *Student) Say() {
	fmt.Println("i‘m astudent")
}

//定义一个类
type Teacher struct {
}

//实现Human方法
func (t *Teacher) Say() {
	fmt.Println("i‘m teacher")
}
func main() {
	//已经有两个类型Student,Teacher实现了Human接口，一个接口的不同提现
	var h Human
	s := &Student{}
	h = s
	h.Say()

	t := &Teacher{}
	h = t
	t.Say()

}
