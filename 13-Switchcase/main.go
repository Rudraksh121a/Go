package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	fmt.Println("Welcome to Switchcase")
	rand.Seed(time.Now().UnixNano())
	diceNumber := rand.Intn(6) + 1
	fmt.Println("Your Number is ", diceNumber)

	switch diceNumber {
	case 1:
		fmt.Println("1 is come")

	case 2:
		fmt.Println("2 is come")
	case 3:
		fmt.Println("3 is come")
	case 4:
		fmt.Println("4 is come")
		fallthrough
	case 5:
		fmt.Println("5 is come")
	case 6:
		fmt.Println("6 is come")
	default:
		fmt.Println("Invalid Number")

	}
}
