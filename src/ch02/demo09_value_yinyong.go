package main

import "fmt"

func main() {
	//值传递和引用传递
	//!!!非引用类型（int、string、bool、数组、struct），无法修改原变量数据；引用类型（指针、slice、map、chan）会改变原变量的值
	a := 10
	value(a)
	fmt.Println(a)

	yinyong(&a)
	fmt.Println(a)
}

func value(a int) {
	a = 90
}

func yinyong(a *int) {
	*a = 80
}
