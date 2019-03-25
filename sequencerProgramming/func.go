package main

import (
	"errors"
	"fmt"
)

func main() {
	sum, err := Add(1, 2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}

	ret2 := getSum(1, 2, 3, 4)
	fmt.Println(ret2)
	fmt.Println("------------------------------------------------------------------------")
	varArgs(sum, err, ret2)
}

func Add(val1, val2 int) (sum int, err error) {

	if val2 < 0 || val1 < 0 {
		err = errors.New("val1或val2不能为负数")
		return
	}
	return val1 + val2, nil
}

/**
不定参数
*/
func getSum(args ...int) int {
	sum := 0
	for _, v := range args {
		sum += v
	}

	return sum
}

/**
任意类型不定参数
*/
func varArgs(args ...interface{}) {
	for _, arg := range args {
		switch arg.(type) {
		case int:
			fmt.Println(arg, "is a int value")
		case string:
			fmt.Println(arg, "is a string value")
		case int64:
			fmt.Println(arg, "is a int64 value")
		default:
			fmt.Println(arg, "is a unkonw type")
		}
	}
}
