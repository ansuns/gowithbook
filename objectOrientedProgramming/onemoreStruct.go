package main

import "fmt"

type Recta struct {
	x, y          float64
	width, height float64
}

func (r *Recta) Area() float64 {
	return r.height * r.width
}

func main() {
	//r := new(Recta)
	r := &Recta{0, 0, 4, 9}
	res := r.Area()
	fmt.Println(res)
}
