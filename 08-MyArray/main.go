package main

import "fmt"

func main() {
	fmt.Println("Welcome to Array")

	var FruitList [4]string
	var veglist = [3]string{"A", "B", "c"}
	FruitList[0] = "A"
	FruitList[2] = "C"

	fmt.Println("List is : ", FruitList)
	fmt.Println("veg is : ", veglist[0])
	fmt.Println("veg is : ", veglist[1])
	fmt.Println("veg is : ", veglist[2])
}
