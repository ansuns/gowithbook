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

func main() {
	//创建Rect类型对象实例(1-4种方法都可以)
	rect1 := new(Rect)
	rect2 := &Rect{}
	rect3 := &Rect{0, 0, 100, 200}
	rect4 := &Rect{width: 300, height: 600}

	fmt.Println(rect4.Area())
	_, _, _, _ = rect1, rect2, rect3, rect4

}
