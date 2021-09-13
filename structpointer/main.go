package main

import "fmt"

type Student struct {
	ID   int
	Name string
	GPA  float64
}

func (student *Student) graduate() {
	student.Name = student.Name + ", S.Kom."
	// fmt.Println(student.Name)
}

func main() {
	student := Student{1, "Hafiz Arrahman", 3.74}
	fmt.Println(student.Name)

	// graduate(&student) // Function
	student.graduate() // Method
	fmt.Println(student.Name)
}

func graduate(student *Student) {
	student.Name = student.Name + ", S.Kom."
	// fmt.Println(student.Name)
}
