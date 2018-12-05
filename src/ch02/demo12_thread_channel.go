package main

import "fmt"

func main() {
	ch := make(chan string)
	go worker(ch)
	fmt.Println(<-ch)
}

func worker(ch chan string) {
	fmt.Println("111111111")
	ch <- "chstring"
}
