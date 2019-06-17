package main

import (
	"fmt"
)

// Arrays
func arrays() {
	var numbers1 [2]int
	numbers1[0] = 1000
	numbers1[1] = 1001
	fmt.Printf("numbers1 (%T) : %v\n", numbers1, numbers1)

	// This is a slice
	var numbers2 = []int{2000, 2001}
	fmt.Printf("numbers2 (%T) : %v\n", numbers2, numbers2)

	// This is an array
	var numbers3 = [2]int{3000, 3001}
	fmt.Printf("numbers3 (%T) : %v\n", numbers3, numbers3)

	var length int
	length = len(numbers3)
	fmt.Println("Length of numbers3 array :", length)
}

// Slices using slice literal
func slices1() {
	var colors = []string{"Red", "Green", "Blue"}
	fmt.Println("colors slice :", colors)

	// Append to existing slice
	colors = append(colors, "Puple")
	fmt.Println("colors slice after append :", colors)

	// Remove first item of slice
	colors = append(colors[1:len(colors)])
	// Same as
	// colors = append(colors[1:])
	fmt.Println("colors slice after removing first item :", colors)

	// Remove last item from slice
	colors = append(colors[0 : len(colors)-1])
	fmt.Println("colors slice after removing last item :", colors)
}

// Slices
func slices2() {
	numbers := make([]int, 2, 2)

	// To create slices we have to use make function.
	// var numbers = []int
	// This causes a panic, because memory is not initialized for declared slice.

	length := len(numbers)
	capacity := cap(numbers)
	fmt.Printf("Length of slice: %v, Capacity of slice: %v\n", length, capacity)

	numbers[0] = 4000
	numbers[1] = 4001

	// Append an item beyond initial size
	numbers = append(numbers, 4002)
	fmt.Printf("Length of slice: %v, Capacity of slice: %v\n", len(numbers), cap(numbers))

	// Remove 2nd item from slice
	numbers = append(numbers[:1], numbers[2:]...)
	// Same as
	// numbers = append(numbers[:1], numbers[2], numbers[3], numbers[4], and so on)
	fmt.Println("numbers slice after removing 2nd item :", numbers)
}

// Slices iterate and remove
func slices3() {
	var numbers []int
	numbers = append(numbers, 1, 2, 3)
	fmt.Println("Initial slice:", numbers)

	// Iterate through slice
	// For slices range gives index and value
	for i, _ := range numbers {
		// Remove 2nd item
		if 1 == i {
			numbers = append(numbers[:i], numbers[i+1:]...)
		}
	}
	fmt.Println("slice after removing 2nd item:", numbers)
}

// Maps
func maps() {
	states := make(map[string]string)
	states["WA"] = "Washington"
	states["OR"] = "Oregon"
	states["CA"] = "California"
	fmt.Println("map:", states)

	delete(states, "OR")
	fmt.Println("map after deleting OR:", states)

	// Iterate through map
	// For maps range gives key and value
	for k, _ := range states {
		// delete an item
		if k == "CA" {
			delete(states, k)
		}
	}
	fmt.Println("map after deleting CA:", states)

	// a slice to store keys
	keys := make([]string, len(states))
	i := 0
	// Here range give only the key
	for k := range states {
		keys[i] = k
		i++
	}
	fmt.Println("keys:", keys)
}

func main() {
	fmt.Println("<====== arrays ======>")
	arrays()

	fmt.Println("<====== slices 1 ======>")
	slices1()

	fmt.Println("<====== slices 2 ======>")
	slices2()

	fmt.Println("<====== slices 3 ======>")
	slices3()

	fmt.Println("<====== maps ======>")
	maps()
}
