package main

import (
	"fmt"
	"time"
)

func main() {
	go Add()
	for i := 0; i < 10; i++ {
		fmt.Println("first Add")
		time.Sleep(time.Second)

	}
}

func Add() {
	for {
		fmt.Println("new task Add() ")
		time.Sleep(time.Second)
	}
}
