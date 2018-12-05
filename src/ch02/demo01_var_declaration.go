package main

import "fmt"

func main() {
	var (
		a int
		b string
		c float32
		d []int //切片（动态数组）
		e int32 = 100
		f [2]int
		g bool
		h func() string
		i = 2.11 //编译器自动推断数据类型
	)

	j := 1             //声明初始化变量，只能用在函数体能作为局部变量的声明
	j, _, k := 2, 2, 3 //"_"匿名变量不会占用内存空间，用户函数返回值，如果不希望返回某个值
	j, k = k, j        //多重赋值
	//uint8、uint16、uint32、uint64（分别表示0到2的n次方）在变成过程中数据声明合适的大小节约内存空间  int8...(负到正)

	fmt.Printf("%T %v \n", a, a)
	fmt.Printf("%T %q \n", b, b)
	fmt.Printf("%T %v \n", c, c)
	fmt.Printf("%T %v \n", d, d)
	fmt.Printf("%T %v \n", e, e)
	fmt.Printf("%T %v \n", f, f)
	fmt.Printf("%T %v \n", g, g)
	fmt.Printf("%T %v \n", h, h)
	fmt.Printf("%T %v \n", i, i)
	fmt.Printf("%T %v \n", j, j)
	fmt.Printf("%T %v \n", k, k)
}
