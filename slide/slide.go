package slide

import (
	"fmt"
)

var intSlide = []int{1, 2, 3, 4}
var strSlide = []string{"one", "two", "three", "four"}

func Declaration() {
	fmt.Println("\nSlide Declaration")

	boolSlide := []bool{true, false, true}

	// fmt.Println("intSlide:", intSlide)
	fmt.Printf("intSlide: %T -- %v \n", intSlide, intSlide)

	fmt.Printf("strSlide: %T -- %v \n", strSlide, strSlide)

	fmt.Printf("boolSlide: %T -- %v \n", boolSlide, boolSlide)
}

func Iterator() {
	fmt.Println("Slide Iterator")

	fmt.Println("\nfor")
	for i := 0; i < len(intSlide); i++ {
		fmt.Println("i:", i, "val:", intSlide[i])
	}

	fmt.Println("\nfor range")
	for i, v := range intSlide {
		fmt.Println("i:", i, "val:", v)
	}

}
