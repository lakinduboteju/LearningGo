# Everything needed to get started with Go

###### http://www.golangbootcamp.com/book

* Go is designed to make code as simple as possible, and increase code readability.

    * Language spec is very small.

    * Designed to write small amount of code.

* Compiled language

* Statically typed language

* Statically linked runtime

    * Go programs runs on Go runtime, which is linked to every Go executable.

    * Memory allocation, garbage collection, thread management are done by the runtime.

* Packages and Types

    * Code is organized in to packages and types.

## Go program structure

``` go
// Package declaration
package main

// Import packages
import (
    "fmt"
)

// Main function
func main() {
    fmt.Println("Hello, Go!")
}
```

## Define variables

### Declare one-by-one

``` go
// Explicit declaration
var firstName1 string
firstName1 = "Lakindu"
// Explicit declaration and assignment
var lastName1 string = "Boteju"

// Implicit declaration and instantiation
firstName2 := "Lakindu"
lastName2 := "Boteju"

var firstName3, lastName3 age3 = "Lakindu", "Boteju", 27
firstName4, lastName4 age4 := "Lakindu", "Boteju", 27
```

### Declare list of variables (using var statement)

``` go
// Explicit declaration
var (
    name string
    age int
)
// Assignment
name = "Lakindu Boteju"
age = 27
```

``` go
var (
    firstName, lastName string
    age int8
)

firstName = "Lakindu"
lastName = "Boteju"
age = 27
```

### Assign functions to variables

``` go
// Assign function to a variable
add := func(val1, val2 int) int {
    return val1 + val2
}

// Call function using variable
sum := add(1, 2)
```

# Define constants

``` go
const PI = 3.14
const (
    STATUS_OK = 200
    STATUS_CREATED = 201
    STATUS_ACCEPTED = 202
)

// Constants can only be numeric values, bool, character or string
const (
    PI = 3.14
    TRUTH = true
    LETTER = 'a'
    WORD = "hello"
)
```
