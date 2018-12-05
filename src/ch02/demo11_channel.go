package main

import "fmt"

func main() {
	ch := make(chan string)
	ch1 := make(chan string, 2)
	ch1 <- "ch1"
	go func() {
		ch <- "str"
	}()

	fmt.Println(<-ch)
	fmt.Println(<-ch1)
}
