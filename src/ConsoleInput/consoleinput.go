package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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
	f, err := strconv.ParseInt(str, 10, 64)
	// Handling error
	if nil != err {
		fmt.Println("Error :", err)
	} else {
		fmt.Println("Number is", f)
	}
}
