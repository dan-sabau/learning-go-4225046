package main

import (
	"fmt"
)

func main() {

	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy brown dog."
	aNumber := 3

	fmt.Println(str1, str2, str3)
	stringLenght, err := fmt.Println("valoare", aNumber)
	if err == nil {
		fmt.Println("lungime", stringLenght)
	}

	fmt.Printf("valoare numar: %v\n", aNumber)
	fmt.Printf("data type: %T\n", aNumber)
}
