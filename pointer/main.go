package main

import "fmt"

func main() {
	numberA := 5
	numberB := &numberA // Lokasi alamat memori numberA disimpan di memori

	fmt.Println(numberA)
	fmt.Println(numberB)
	fmt.Println(*numberB) // Menampilkan value dari alamat memori

	*numberB = 10
	fmt.Println(*numberB)
	fmt.Println(numberA)

	fmt.Println("========")

	var number1 int = 5
	var number2 *int = &number1 // * didepan digunakan untuk menyimpan alamat memori

	fmt.Println(number1)
	fmt.Println(number2)

	number1 = 20

	fmt.Println(number1)
	fmt.Println(*number2)

	fmt.Println("========")

	number := 5
	fmt.Println("Alamat Memori : ", &number)
	fmt.Println("Nilai Awal : ", number)

	change(&number, 100)
	fmt.Println("Alamat Memori : ", &number)
	fmt.Println("Nilai Akhir : ", number)
}

func change(old *int, new int) {
	*old = new
	fmt.Println("Alamat Memori : ", old)
	fmt.Println("Di dalam func : ", *old)
}
