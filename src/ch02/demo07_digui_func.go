package main

import "fmt"

func main() {
	//递归函数
	fmt.Println(factorial(10))
}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	return n * factorial(n-1)
}
