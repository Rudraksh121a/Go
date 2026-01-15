package main

import (
	"fmt"
)

func main() {
	fmt.Println("Welcome To Method")
	rudra := User{"rudraksh", "email.com", true, 21}
	fmt.Println("User Detail is ", rudra)
	rudra.Getstatus()
	rudra.Updatemail("sdflf@mail")
	fmt.Println(" new updatedx Detail is ", rudra)

	rudra.RefUpdatemail("sdflf@mail")
	fmt.Println(" new redf Detail is ", rudra)

}

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func (u User) Getstatus() {
	fmt.Println(u.Status)
}
func (u User) Updatemail(NewMail string) {

	u.Email = NewMail //it make copy orignal not change
}
func (u *User) RefUpdatemail(NewMail string) {

	u.Email = NewMail //it make copy orignal not change
}
