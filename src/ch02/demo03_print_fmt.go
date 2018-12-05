package main

import "fmt"

type Student struct {
	x, y int
}

func main() {
	p := Student{1, 2}
	fmt.Printf("%T %v \n", p, p)

	a := ""
	a = fmt.Sprintf("%X", 123) //进制转换
	fmt.Println(a)
}
