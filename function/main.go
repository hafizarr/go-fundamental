package main

import "fmt"

func main() {
	fmt.Println(printMyResult("go"))
	fmt.Println(printMyResult("java"))
	fmt.Println(printMyResult("php"))

	result := add(10, 20)
	fmt.Println(result)

	fmt.Println("===========")
	// Kalau ga butuh salah satu variable bisa dihilangkan dengan _
	luas, keliling := calculate2(10, 2)
	fmt.Println(luas)
	fmt.Println(keliling)
}

// Nama func, input, output
func printMyResult(sentence string) string {
	newSentence := "Saya sedang belajar " + sentence
	return newSentence
}

// Mempersingkat deklarasi tipe data di  func
func add(number, numberTwo int) int {
	return number + numberTwo
}

// Mengembalikan 2 return
func calculate(panjang, lebar int) (int, int) {
	luas := panjang * lebar
	keliling := 2 * (panjang + lebar)

	return luas, keliling
}

// predifine return value
func calculate2(panjang, lebar int) (luas int, keliling int) {
	luas = panjang * lebar
	keliling = 2 * (panjang + lebar)

	return
}
