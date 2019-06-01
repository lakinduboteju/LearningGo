# Learning Go
My practice code from Linkedin Learning course of Learning Go.

## 1. Getting Started

* `Compiled` language

    Not an interpreted language. Not translated in to an intermediate code like byte code in Java. Just directly compiled in to machine code bound to a particular platform and an architecture.

* `Statically typed` language

    A variable has a specific type always. No need to explicitly declare the type always, sometime compiler is capable of declaring variable for us.

* Statically linked runtime

    Every program compiled in to executable has it's own statically linked runtime. User do not need to have an installed runtime or a virtual machine to run Go compiled program on their platform.

* Object-oriented 'kind of'

    Go supports object oriented features like `interfaces`. Programmer can define custom `types`. A Type should implement at least 1 interface. Can add functions to a type (`methods`). Can create custom data structures (`structs`) and they can have member fields. Types and structs in Go like classes in languages like Java or C#.

    But there's `no type inheritance` in Go (no super type/class sub type /class).

    `No method or operator overloading`. Each function should have a unique name within a package or a type (class).

    `No exception handling` (no try catch or finally clauses). Error objects are being return from functions. Programmer has to handle `error` objects and handle errors with conditional logic.

    `No implicit type conversion`. All type conversions have to be done explicitly by programmer.

* `Package` names and `variable` names start in lower-case letter and contiue in camel-case.

* `Exported` (public) functions/methods or fields of Types starts with upper-case letters. That's how they become public.

* `Private` methods/fields starts with lower-case letter.

* `Semicolon` is optional. But sometimes it is required to use semicolons to prevent the code mis-interpreted by compiler. Lexer that analyses the code before compiling, adds `;` after each statement, because Go compiler expects `;` for each statement. Most of the time, Go takes line break as the end of the preceding statement.

* In code blocks, starting brace should be on the same line as preceding statement. Go does sensitive to the format of the written code. Moreover, sometimes parentheses are not needed.

    ``` go
    for i := 0; i < 10; i++ {
        // Loop code here
    }
    ```

* `builtin` is a default imported package for every Go file, which contains some built-in functions and constants. Programmer does not have to import it explicitly.

## 2. Essential Go Skills

Print string and value of other data types.
Get inputs from stdin. Read strings from stdin and convert to other types.

### 1. [HelloWorld](src/HelloWorld/helloworld.go)

* Structure of a Go program (`package` declaration, `import` required packages/libraries, `main()` function).

* How to call exported (public) functions from a package (library).

``` go
// Every file belongs to a package.
// File that contains main() function should belong to main package.
// This main is an identifier not belong to any type.
package main

// Package imports (just like library imports)
// There's a package called builtin, which contains a common set of functions and constants.
// builtin package don't have to be imported explicitly.
// But all other packages should be imported explicitly.
import (
    "fmt"
    // Other package names comes here (as strings).
)

// main() function of the program.
// All functions starts with func identifier.
// main() function takes no arguments.
func main() {

    // Println() function is an exported function from fmt package.
    // Exported = public
    // No need for a semicolon at the end of the statement.
    fmt.Println("Hello, World!")
    
    // Your other program code comes here.
}
```

### 2. [Strings](src/Strings/strings.go)

* Declare and assign variables (`explicit` variable declaration, `implicit` variable declaration, `short` variable definition).

``` go
// Explicit variable declaration
var str string = "some"
// Implicit variable declaration
var str = "some"
// Short variable definition = Shortand of declaring and assigning variable
str := "some"
// Declare and assign value
var str string
str = "some"
```

* Function defination.

``` go
// All functions starts with func identifier.
// Then function name. One package, type or struct cannot have functions with same name (no overloading).
// Input arguments are declared next in parentheses.
// At the end there's return type or multiple types (multiple types in parentheses).
func some(i int) int {
    return i+1
}

// Multiple return values (may be different types)
func some(i int) (int, float64) {
    // Type conversion has to be explicit as follows.
    return i+1, float64(i)+2.0
}
```

* Error object as one of multiple return values.

``` go
// Declare and assign multiple values returned from a function.
var strLen int = 0
var err error = nil
strLen, err = fmt.Println("some str")

// Using short variable declaration make sense in this situation.
strLen, err := fmt.Println("soem str")

// If we want to ignore error object, just use _ instead of variable name.
strLen, _ := fmt.Println("some str")
```

* String formatting.

``` go
str1, str2, str3 := "one", "two", "three"
// Format and print to stdout
fmt.Printf("%v, %v, %v\n", str1, str2, str3);
// Print variable type
fmt.Printf("%T\n", str1)
// Format and print floating point number with 2 decimal places
f := 3.1468965
fmr.Printf("%.2f\n", f)
// Format and create string
str := fmt.Sprintf("%v, %.2f, %T", f, f, f)
```

### 3. [ConsoleInput](src/ConsoleInput/consoleinput.go)

* Read string line from stdin. Convert read strings in to different types.

``` go
import (
    "bufio"
    "os"
    "strconv"
)

reader := bufio.NewReader(os.Stdin)
str, err := reader.ReadString('\n')
if nil != err {
    // Handle error.
} else {
    // Type conversion
    str = strings.TrimSpace(str)
    i, err1 := strconv.ParseInt(str, 10, 64)
    // err1 should be handled
    f, err2 := strconv.ParseFloat(str, 64)
    // err2 should be handled
}
```

* Read space separated tokens.

``` go
var s1 string
var s2 string
var i1 int
fmt.Scanln(&s1, &s2, &i1)
// Input: "one two 3\n"
// s1 = "one"
// s2 = "two"
// i1 = 3
```

### 4. [UsefulTypes](src/UsefulTypes/usefultypes.go)

* Builtin types

``` go
uint8, uint16, uint32, uint64
int8, int16, int32, int64
float32, float64
bool
string
byte // alias to uint8
int // will be int32 or int64 depending on host
// Character value will be assigned to int32 type
var aChar1 rune = 'a' // rune is alias to int32
var aChar2 int32 = 'b'
```

* Constant notation

``` go
// var identifier changes to const identifier
const MY_CONSTANT int = 42
const MY_WORD = "Hello"
// := is available only for variables not for const
```

* Operations on string variables

``` go
import (
    "strings"
)

// string is immutable
str1 := "hello, world"
str2 := strings.ToUpper(str1)
str3 := strings.Title(str1)

// string comparison
isEqual := str1 == str2
// string comparison ignore case
isEqual = strings.EqualFold(str1, str2)
// check if string contains sub-string
strings.Contains(str1, "hello")
```
