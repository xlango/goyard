package main

import "fmt"

type Person struct {
	name string
	age  int
}

func main() {
	a := 1
	var ip *int
	ip = &a
	fmt.Printf("%T , %v \n", ip, ip)
	fmt.Printf("%T , %v \n", *ip, *ip)
	fmt.Println(ip, &ip, *ip, *&ip, &*ip)

	s1 := Person{"x1", 11}
	var b *Person = &s1
	fmt.Println(b)
	fmt.Println(b.name, b.age)
	fmt.Println(&b.name, &b.age)

	c := 10
	d := &c
	fmt.Println(change(d))

	e := [3]int{1, 2, 3}
	var f [3]*int
	for i := 0; i < len(e); i++ {
		f[i] = &e[i]
	}
	fmt.Println(f)
}

func change(num *int) int {
	*num = 20
	return *num
}
