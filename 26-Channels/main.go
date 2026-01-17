package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Hello from Channels")

	mych := make(chan int, 2)
	wg := &sync.WaitGroup{}
	// mych <- 5

	wg.Add(2)
	go func(ch chan<- int, wg *sync.WaitGroup) {
		mych <- 5
		mych <- 6
		defer wg.Done()
		close(mych)
	}(mych, wg)

	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChanelOpen := <-mych
		fmt.Println(val)
		fmt.Println(isChanelOpen)
		fmt.Println(<-mych)
		defer wg.Done()
	}(mych, wg)

	wg.Wait()
}
