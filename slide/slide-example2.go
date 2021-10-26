package slide

import (
	"fmt"
	// "log"
	// "time"
)

/*
func main() {
	fmt.Println("Welcome to Slide")

	// var sl = []int{1, 2, 3, 4, 5, 6, 7}
	sl1 := []int{1, 2, 3, 4, 2, 5, 6, 3, 7, 8, 9, 10, 11, 12, 1, 13, 14, 15, 16}
	var sl2 = sl1[1:]

	fmt.Println("sl1:", sl1)
	fmt.Println("sl2:", sl2)

	// for _, v := range sl1 {
	// 	fmt.Println(v)
	// }

	//Print out duplicates in sl1 slice
	start := time.Now()
	Duplicate1(sl1)
	elapsed := time.Since(start)
	log.Printf("Duplicate1 took %s", elapsed)

	start = time.Now()
	Duplicate2(sl1)
	elapsed = time.Since(start)
	log.Printf("Duplicate2 took %s", elapsed)

	start = time.Now()
	Duplicate3(sl1)
	elapsed = time.Since(start)
	log.Printf("Duplicate3 took %s", elapsed)
}
*/

// Example of finding duplicates in a slide

// Duplicate1 - 1 approach
func Duplicate1(sl []int) {
	for i, v1 := range sl {
		_sl := sl[i+1:]
		for _, v3 := range _sl {
			if v1 == v3 {
				fmt.Println("1-Duplicate:", v1)
				break
			}
		}
	}
}

func Duplicate2(sl []int) {
	for i, v1 := range sl {
		for j, v2 := range sl {
			if v1 == v2 && i < j {
				fmt.Println("2-Duplicate:", v1)
			}
		}
	}
}

func Duplicate3(sl []int) {
	for i, v1 := range sl {
		for j := i + 1; j < len(sl); j++ {
			v2 := sl[j]
			if v1 == v2 {
				fmt.Println("3-Duplicate:", v1)
				break
			}
		}
	}
}
