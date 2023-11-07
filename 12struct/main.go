package main

import "fmt"

type User struct {
	Name   string
	Email  string
	Status bool
	Age    int
}

func main() {

	// No inheritence in golang
	// No super or parent

	user1 := User{"Madhurendra Nath Tiwari", "dev.madhurendra@gmail.com", true, 22}
	user2 := User{"Sai prabhu", "sai.prabhu@gmail.com", false, 26}
	user3 := User{"Shivam Kumar", "shivam.kumar@gmail.com", true, 24}
	fmt.Println("User1 details : ", user1)
	fmt.Printf("User1 details are: %+v\n", user1)
	fmt.Printf("User %v has %v email : \n", user1.Name, user1.Email)

	users := make([]User, 3)
	users[0] = user1
	users[1] = user2
	users[2] = user3

	fmt.Println("List of all the user: ", users)
}
