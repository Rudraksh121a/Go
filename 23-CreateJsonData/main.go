package main

import (
	"encoding/json"
	"fmt"
)

type course struct {
	Name     string   `json:"coursename"`
	Age      int      `json:"age"`
	Platform string   `json:"platform"`
	Password string   `json:"-"`
	Tags     []string `json:"tags,omitempty`
}

func main() {
	fmt.Println("Hello from CreateJsonData")
	encodeJson()

}

func encodeJson() {

	rkCourses := []course{
		{"Python Course", 233, "rk", "asdf", []string{"web", "backend"}},
		{"Java Course", 233, "rk", "asdf", []string{"web", "backend"}},
		{" Go Course", 233, "rk", "asdf", []string{"web", "backend"}},
	}

	//packege as json

	finalJson, err := json.MarshalIndent(rkCourses, "", "\t")
	if err != nil {
		panic(err)
	}
	fmt.Println(string(finalJson))
}
