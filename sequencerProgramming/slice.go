package main

import "fmt"

var array [5]int //数组定义

func main() {
	//数组赋值
	array = [5]int{1, 2, 3, 4, 5}
	var slice1 []int = array[:5] //数组切片 0-5
	for _, v := range slice1 {
		fmt.Println("slice1 v: ", v)
	}
	fmt.Println("------------------------------------------------------------------------")
	//也可以使用Make 创建切片
	newSlice1 := make([]int, 5)     //初始个数为5，元素初始值为0的切片
	newSlice2 := make([]int, 5, 10) //初始个数为5，元素初始值为0的切片,并预留10个元素的存储空间，指10个，不是5+10=15
	newSlice3 := []int{1, 2, 3}     //直接创建并初始化包含3个元素的切片

	newSlice1, newSlice2, newSlice3 = newSlice1, newSlice2, newSlice3

	//元素遍历
	for _, v := range newSlice2 {
		fmt.Println(" ", v)
	}
	fmt.Println("------------------------------------------------------------------------")
	newSlice2 = append(newSlice2, 8, 9, 10)          //动态追加元素
	fmt.Println("len of newSlice2 ", len(newSlice2)) //返回元素个数
	fmt.Println("cap of newSlice2 ", cap(newSlice2)) //返回分配的空间大小
	fmt.Println("------------------------------------------------------------------------")
	for _, v := range newSlice2 {
		fmt.Println(" ", v)
	}
	fmt.Println("------------------------------------------------------------------------")
	newSlice1 = append(newSlice1, newSlice3...) //...相当于把newSlice3打散后传入：append(newSlice1,1,2,3)
	for _, v := range newSlice1 {
		fmt.Println(" ", v)
	}
	fmt.Println("------------------------------------------------------------------------")

	//基于切片创建切片
	sSlice1 := []int{100, 200, 300, 400, 500}
	sSlice2 := sSlice1[2:]
	for _, v := range sSlice2 {
		fmt.Println(" ", v)
	}
	fmt.Println("------------------------------------------------------------------------")

	//复制切片：从一个数组切片复制到另一个数组切片
	dSlice1 := []string{"a", "b", "c", "d", "e"}
	dSlice2 := []string{"z"}
	copy(dSlice2, dSlice1)
	for _, v := range dSlice2 {
		fmt.Println("copy value:", v) //结果a,复制1的一个元素到2的前一个位置
	}
}
