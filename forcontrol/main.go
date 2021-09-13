package main

import "fmt"

func main() {
	// for i := 1; i <= 10; i++ {
	// 	fmt.Println("Saya tidak akan melakukan ini lagi: ", i)
	// }

	// i := 1
	// for i <= 10 {
	// 	fmt.Println(i)
	// 	i++
	// }

	title := "Golang the best language"
	for index, letter := range title {
		if index%2 == 0 {
			fmt.Println("index : ", index, "letter : ", string(letter))
		}
	}

	// _ digunakan untuk variable yang ga kepake
	for _, letter := range title {
		fmt.Println("index : ", 0, "letter : ", string(letter))
	}

	for index, letter := range title {
		letterSring := string(letter)
		// if letterSring == "a" || letterSring == "i" || letterSring == "u" || letterSring == "e" || letterSring == "o" {
		// 	fmt.Println("index : ", index, "letter : ", string(letter))
		// }

		switch letterSring {
		case "a", "i", "u", "e", "o": // sama seperti or
			fmt.Println("index : ", index, "letter : ", string(letter))
		}
	}
}
