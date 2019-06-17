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

* Arithmetic operators on numeric types

``` go
// Math operators
+  // addition
-  // subtraction
*  // multiplication
/  // division
%  // remainder

// Bitwise operators
&  // bitwise AND
|  // bitwise OR
^  // bitwise XOR
&^ // bit clear
<< // left shift
>> // right shift

++ // increment by 1
-- // decrement by 1
+= // add and assign
-= // subtract and assign
*= // multiply and assign
/= // divide and assign
%= // get remainder and assign

&= // get bitwise AND and assign
|= // get bitwise OR and assign
^= // get bitwise XOR and assign
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

* Interfaces

``` go
type Location interface {
    GetX() int
    GetY() int
}
```

* Implement an interface

``` go
type Point struct {
    X, Y int
}

func (p *Point) GetX() int {
    return p.X
}

func (p *Point) GetY() int {
    return p.Y
}

// Now type Point implements Location interface.
```

* Take any typed value as an input

``` go
func printany(v interface{}) {
    fmt.Println(v)
}

// Can pass any typed value to the function
printany(1)
printany(3.14)
printany("hello")
printany(struct{a, b int}{1, 2})
```

### Type conversion and assertion

* Type conversion example

``` go
// Location interface
type Location interface {
    GetX() int
    GetY() int
}

// Pointer type implements Location interface
type Point struct {
    X, Y int
}

// Implementing Location interface
func (p *Point) GetX() int {
    return p.X
}

// Implementing Location interface
func (p *Point) GetY() int {
    return p.Y
}

// Convert a Point to Location type
point := Point{X: 1, Y: 2}
location := Location(&point) // have to pass pointer of Point value
```

* Type assertion example 1

``` go
// Using type assertion to convert types
func timeMap(y interface{}) {
    // try type conversion
    z, ok := y.(map[string]interface{})
    // type conversion success?
	if ok {
		z["updated_at"] = time.Now()
	}
}
```

* Type assertion example 2

``` go
func printString(value interface{}) {
    // Do differently depending on type
	switch str := value.(type) {
	case string:
		fmt.Println(str)
	case Stringer:
		fmt.Println(str.String())
	}
}
```

* Type assertion example 3

``` go
if err != nil {
    // if clause with small declaration before condition
  if msqlerr, ok := err.(*mysql.MySQLError); ok && msqlerr.Number == 1062 {
    log.Println("We got a MySQL duplicate :(")
  } else {
    return err
  }
}
```
* Type assertion facts :-

    * Type assertion can be used only with `interface{}` typed values.

    * `type` keyword for type assertion can be used only with a switch-case.

## Control logic

### If-else

``` go
// A short declaration statement preceding the conditional expression.
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
  // or use break statement to break out of loop
  
  // continue statement can be used to
  // stop running the loop body midway and
  // continue to the next iteration
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

* Switch with no expression

``` go
// Here we use switch as an if-else chain
var BMI = 21.0 
switch {
case BMI < 18.5:
    fmt.Println("You're underweight")
case BMI >= 18.5 && BMI < 25.0:
    fmt.Println("Your weight is normal")
case BMI >= 25.0 && BMI < 30.0:
    fmt.Println("You're overweight")
default:
    fmt.Println("You're obese")	
}
```

* Conditional / Logical operators

``` go
&& // And
|| // Or
== // Equals
!= // Not equal
!  // Not
<  // Less than
>  // Greater than
<= // Less than or equal
>= // Greater than ot equal
```

## Memory allocation

### new

* Built-in function, allocates memory and does not initialize it, and it only zeros it.

* Returns pointer to allocated empty memory block.

``` go
func NewFile(fd int, name string) *File {
    // Allocate zero memory for a File struct
    f := new(File)
    // After allocating zero memory we have manually assign
    // values for fields.
    f.fd = fd
    f.name = name
    f.dirinfo = nil
    f.nepipe = 0

    // This is same as
    // f := File{fd, name, nil, 0}
    
    return f
}
```

### make

* Used only to create slices, maps, and channels.

* It returns an initialized value, not a pointer to the value.

``` go
// Make a slice
s := make([]string)
// Make a map
m := make(map[int]string)
```

## Collections

* Create collections

``` go
type Vertex struct {
    Lat, Long float64
}

// Array
a1 := [2]Vertex
a1[0] = Vertex{40.68433, -74.39967}
a1[1] = Vertex{37.42202, -122.08408}

// Array using literal
a2 := [2]Vertex{
    Vertex{
        40.68433, -74.39967
    },
    Vertex{
        37.42202, -122.08408
    }
}

// Slice
s1 := make([]Vertex, 2, 2)
s1[0] = Vertex{40.68433, -74.39967}
s1[1] = Vertex{37.42202, -122.08408}

// Slice using literal
s2 := []Vertex{
    Vertex{
        40.68433, -74.39967
    },
    Vertex{
        37.42202, -122.08408
    }
}

// Map
m1 := make(map[string]Vertex)
m1["Bell Labs"] = Vertex{40.68433, -74.39967}
m1["Google"] = Vertex{37.42202, -122.08408}

// Map using literal
m2 := map[string]Vertex{
	"Bell Labs": Vertex{
		40.68433, -74.39967,
	},
	"Google": Vertex{
		37.42202, -122.08408,
	},
}
```

* Append to collections

``` go
// Cannot append to arrays, because they are fixed size.

// Append to slice
s2 = append(s2, Vertex{20.5636, -133.46063})

// To maps just add new key value pair
m2["Home"] = Vertex{
    20.5636, -133.46063
}
```

* Remove from collections

``` go
// Cannot remove from arrays, because they are fixed size.

// Remove first 2 items from slice
s2 = append(colors[2:])
// Remove last 3 items from slice
s2 = append(colors[:len(colors)-3]
// Remove 5th items from slice
i := 5
s2 = append(colors[:i], colors[i+1:]...)

// Remove from map
delete(m2, "Google")
```
