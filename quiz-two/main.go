package main

import "fmt"

func main() {
	// hitung rata-rata
	scores := [8]int{100, 80, 75, 92, 70, 93, 88, 67}

	n := len(scores)

	var total int
	for _, score := range scores {
		total = total + score
	}

	var average float64
	average = float64(total) / float64(n)

	fmt.Println(average)

	fmt.Println("=============")

	var goodScores []int
	for _, score := range scores {
		if score > 90 {
			goodScores = append(goodScores, score)
		}
	}

	fmt.Println(goodScores)
}
