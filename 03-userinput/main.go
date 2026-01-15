package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "Welcome to user input"
	fmt.Println(welcome)

	reader := bufio.NewReader(os.Stdin)
	fmt.Println("Enter the rating for our shop:")

	// comma  ok ||  comma error
	// input, _ := reader.ReadString('\n')  #if we dont care about error we use _ in the place of error
	// _, err := reader.ReadString('\n')  #if we dont care about input we use _ in the place of input
	input, _ := reader.ReadString('\n')
	fmt.Println("Thanks for rating,", input)
	fmt.Printf("Type of this rating is %T", input)
}
