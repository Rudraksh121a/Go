package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Welcome to the Slices")
	var flist = []string{}
	fmt.Printf("Type of flist is %T \n", flist)

	flist = append(flist, "A", "B", "C")

	fmt.Println("items of flist is \n", flist)
	fmt.Println("items with slices is \n", flist[1:])
	fmt.Println("items with slices is \n", flist[:1])
	fmt.Println("items with slices is \n", flist[:])

	highscore := make([]int, 4)
	highscore[0] = 1234
	highscore[1] = 123
	highscore[2] = 1223
	highscore[3] = 1223345
	// highscore[4] = 1223345 //give error out of range but we can use append

	highscore = append(highscore, 234, 234, 324, 324, 234, 32, 43, 4, 23)

	fmt.Printf("high scores with type %T \n", highscore)
	fmt.Println("high scores \n", highscore)

	sort.Ints(highscore)
	fmt.Println("high scores with sort \n", highscore)

	fmt.Println("high scores with sort \n", sort.IntsAreSorted(highscore))

	// Remoce the data for slices with index

	var courses = []string{"py", "js", "Ruby", "rn", "go"}
	fmt.Println(courses)
	var index int = 2
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("After delete index 2 " ,courses)

}
