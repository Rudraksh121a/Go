package main

import "fmt"

func main() {
	// defer fmt.Println("World")
	// defer fmt.Println("1")
	defer fmt.Println("kdfjls;")
	fmt.Println("Welcome To Defer")

	deferex()
}

// Welcome To Defer
// 2
// 1
// World

func deferex() {
	//here all go to stack because of defer
	for i := 0; i < 5; i++ {
		defer fmt.Println(i)
	}
}
