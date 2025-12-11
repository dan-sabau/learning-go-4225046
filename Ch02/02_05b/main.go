package main

import (
	"fmt"
	"math"
)

func main() {

	f1, f2, f3 := 23.5, 65.1, 76.3
	sum := f1 + f2 + f3
	fmt.Println("Float sum:", sum)

	sum = math.Round(sum*100)/100
	fmt.Printf("noua suma: %v\n", sum)

	fmt.Println("PI: ", math.Pi)

	raza := 12.1
	circumferinta := 2 * raza * math.Pi
	fmt.Printf("circumferinta cu 2 zecimale: %.2f\n", circumferinta)

}
