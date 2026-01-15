package main

import "fmt"

func main() {
	fmt.Println("Welocme to Maps")

	language := make(map[string]string)
	language["js"] = "javaScript"
	language["Py"] = "Python"
	language["cpp"] = "c++"

	// fmt.Printf("Type of language are %T \n", language)
	fmt.Printf("List of All language are %v\n", language)

	for key, value := range language {
		fmt.Printf("for key %v and value is %v\n", key, value)
	}

}
