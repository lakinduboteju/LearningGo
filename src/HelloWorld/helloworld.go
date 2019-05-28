// Package declaration
// Startup file (file with main function) should always be in main package.
// main is an identifier not belong to any type like string.
// Not a requirement to match package name with dir structure like in Java.
package main

// Go libraries are organized as packages.
// There's a package called builtin which contains functions and constants
// that are always available and imported in to the code.
// No need explicitly import builtin package.
// But need to explicitly import every other package.
import (
	"fmt"
	"strings"
)

// Functions starts with identifier func.
// Main function always will be main() and doesn't take any args.
func main() {
	// To print Hello World, start with the package name fmt.
	// Then call exported function from fmt package called Println.
	// We say it's exported (public) because it starts with a upper-case letter.
	// Exported is equivalent to make method public like in Java.
	fmt.Println("Hello, World!")
	fmt.Println(strings.ToUpper("hello really loud!"))
}
