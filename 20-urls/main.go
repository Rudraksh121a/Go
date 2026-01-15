package main

import (
	"fmt"
	"net/url"
)

const myurl string = "https://api.freeapi.app/api/v1/public/randomusers?page=1&limit=10"

func main() {
	fmt.Println("Welcome To urls")
	result, _ := url.Parse(myurl)
	fmt.Println(result.Scheme)
	fmt.Println(result.Host)
	fmt.Println(result.Path)
	fmt.Println(result.Port())
	fmt.Println(result.RawQuery)
}
