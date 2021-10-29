package slide

import (
	"fmt"
)

var intSlice = []int{1, 2, 3, 4}
var sliceStrings = []string{"one", "two", "three", "four"}

func Declaration() {
	fmt.Println("\nSlice Declaration")

	// boolSlide := []bool{true, false, true}

	fmt.Println("sliceStrings:", sliceStrings)

	fmt.Printf("%T\n",  sliceStrings)
	fmt.Printf("%v\n",  sliceStrings)
	fmt.Printf("%+v\n", sliceStrings)
	fmt.Printf("%#v\n", sliceStrings)
	fmt.Printf("%+q\n", sliceStrings)
}

func Iterator() {
	fmt.Println("\nSlice Iterator")

	fmt.Println("for")
	for i := 0; i < len(intSlice); i++ {
		fmt.Println("i:", i, "val:", intSlice[i])
	}

	fmt.Println("\nfor range")
	for i, v := range intSlice {
		fmt.Println("i:", i, "val:", v)
	}

}
