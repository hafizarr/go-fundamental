package main

import "fmt"

func main() {
	students := []map[string]string{
		{
			"name":  "Agung",
			"score": "A",
		},
		{
			"name":  "Ling",
			"score": "A",
		},
		{
			"name":  "Mario",
			"score": "B",
		},
	}

	fmt.Println(students)

	for _, student := range students {
		fmt.Println(student["name"], " - ", student["score"])
	}
}
