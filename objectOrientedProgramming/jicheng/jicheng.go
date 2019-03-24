package main

import "fmt"

//基类
type Base struct {
}

//基类的一个成员方法
func (b *Base) Add(x, y int) int {
	return x + y
}

//基类的一个成员方法
func (b *Base) Less() {

}

//定义一个类，从Base继承
type My struct {
	Base //匿名组合
}

//
func (m *My) Add(x, y int) int {
	return m.Base.Add(x, y)
}

func main() {
	m := &My{}
	res := m.Add(7, 6)
	fmt.Println(res)
}
