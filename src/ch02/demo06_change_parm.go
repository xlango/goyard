package main

import "fmt"

func main() {
	//可变参数
	sum, count, name := AddSum("李四", 2, 3, 4)
	fmt.Println(sum, count, name)

	arr := []int{1, 2, 3, 4, 5}
	fmt.Println(AddSum("张三", arr...))
}

func AddSum(name string, nums ...int) (sum, count int, n string) {
	for _, value := range nums {
		sum += value
		count++
	}
	n = name
	return
}
