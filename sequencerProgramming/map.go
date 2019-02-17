package main

import "fmt"

type PersonInfo struct {
	ID      string
	Name    string
	Address string
}

func main() {

	var persionDB map[string]PersonInfo //声明MAP
	persionDB = make(map[string]PersonInfo)
	persionDB["12345"] = PersonInfo{"12345", "Jack", "Beijing, China"} //persionDB["12345"] 只能用双引号
	persionDB["007"] = PersonInfo{"007", "007", "Guaongdong, China"}   //persionDB["12345"] 只能用双引号

	persion, ok := persionDB["12345"]
	//ok 返回的是一个布尔类型
	if ok {
		fmt.Printf("Persion find name:%s", persion.Name)
	} else {
		fmt.Println("persion not found")
	}
	fmt.Println()
	fmt.Println("------------------------------------------------------------------------")
	fmt.Println("len of persionDB:", len(persionDB))
	//删除元素
	delete(persionDB, "007")
	fmt.Println("len of persionDB:", len(persionDB)) //len：1
}
