package main

import "fmt"

/**
流程控制
 */

func main() {
	val := first(7)
	fmt.Println(val)

	i := 1
	switch i {
	case 0:
		fmt.Println(0)
	case 1:
		fmt.Println(1)
	case 2:
		fmt.Println(2)
	default:
		fmt.Println("Default")

	}

	//switch后面的表达式可能也不是必须的：
	switch {
	case 0 < val:
		fmt.Println("aa")
	default:
		fmt.Println("not aa")
	}

	testFor()
}

func first(a int) int {
	if a == 0 {
		return 4
	} else {
		return a * a
	}
}

func testFor() {
	sum := 0
	for {
		sum++
		if sum > 10 {
			break
		}
		fmt.Println(sum)
	}
}
