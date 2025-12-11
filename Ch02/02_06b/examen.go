// Write your answer here, and then test your code.
package main

// Uncomment this import section to use required Go packages
 import (
 	"fmt"
 	"strconv"
 	"strings"
 )

// Change these boolean values to control whether you see 
// the expected answer and/or hints.
const showExpectedResult = true;
const showHints = true;

// calculate() returns the the result of adding 2 numbers 
// after parsing them from strings
func calculate(value1 string, value2 string) float64 {
	float1, err := strconv.ParseFloat(strings.TrimSpace(value1), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 1 must be a number")
	}
	float2, err := strconv.ParseFloat(strings.TrimSpace(value2), 64)
	if err != nil {
		fmt.Println(err)
		panic("Value 2 must be a number")
	}
	return float1 + float2
}