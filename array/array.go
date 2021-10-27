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

    // to print each element quoted.
    fmt.Printf("%q\n", intArray)
}

func Iterator() {
    fmt.Println("\nArray Iterator")

    fmt.Println("for")
    for i := 0; i < len(intArray); i++ {
        fmt.Println("i:", i, "val:", intArray[i])
    }

    fmt.Println("\nfor range")
    for i, v := range intArray {
        fmt.Println("i:", i, "val:", v)
    }
}

func CreateArray() {
    fmt.Println("\nArray Create a new one")

    fmt.Println("intArray[0:2]", intArray[0:2])
    fmt.Println("intArray[1:3]", intArray[1:3])
    fmt.Println("intArray[2:]", intArray[2:])
}
