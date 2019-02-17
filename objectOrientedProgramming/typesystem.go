package main

import "fmt"

type Interger int //定义一个新的类型Interger

/**
增加一个Less方法，是Interger的内部方法
 */
func (a Interger) Less(b Interger) bool {
	return a < b
}

func main() {
	var a Interger = 1
	if a.Less(2) {
		fmt.Println(a, "less 2")
	}
}
