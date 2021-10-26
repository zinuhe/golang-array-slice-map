package array

import (
	"fmt"
)

var intArray = [4]int{1, 2, 3, 4}
var stringArray = [4]string{"one", "two", "three", "four"}

func Declaration() {
	fmt.Println("Simple Array Declaration")

	boolArray := [2]bool{true, false}

	// fmt.Println("intArray:", intArray)
	fmt.Printf("intArray: %T -- %v \n", intArray, intArray)

	fmt.Printf("stringArray: %T -- %v \n", boolArray, boolArray)

	fmt.Printf("boolArray: %T -- %v \n", boolArray, boolArray)
}

func Iterator() {
	fmt.Println("Array Iterator")

	fmt.Println("\nfor")
	for i := 0; i < len(intArray); i++ {
		fmt.Println("i:", i, "val:", intArray[i])
	}

	fmt.Println("\nfor range")
	for i, v := range intArray {
		fmt.Println("i:", i, "val:", v)
	}
}
