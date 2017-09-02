package main

import (
	"time"
	"fmt"
)


func main() {

	t := "2017-08-31 14:15:00"
	tp, _ := time.Parse("2006-01-02 15:04:05", t)
	fmt.Println(tp)

	t2, _ := time.LoadLocation("Asia/Manila")
	t3 := time.Now().In(t2)
	fmt.Println("Location: ", t3.Location(), " Time: ", t3)

	td := t3.Sub(tp)
	fmt.Println(td)
	if td > (time.Minute * 15) {
		fmt.Println("Greater than")
	}

}