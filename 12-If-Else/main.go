package main

import "fmt"

func main() {
	fmt.Println("Welcome to If-Else")
	loginCount := 21
	var result string

	if loginCount < 10 {
		result = "Regular User"
	} else if loginCount > 10 {
		result = "Watch Out"
	} else {
		result = "something"
	}

	fmt.Println(result)

	if num := 3; num == 3 {
		fmt.Println("num==3")
	} else {
		fmt.Println("num !=3")
	}

}
