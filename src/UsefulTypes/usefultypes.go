package main

import (
	"fmt"
	"strings"
)

func main() {
	// Explicitly typed declaration and assignment
	var anInteger8 int8 = -128
	var anInteger16 int16 = -32768
	var anInteger32 int32 = -2147483648
	var anInteger64 int64 = -9.223372036e+18
	fmt.Printf("int8 min = %v\nint16 min = %v\nint32 min = %v\nint64 min = %v\n\n",
		anInteger8, anInteger16, anInteger32, anInteger64)

	var anUInteger8 uint8 = 255
	var anUInteger16 uint16 = 65535
	var anUInteger32 uint32 = 4294967295
	var anUInteger64 uint64 = 1.844674407e+19
	fmt.Printf("uint8 max = %v\nuint16 max = %v\nuint32 max = %v\nuint64 max = %v\n\n",
		anUInteger8, anUInteger16, anUInteger32, anUInteger64)

	var aFloat32 float32 = -4294967295.55445677
	var aFloat64 float64 = -1.844674407e+19
	fmt.Printf("float32 = %v\nfloat64 = %v\n\n",
		aFloat32, aFloat64)

	var aBoolean bool = true
	// byte = uint8
	var aByte byte = 255
	// int = int64 in an 64 bit machine
	// int = int32 in an 32 bit machine
	var anInteger int = -9.223372036e+18
	// uint = uint64 in an 64 bit OS
	// uint = uint32 in an 32 bit OS
	var anUInteger uint = 1.844674407e+19
	fmt.Printf("bool = %v\nbyte (%T) = %v\nint (%T) = %v\nuint (%T) = %v\n\n",
		aBoolean, aByte, aByte, anInteger, anInteger, anUInteger, anUInteger)

	// Implicitly typed declaration and assignment
	anInteger_2 := 32768
	// If machine is 32 bit, type will be int32
	// If machine is 64 bit, type will be int64
	fmt.Printf("%T = %v\n\n", anInteger_2, anInteger_2)
	aFloat_2 := -1.844674407e+19
	// If machine is 32 bit, type will be float32
	// If machine is 64 bit, type will be float64
	fmt.Printf("%T = %v\n\n", aFloat_2, aFloat_2)
	// There's no char type. Character values are assigned to int32 type.
	aChar := 'a'
	fmt.Printf("%T = %v (%c)\n\n", aChar, aChar, aChar)

	// Constants
	// Explicit typing
	const anINTEGER int = 42
	// Implicit typing
	const aSTRING = "This is Go!"
	// There's no := in implicit typing when it comes to constants

	// Strings
	str1 := "An implicitly typed string"
	var str2 string = "an explicitly typed string"

	// ToUpper() function from strings package
	fmt.Println(strings.ToUpper(str1))
	// Title() function from strings package
	fmt.Println(strings.Title(str2))

	// String comparison
	lValue := "hello"
	uValue := "HELLO"
	// Indentifier for bool is t in formatting
	fmt.Printf("%v Equals %v? %t\n", lValue, uValue, lValue == uValue)
	// Non-case-sensitive string comparison
	fmt.Printf("%v Equals non-case-sensitive? %v? %t\n", lValue, uValue, strings.EqualFold(lValue, uValue))

	// String contains check
	fmt.Println(str1, "contains exp?", strings.Contains(str1, "exp"))
	fmt.Println(str2, "contains exp?", strings.Contains(str2, "exp"))
}
