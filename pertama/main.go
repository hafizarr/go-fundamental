package main

import (
	"fmt"
	"pertama/calculaion"
)

// methode yang akan di jalankan di package main
func main() {
	fmt.Println("Halo, belajar golang")
	// sentence := TestAja()

	// fmt.Println(sentence)

	result := calculaion.Add(26, 29)
	fmt.Println(result)
}
