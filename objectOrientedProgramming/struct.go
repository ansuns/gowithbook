package main

import "fmt"

/**
类似其他语言的class 类
 */

//定义一个sturcy
type Rect struct {
	x, y          float64
	width, height float64
}

//添加一个计算面积的方法
func (r *Rect) Area() float64 {
	return r.width * r.height
}

//提供继承，但是已组合的文法（匿名组合）
//定义一个Base类，实现Foo(),Bar()两个成员方法
type Base struct {
	Name string
}

//为Base类增加Foo方法
func (base *Base) Foo() {

}

//为Base累增加Bar方法
func (base *Base) Bar() {

}

//定义Foo类，从Base继承
type Foo struct {
	Base
	//...其他成员
}

//改写Bar方法
func (foo *Foo) Bar() {
	foo.Base.Bar()
}

func main() {
	//创建Rect类型对象实例(1-4种方法都可以)
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 300, height: 600}

	fmt.Println(rect4.Area())
	_, _, _, _ = rect1, rect2, rect3, rect4

}
