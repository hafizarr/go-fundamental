package main

import "fmt"

func main() {
	var languages [5]string

	languages[0] = "hai0"
	languages[1] = "hai1"
	languages[2] = "hai2"
	languages[3] = "hai3"
	languages[4] = "hai4"

	fmt.Println(languages)
	length := len(languages) // Panjang array
	fmt.Println(length)

	languagesIsi := [5]string{"coba", "ah", "oh", "gini", "ya"} // Langusng diisi
	// languagesIsi := [5]string{
	// 	"coba",
	// 	"ah",
	// 	"oh",
	// 	"gini",
	// 	"ya",
	// }
	languagesIsiMales := [...]string{
		"coba",
		"ah",
		"oh",
		"gini",
		"ya",
		"males",
	}

	fmt.Println(languagesIsi)
	fmt.Println(languagesIsiMales)

	// sama seperti foreach
	for _, lang := range languagesIsiMales {
		fmt.Println(lang)
	}
}
