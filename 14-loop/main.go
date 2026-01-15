package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome To loop")

	list := []int{1, 23, 34, 43, 43, 43, 4, 43, 454, 53, 45, 345, 345}

	for i := 1; i < len(list); i++ {
		fmt.Printf("i= %v\n", list[i])
	}

	for i, data := range list {
		fmt.Printf("index %v  and data%v\n", i, data)
	}
}
