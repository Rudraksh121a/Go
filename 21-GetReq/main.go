package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {
	fmt.Println("Hello from GetReq")
	url := "https://api.freeapi.app/api/v1/public/randomusers?page=1&limit=10"
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()
	fmt.Println("statuscode", response.StatusCode)
	content,_:=io.ReadAll(response.Body)
	fmt.Println(string(content))
	
}
