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
	fmt.Print("Enter your name here: ")
	// Read stdin until line break
	str, _ := reader.ReadString('\n')
	fmt.Print("Hello, ", str)

	// Read an int from stdin
	fmt.Print("Enter a integer here: ")
	str, _ = reader.ReadString('\n')
	// Trim away left and right white spaces
	str = strings.TrimSpace(str)
	// Convert string value to int
	i, err := strconv.ParseInt(str, 10, 64)
	// Handling error
	if nil != err {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Integer is", i)
	}

	// Read an float from stdin
	fmt.Print("Enter a float here: ")
	str, _ = reader.ReadString('\n')
	// Trim away left and right white spaces
	str = strings.TrimSpace(str)
	// Convert string value to float
	f, err := strconv.ParseFloat(str, 64)
	// Handling error
	if nil != err {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Float is", f)
	}

	// Read space seperated values
	fmt.Print("Enter a string, int and a float seperated by spaces: ")
	str1, i1, f1 := "", 0, 0.0
	// Scanln() reads space seperated values.
	// They can be in different types.
	fmt.Scanln(&str1, &i1, &f1)
	fmt.Println("The string is", str1)
	fmt.Println("The integer is", i1)
	fmt.Println("The float is", f1)
}
