package main

import (
	"fmt"
	"io"
	"os"
)

func main() {
	fmt.Println("Welcome to Files")
	content := "print(\"Hello\")"
	makeAndWriteFile(content)
	readfile("test.py")

}

func makeAndWriteFile(content string) {
	file, err := os.Create("./test.py")

	if err != nil {
		panic(err)
	}

	length, error := io.WriteString(file, content)
	if error != nil {
		panic(error)
	}

	defer fmt.Println(length)
	file.Close()
}

func readfile(filename string) {
	databyte, error := os.ReadFile(filename)
	if error != nil {
		panic(error)
	}
	fmt.Println("Text data inside\n", string(databyte))

}
