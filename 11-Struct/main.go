package main

import "fmt"

func main() {
	fmt.Println("Welcome To Struct")
	rudra := User{"rudraksh", "email.com", true, 21}
	fmt.Println("User Detail is ", rudra)
	fmt.Println("User Detail is ", rudra.Name)
	fmt.Printf("User Detail is %+v ", rudra)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}
