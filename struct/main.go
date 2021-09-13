package main

import (
	"fmt"
	"struct/management"
)

func main() {
	// Seter
	user := management.User{}
	user.ID = 1
	user.FirstName = "Hafiz"
	user.LastName = "Arrahman"
	user.Email = "hafiz@gmail.com"
	user.IsActive = true

	user2 := management.User{}
	user2.ID = 2
	user2.FirstName = "Hafiz2"
	user2.LastName = "Arrahman2"
	user2.Email = "hafiz2@gmail.com"
	user2.IsActive = true

	// Urutan bisa acak
	userIsi := management.User{ID: 3, FirstName: "Annisa", LastName: "Noviani", Email: "nisanvani@gmail.com", IsActive: true}

	// Harus urut
	userIsi2 := management.User{
		4,
		"Nisa",
		"Novi",
		"novi@gmail.com",
		true,
	}

	// Geter
	fmt.Println(user.FirstName)
	fmt.Println(user2)
	fmt.Println(userIsi)
	fmt.Println(userIsi2)

	fmt.Println("==========")
	fmt.Println(management.DisplayUser(userIsi2))

	fmt.Println("==========")
	users := []management.User{user, user2}
	group := management.Group{"Gamer", user, users, true}

	management.DisplayGroup(group)

	fmt.Println("==========")
	// Method

	result := user.DisplayMUser()
	fmt.Println(result)
	fmt.Println(user2.DisplayMUser())

	group.DisplayMGroup()
}
