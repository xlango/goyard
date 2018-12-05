package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	/*var wg sync.WaitGroup
	wg.Add(2)
	go a(&wg)
	go a(&wg)
	wg.Wait()*/

	for i := 0; i <= 2; i++ {
		go func(v int) {
			fmt.Println(v)
		}(i)
		time.Sleep(1 * time.Second)
	}
}

func a(wg *sync.WaitGroup) {
	for i := 1; i <= 2; i++ {
		fmt.Println(i)
	}
	wg.Done()
}
