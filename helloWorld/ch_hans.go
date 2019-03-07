package main

import "fmt"

func main() {

	//1 先将字符串转成 []rune 切片
	//2 再用常规方法进行遍历
	//  中文字符串遍历
	str := "我爱中国123"
	str2 := []rune(str)
	for i := 0; i < len(str2); i++ {
		fmt.Printf("str[%d]=%c\n", i, str2[i])
	}
}
