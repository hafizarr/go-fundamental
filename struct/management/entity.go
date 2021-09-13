package management

import "fmt"

type User struct {
	ID        int
	FirstName string
	LastName  string
	Email     string
	IsActive  bool
}

func (user User) DisplayMUser() string {
	return fmt.Sprintf("name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
}

func DisplayUser(user User) string {
	result := fmt.Sprintf("name : %s %s, Email : %s", user.FirstName, user.LastName, user.Email)
	return result
}

type Group struct {
	Name        string
	Admin       User
	Users       []User
	IsAvailable bool
}

func DisplayGroup(group Group) {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("USer Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}

func (group Group) DisplayMGroup() {
	fmt.Printf("Name : %s", group.Name)
	fmt.Println("")
	fmt.Printf("Member Count : %d", len(group.Users))
	fmt.Println("")
	fmt.Println("USer Name : ")
	for _, user := range group.Users {
		fmt.Println(user.FirstName)
	}
}
