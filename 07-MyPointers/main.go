package main

import "fmt"

func main() {
	fmt.Println("Welcome to Pointers")

	var MYPTR *int

	fmt.Println("Value of Pointer is : ",MYPTR)  //if we not define any value so  `Value of Pointer is :  <nil>`
	


	mynum:=4
	var ptr = &mynum
	fmt.Println("Value of  Actual Pointer Address is : ",ptr) //address of that value
	fmt.Println("Value of  Actual Pointer is : ",*ptr)   //actaul value

	
}
