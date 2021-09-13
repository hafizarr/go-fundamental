package main

import "fmt"

type Gamer struct {
	Name  string
	Games []string
}

func (gamer *Gamer) AddGame(game string) {
	gamer.Games = append(gamer.Games, game)
}

func main() {
	gamer := Gamer{Name: "Zelda"}

	gamer.AddGame("PES 2021")
	gamer.AddGame("FIFA 21")

	for _, game := range gamer.Games {
		fmt.Println(game)
	}
}
