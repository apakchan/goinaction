package main

import "fmt"

func createArrayAndPrint() {
	var array2 [5]int // 0 0 0 0 0
	array1 := [5]int{1, 2, 3, 4, 5}
	array3 := [...]int{3, 4, 5, 6}
	array4 := [5]int{1: 10, 3: 40} // 0 10 0 40 0
	for _, item := range array2 {
		// 0
		fmt.Printf("%d\n", item)
	}
	for _, item := range array1 {
		fmt.Printf("%d\n", item)
	}
	for _, i := range array3 {
		fmt.Printf("%d\n", i)
	}
	for _, i := range array4 {
		fmt.Printf("%d\n", i)
	}
}

func useArray() {
	array := [5]*int{0: new(int), 1: new(int)}
	*array[0] = 10
	*array[1] = 20
	// index 0: 10, index 1: 20
	// index > 1 取值会报 NPE
	fmt.Printf("index 0: %d, index 1: %d\n", *array[0], *array[1])

	var array1 [3]*string
	array2 := [3]*string{new(string), new(string), new(string)}
	*array2[0] = "Red"
	array1 = array2
	*array1[0] = "Blue"
	// arr1 0: Blue, arr2 0: Blue
	fmt.Printf("arr1 0: %s, arr2 0: %s\n", *array1[0], *array2[0])
}

func multiDimensionArray() {
	var array1 [4][2]int
	array2 := [4][2]int{{1, 2}, {3, 4}, {5, 6}, {7, 8}}
	array3 := [4][2]int{1: {1, 2}, 3: {3, 4}}
	array4 := [4][2]int{1: {1: 2}, 3: {0: 3}}
	printMultiDimensionArray(&array1)
	printMultiDimensionArray(&array2)
	printMultiDimensionArray(&array3)
	printMultiDimensionArray(&array4)
}

func printMultiDimensionArray(array *[4][2]int) {
	fmt.Println("------------------------------------------------")
	for _, arr := range array {
		for _, item := range arr {
			fmt.Printf("%d ", item)
		}
		fmt.Printf("\n")
	}
}

func main() {
	// createArrayAndPrint()
	// useArray()
	multiDimensionArray()
}
