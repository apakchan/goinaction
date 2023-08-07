package main

import "fmt"

func createSlice() {
	// 使用embedded make 函数
	// 规定了最大容量
	_ = make([]string, 5)
	_ = make([]int, 3, 5)                                  // 长度为 3 容量为 5, 不允许容量 < 长度
	_ = []string{"Red", "Blue", "Green", "Black", "White"} // 长度 容量都为 5
	_ = []string{99: "99"}                                 // 初始化第 100 个元素

	_ = make([]int, 0) // 空的整型切片, 长度为 0, 容量为 0, 地址指针为空
}

func useSlice() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3] // [1, 3)  长度为 2 (3 - 1) 容量为 4 (5 - 1) 5 是 srcSlice 的容量
	// 20 30
	printIntSlice(&newSlice)
	newSlice[0] = 100
	printIntSlice(&slice)    // 10 100 30 40 50
	printIntSlice(&newSlice) // 100 30
}

func sliceAppend() {
	slice := []int{10, 20, 30, 40, 50}
	newSlice := slice[1:3] // [1, 3)  长度为 2 (3 - 1) 容量为 4 (5 - 1) 5 是 srcSlice 的容量
	newSlice = append(newSlice, 100)
	printIntSlice(&newSlice) // 20 30 100
	printIntSlice(&slice)    // 10 20 30 100 50

	slice = append(slice, 200)
	slice[1] = 300
	printIntSlice(&slice)    // 10 300 30 100 50 200
	printIntSlice(&newSlice) // 20 30 100

	// 3 个索引
	// 长度为 3 - 2 = 1, 容量为 4 - 2 = 2
	newSlice1 := slice[2:3:4]
	printIntSlice(&newSlice1)

	/*
		尽量设置切片容量 = 长度，这样可以强制让新切片的第一个 append 操作创建新的底层数组
		好处是可以避免修改新切片时连同老切片修改的 bug, 避免不必要的错误
	*/
	source := []string{"a", "b", "c", "d", "e"}
	sliceString := source[2:3:3] // len: 1, cap: 1, content: ["c"]
	sliceString = append(sliceString, "z")
	printStringSlice(&source)      // a b c d e
	printStringSlice(&sliceString) // c z

	strings := append(source, sliceString...)
	printStringSlice(&strings) // len: 7, cap: 10, content: a b c d e c z
}

func createMultipleSlice() {
	// 两个 slice, 第一个 len: 1, cap: 1, content: 10, 第二个 len: 2, cap: 2, content: 100, 200
	slice := [][]int{{10}, {100, 200}}
	fmt.Printf("%v\n", slice)
	slice[0] = append(slice[0], 20)
	fmt.Printf("%v\n", slice)
}

func printIntSlice(slice *[]int) {
	fmt.Println("---------------------------------------------")
	fmt.Printf("length: %d, capicity: %d\n", len(*slice), cap(*slice))
	for _, item := range *slice {
		fmt.Printf("%d ", item)
	}
	fmt.Printf("\n")
}

func printStringSlice(slice *[]string) {
	fmt.Println("---------------------------------------------")
	fmt.Printf("length: %d, capicity: %d\n", len(*slice), cap(*slice))
	for _, item := range *slice {
		fmt.Printf("%s ", item)
	}
	fmt.Printf("\n")
}

func main() {
	//createSlice()
	//useSlice()
	//sliceAppend()
	createMultipleSlice()
}
