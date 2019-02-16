package main

import (
	"calc/lib"
	"fmt"
	"os" //用于获得命令行参数os.Args
	"strconv"
)

var Usage = func() {
	fmt.Println("USAGE:calc command [arguments] ...")
	fmt.Println("\nThe commands are \n\tadd\tAddtition of two values.\n\tsqrt\tSqure rot of a non-negative value.")
}

func main() {
	args := os.Args[1:]
	if args == nil || len(args) < 2 {
		Usage()
		return
	}

	switch args[0] {

	case "add":
		if len(args) != 3 {
			fmt.Println("USAGE:calc add <interger1><interger2>")
			return
		}
		v1, err1 := strconv.Atoi(args[1])
		v2, err2 := strconv.Atoi(args[2])
		if err1 != nil || err2 != nil {
			fmt.Println("USAGE:calc add <interger1><interger2>")
			return
		}

		ret := lib.Add(v1, v2)
		fmt.Println("Result: ", ret)
	case "sqrt":

		if len(args) != 2 {
			fmt.Println("USAGE:calc sqrt <interger>")
			return
		}
		v, err := strconv.Atoi(args[1])
		if err != nil {
			fmt.Println("USAGE:calc add <interger>")
			return
		}

		ret := lib.Sqrt(v)
		fmt.Println("Result: ", ret)
	default:
		Usage()

	}
}
