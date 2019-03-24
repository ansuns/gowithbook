package main

import "fmt"

func main() {
	//定义一个从Animal继承的Dog对象
	var ddog Animal = new(Dog)
	ddog.Run()
}

//定义一个接口
type Animal interface {
	//定义一个抽象方法
	Run()
}

//定义一个类
type Dog struct {

}

//实现Animal接口,只要实现了Animal的方法，就实现了该接口
func (d *Dog) Run(){
	fmt.Println("Dog running")
}

