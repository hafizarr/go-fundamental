package main

import "fmt"

type BangunDatar interface {
	HitungLuas() int
}

type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int {
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar   int
}

func (persegiPanjang PersegiPanjang) HitungLuas() int {
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

type Asal struct {
	Angka int
}

func (asal Asal) HitungLuas() int {
	return 10001
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 5, Lebar: 4}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas : ", luas)

	persegi := Persegi{Sisi: 5}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas : ", luas)

	asal := Asal{Angka: 5}
	luas = SeberapaLuas(asal)
	fmt.Println("Luas : ", luas)
}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
}
