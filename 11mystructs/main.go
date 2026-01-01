package main

import "fmt"
type User struct {
	Name string
	Email string
	Status bool
	Age int
}

func (u *User) changeStatus(status bool) {
	u.Status = status
}

func main() {
	fmt.Println("Intro to structs")

	// Go’s structs are typed collections of fields. They’re useful for grouping data together to form records.

	// no inheritance in Golang, No super and parent

	user1 := User{"farrow","user1@test.com",true,16}
	fmt.Println(user1)
	fmt.Printf("User1 details are: %+v\n", user1) //%+v prints field names and values, kind of verbose
	fmt.Printf("Name is: %v, and email is %v\n", user1.Name, user1.Email)
	user1.changeStatus(false)
	fmt.Printf("User1 details are: %+v\n", user1) //%+v prints field names and values, kind of verbose
}
