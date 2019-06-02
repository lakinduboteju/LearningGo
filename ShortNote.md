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

### Executable code structure

Main entry point of the executable should be in `main` package. Function should be `main()`.

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

### Package imports

* Every compiled program (executable / library) made up of code spread across multiple packages.

* 3rd party libraries (packages) from net will be downloaded in to the `$GOPATH` path.

* Start your new program / library inside the `$GOPATH/src` directory.

    ``` bash
    $ cd $GOPATH/src
    # If you publish your work in github, create the path as this.
    $ mkdir -p github.com/<your username>/<project name>
    ```

``` go
import (
    // This package name is rand, and it's located in a directory called math
    "math/rand"
    // Most 3rd patry libraries and referred to by the location of its web URL.
    // Example : cryptography library from github
    // To download the library use following command:-
    // $ go get github.com/mattetti/goRailsYourself/crypto
    // The library will be downloaded to the path referred by GOPATH environment variable.
    "github.com/mattetti/goRailsYourself/crypto"
)
```

## Built-in types

* Numeric types

``` go
int // either 32 or 64 bits
uint // either 32 or 64 bits

int8, int16, int32, int64
uint8, uint16, uint32, uint64

float32, float64

complex64, complex128

byte // alias for int8
// rune is used to store a unicode character
rune // alias for int32
```

* Other build-in types

``` go
bool
string
```

* Examples of different typed variable instantiation

``` go
var ui uint = 18446744073709551615
var cplx complex128 = (2+3i)
var f float64 = 3.14
var c rune = 'a'

var b bool = true
var s string = "hello"
```

* Type conversion

    * `T(v)` conversion functions can be used to convert values between numeric types.

``` go
var i int = 42
var f float64 = float64(i)
var u uint = uint(f)
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

## Define constants

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

## Functions

* Pass by value - all values passed in and out of functions are passed by value.

* Passed variable's value will be copied to the local variable wihin the function.

* Go channels are also being passed by value, not by reference.

* To pass by reference, pointers can be passed and use `*` operator inside the function to access the value (dereference).

* To get the pointer of a value, use `&` operator.

``` go
func add(a, b int) int {
    return a + b
}

// Multiple return values
func location(city string) (float32, float32) {
    // some code
    return fLongitude, fLatitude
}
// Calling function with multiple return values
longitude, latitude := location("Colombo")

// Named return parameters
func location(city string) (fLongitude, fLatitude float32) {
    // some code
    return
}
```

## Defining Types

* Types

    * interfaces - are named collection of method signatures

    * structs - are named collection of data / fields / properties.

* Struct

``` go
type Point struct {
    X, Y int
}

var (
    p Point = Point{1, 2}
    q *Point = &Point{1, 2}
    r Point = Point{X: 1} // Y:0
    s Point = Point{} // X:0, Y:0
)
```

## Control logic

### If-else

``` go
// Below an assignment operation happens first.
if num := 9; num < 0 {
    fmt.Println(num, "is negative")
} else if num < 10 {
    fmt.Println(num, "has 1 digit")
} else {
    fmt.Println(num, "has multiple digits")
}
```

### Loops

``` go
// Traditional for loops
sum := 0
for i := 0; i < 10; i++ {
    sum += i
}

sum := 1
for ; sum < 1000; {
    sum += sum
}

// For as a while loop
sum := 1
for sum < 1000 {
    sum += sum
}

// Infinite loop
for {
  // do something in a loop forever
  // or break
}
```

### Switch-case

``` go
score := 7
// No need of using break at the end of every case block
switch score {
case 0, 1, 3:
    fmt.Println("Terrible")
case 4, 5:
    fmt.Println("Mediocre")
case 6, 7:
    fmt.Println("Not bad")
case 8, 9:
    fmt.Println("Almost perfect")
case 10:
    fmt.Println("hmm did you cheat?")
default:
    fmt.Println(score, " off the chart")
```

* Execute all following case blocks after a match using `fallthrough`

``` go
n := 4
// Even in a fallthrough switch-case,
// break can be used to break out of switch block.
switch n {
case 0:
    fmt.Println("is zero")
    fallthrough
case 1:
    fmt.Println("is <= 1")
    fallthrough
case 2:
    fmt.Println("is <= 2")
    fallthrough
case 3:
    fmt.Println("is <= 3")
    fallthrough
case 4:
    fmt.Println("is <= 4")
    fallthrough
case 5:
    fmt.Println("is <= 5")
    fallthrough
case 6:
    fmt.Println("is <= 6")
    fallthrough
case 7:
    fmt.Println("is <= 7")
    fallthrough
case 8:
    fmt.Println("is <= 8")
    fallthrough
default:
    fmt.Println("Try again!")
}
```

Output

```
is <= 4
is <= 5
is <= 6
is <= 7
is <= 8
Try again!
```
