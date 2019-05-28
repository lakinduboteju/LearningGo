package main

import (
	"fmt"
)

func main() {
	// := is not assignment operator.
	// It's called short variable declaration.
	// It declares variable and assigns the value.
	// Here we do not explicitly declare variable with the specific type.
	// Compiler does variable declaration for us.
	str1 := "The quick red fox"
	str2 := "jumped over"
	str3 := "the lazy dogs."

	// go functions are cable of returning more than one value.
	// fmt.Println() function returns length of printed string (integer)
	// and potential error object.
	strLen, err := fmt.Println(str1, str2, str3)
	if nil == err {
		fmt.Println("String length:", strLen)
	}

	// If you declare a variable, you must use it.
	// Otherwise it leads to a compilation error.
	// If you don't want to use it just use _ instead of
	// variable name.
	strLen, _ = fmt.Println(str1, str2, str3)
	fmt.Println("String length:", strLen)

	aNumb := 24
	isTrue := true

	// Use %v for value of a variable
	fmt.Printf("Number is %v, boolean is %v\n", aNumb, isTrue)
	// Use %T for type of the variable
	fmt.Printf("%T is %v, %T is %v, a %T is %v\n",
		aNumb, aNumb, isTrue, isTrue, str1, str1)
	// Convert int to float
	fmt.Printf("Number as a float %v\n", float64(aNumb))
	// For floating point number of 2 decimal points
	fmt.Printf("Number as a float %.2f\n", float64(aNumb))

	// Create a formatted string
	myString := fmt.Sprintf("Data types : %T, %T, %T", str1, aNumb, isTrue)
	fmt.Println(myString)
}
