package main

import "fmt"

func main() {
	fmt.Println("Math")
	i1, i2, i3 := 12, 34, 54
	intSum := i1 + i2 + i3
	fmt.Println("int sum:", intSum)

	f1, f2, f3 := 23.5, 65.1, 76.3
	floatSum := f1 + f2 + f3
	fmt.Println("float sum: ", floatSum)

	total := float64(intSum) + floatSum
	fmt.Println("total: ", total)

	prod := float64(intSum) * floatSum
	fmt.Println("total: ", prod)

	div := float64(intSum) / floatSum
	fmt.Println("imp: ", div)

		rest := int64(floatSum) % int64(intSum)
	fmt.Println("imp: ", rest)
}
