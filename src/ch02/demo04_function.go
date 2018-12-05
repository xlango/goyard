package main

import (
	"fmt"
	"math"
)

//type 自定义类型
type caseFunc func(int) int

func main() {

	/*slient1:= []int{ 1,2,3}
	for i,v:= range slient1 {
		fmt.Printf("%d %d \n",i,v)
	}*/

	//fmt.Print(b(1,A))

	//匿名函数
	func(data int) {
		fmt.Print(data + 1)
	}(100)
	result := func(data float64) float64 {
		return math.Sqrt(data)
	}(100.0)
	fmt.Print(result)
	i := func(data string) string {
		return data
	}
	fmt.Println(i("abc"))

	//匿名函数作为回调函数
	i2 := b(1, func(i int) int {
		return i + 2
	})
	fmt.Print(i2)
}

func A(i int) int {
	return i + 1
}

//函数作为参数
func b(i int, c caseFunc) int {
	return c(i)
}
