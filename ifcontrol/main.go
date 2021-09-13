package main

import "fmt"

func main() {
	score := 1
	jomblo := true

	if score > 80 {
		fmt.Println("Grade A")
	} else if score > 50 {
		fmt.Println("Grade B")
	} else {
		if jomblo == true {
			fmt.Println("Ngenes")
		} else {
			fmt.Println("Gblk")
		}
	}
}

// if
// if else
// else if
// nested if
