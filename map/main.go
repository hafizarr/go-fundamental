package main

import "fmt"

func main() {
	var myMap map[string]int
	myMap = map[string]int{}

	myMap["coba1"] = 1
	myMap["coba2"] = 2
	myMap["coba3"] = 3
	myMap["coba4"] = 4
	myMap["coba3"] = 10

	fmt.Println(myMap["coba3"])

	myMapIsi := map[string]string{
		"ruby": "is beautiful",
		"go":   "is super fast",
		"java": "jawa",
	}

	for _, v := range myMapIsi {
		fmt.Println(v)
	}

	fmt.Println("==============")

	// delete(myMapIsi, "go") // Menghapus value

	for _, v := range myMapIsi {
		fmt.Println(v)
	}

	fmt.Println("==============")

	// value := myMapIsi["python"]
	// fmt.Println(value)

	value, isAvailable := myMapIsi["go"]
	fmt.Println(isAvailable)
	fmt.Println(value)

}
