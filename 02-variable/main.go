package main

import "fmt"
const LoginToken string ="fdjksfhsjkodhr"
func main() {
	fmt.Println("Variable")
	var username string = "Rudraksh"
	fmt.Println(username)
	fmt.Printf("Variable of this types: %T \n", username)

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Variable of this types: %T \n", isLoggedIn)

	var smallvalue uint8 = 255
	fmt.Println(smallvalue)
	fmt.Printf("Variable of this types: %T \n", smallvalue)

	var smallfloat float64 = 255.5793485648396789435678945
	fmt.Println(smallfloat)
	fmt.Printf("Variable of this types: %T \n", smallfloat)

	// default value and aliases
	var anothervariableint int
	fmt.Println(anothervariableint)
	fmt.Printf("Variable of this types: %T \n", anothervariableint)

	var anothervariablestr string
	fmt.Println(anothervariablestr)
	fmt.Printf("Variable of this types: %T \n", anothervariablestr)

	// implicit type
	var website = "something"
	fmt.Println(website)

	// no var style  not allow to declear globle
	numberOfUser :=49343
	fmt.Println(numberOfUser)

	fmt.Println(LoginToken)
}
