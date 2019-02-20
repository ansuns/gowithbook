package main

import (
	"book/objectOrientedProgramming/interface/one"
	"book/objectOrientedProgramming/interface/two"
	"fmt"
)

type File struct {
}

func (f *File) Writer(buf []byte) (n int, err error) {
	return n, err
}

func (f *File) Read(buf []byte) (n int, err error) {
	return n, err
}

func main() {

	var file1 two.IStream = new(File)
	var file2 one.ReadWriter = file1
	var file3 two.IStream = file2
	fmt.Println(file1)
	fmt.Println(file3)

}
