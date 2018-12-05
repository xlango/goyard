package main

import "fmt"

func main() {
	//闭包=函数+引用环境  Lambda表达式

	a := add()

	for i := 0; i < 6; i++ {
		fmt.Println(a(i))
	}

	b := func() func() int {
		i := 99
		return func() int {
			i++
			return i
		}
	}()
	fmt.Print(b())
}

func add() func(int) int {
	a := 0
	i := func(n int) int {
		a += n
		return a
	}

	return i
}
