package main

import (
	"fmt"
	"sync"
)

// 协程直接的同步操作可以用sync.WaitGroup，也可以用chan通道
func main() {
	var wg sync.WaitGroup

	c := make(chan int, 5)
	go func() {
		wg.Add(1)

		for i := 0; i < 5; i++ {
			c <- i
		}
		close(c)
		wg.Done()
	}()
	wg.Wait()
	for v := range c {
		fmt.Println(v)
	}
}
