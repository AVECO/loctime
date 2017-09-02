package main

import (
	"time"
	"fmt"
)


func main() {

	tLoc, _ := time.LoadLocation("Asia/Manila")

	t := "2017-09-03 02:00:00"
	tp, _ := time.ParseInLocation("2006-01-02 15:04:05", t, tLoc)
	fmt.Println(tp)

	t3 := time.Now().In(tLoc)
	fmt.Println("Location: ", t3.Location(), " Time: ", t3)

	td := t3.Sub(tp)
	fmt.Println(td)
	if td > (time.Minute * 15) {
		fmt.Println("Greater than")
	}

}
