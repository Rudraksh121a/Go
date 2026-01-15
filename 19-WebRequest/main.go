package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Welcome To WebRequest")
	url := "https://api.freeapi.app/api/v1/public/quotes/quote/random"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	fmt.Printf("Type is %T", res)
	defer res.Body.Close()

	data, _ := io.ReadAll(res.Body)
	fmt.Println(string(data))
}
