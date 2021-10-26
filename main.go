package main

import (
	"fmt"

	array "github.com/zinuhe/golang-array-slide-map/array"
	slide "github.com/zinuhe/golang-array-slide-map/slide"
)

func main() {
	fmt.Println("==============================ARRAYS==============================")
	array.Declaration()
	array.Iterator()
	array.CreateArray()

	fmt.Println("==============================SLIDES==============================")
	slide.Declaration()
	slide.Iterator()

	fmt.Println("===============================MAPS===============================")

}
