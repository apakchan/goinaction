package main

import "fmt"

func createMap() {
	// key: string, value: int
	dict := make(map[string]int)

	dictNew := map[string]string{
		"Red":    "#da1337",
		"Orange": "#e95a22",
		"Dark":   "i dont know"}

	fmt.Printf("%#v\n", dict)
	fmt.Printf("%#v\n", dictNew)

	value, exists := dictNew["Red"]
	if exists {
		// true
		fmt.Println(value) // #da1337
	}

	for key, value := range dictNew {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	delete(dictNew, "Red")

	for key, value := range dictNew {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}

	removeColor(dictNew, "Dark")

	for key, value := range dictNew {
		fmt.Printf("key: %s, value: %s\n", key, value)
	}
}

// 函数间传递 map 不会制造 map 的副本, 当传递映射给一个函数，对这个映射做了修改，所有这个映射的引用都会觉察这个修改
func removeColor(dict map[string]string, key string) {
	delete(dict, key)
}

func main() {
	createMap()
}
