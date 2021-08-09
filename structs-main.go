package main

import (
	"log"
	"time"
)

// variables with first letter capitalized are available outside of this package
type User struct {
	FirstName   string
	LastName    string
	PhoneNumber string
	Age         int
	BirthDate   time.Time
}

func main() {
	user := User{
		FirstName: "Brandon",
		LastName:  "Green",
	}

	log.Println(user.FirstName)
	log.Println(user.LastName)
}
