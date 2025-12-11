package main

import (
	"fmt"
	"time"
)

func main() {

	fmt.Println("Dates and times")

	t:= time.Date(2009,time.November, 10, 23,0,0,0, time.UTC)
	fmt.Printf("go lansare: %s\n", t)

	n:= time.Now()
	fmt.Printf("ora e: %s\n", n)
	fmt.Printf("tipul e: %T\n", n)
	fmt.Printf(n.Format(time.ANSIC) + "\n")

	maine := n.AddDate(0,0,1)
	fmt.Printf("maine e: " + maine.Format(time.ANSIC) + "\n")

	format := "Mon 2006/01/01"
		fmt.Printf("maine e: " + maine.Format(format) + "\n")


}
