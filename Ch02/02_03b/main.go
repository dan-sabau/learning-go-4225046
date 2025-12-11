package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Print("test: ")
	str, _ := reader.ReadString('\n')
	fmt.Printf("valoare: %v\n", str)
	
	fmt.Print("numar: ")
	str, _ = reader.ReadString('\n')
	f, err :=strconv.ParseFloat(strings.TrimSpace(str), 64)
	if err != nil {
		fmt.Println(err)
		
	} else {
		fmt.Println("valoare: ", f)
	}

}
