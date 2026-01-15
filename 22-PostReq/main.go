package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	url := "https://api.freeapi.app/api/v1/kitchen-sink/http-methods/post"

	req, _ := http.NewRequest("POST", url, nil)

	req.Header.Add("accept", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := io.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
