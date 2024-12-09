package main

import (
	"fmt"
	"math/big"

	// "math/rand"
	"crypto/rand"
)

func main() {
	fmt.Println("Welcome to math in golang")

	// var mynumberone int = 2
	// var mynumbertwo float64 = 4

	// fmt.Println("The sum is :",/mynumberone+ int(mynumbertwo))
	// rand.Seed(time.Now().UnixNano())
	// fmt.Println(rand.Intn(5) + 1) //go 1 to 4

	rand, _ := rand.Int(rand.Reader, big.NewInt(5))
	fmt.Println(rand)
}
