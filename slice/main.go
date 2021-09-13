package main

import "fmt"

func main() {
	var gamingConsoles []string

	gamingConsoles = append(gamingConsoles, "PS5")
	gamingConsoles = append(gamingConsoles, "PS4")
	gamingConsoles = append(gamingConsoles, "PS3")
	gamingConsoles = append(gamingConsoles, "PS2")
	gamingConsoles = append(gamingConsoles, "PS1")

	fmt.Println(gamingConsoles)

	gamingConsolesIsi := []string{
		"Nitendo",
		"Xbox",
	}

	fmt.Println(gamingConsolesIsi)

	for _, v := range gamingConsoles {
		fmt.Println(v)
	}
}
