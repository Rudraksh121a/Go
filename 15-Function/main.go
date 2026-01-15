package main

import "fmt"

func main() {
	fmt.Println("Welcome to Function")
	Greet("User")
	Greet1(1)
	caller()
	fmt.Println(adder(1, 2))
	d:=proadder(34, 5, 34, 435, 345, 34, 54, 5)
	fmt.Println(d)
}

func Greet(name string) {
	fmt.Println("i am a Greet Function")
	fmt.Println("name", name)
}
func Greet1(number int) {
	fmt.Println("i am a Greet1 Function")
	fmt.Println("name", number)

}

func caller() {
	Greet("new")
	Greet1(12)
}
func adder(a int, b int) int {
	return a + b
}

func proadder(a ...int) int { //...int make a array
	total := 0
	for _ ,val:=range a{
		total+=val
	}
	return total
}
