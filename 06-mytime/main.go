package main

import (
	"fmt"
	"time"
)

func main() {
	Presenttime := time.Now()
	fmt.Println(Presenttime)

	fmt.Println(Presenttime.Format("01-02-2006 15:04:05 Monday"))
}
