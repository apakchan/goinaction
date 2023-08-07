package main

import (
	"../counters"
	"../entities"
	"fmt"
)

func main() {
	counter := counters.New(10)
	fmt.Println(counter)
	user := entities.User{Name: "name", Gender: "male"}
	fmt.Println(user)

	admin := entities.Admin{Level: "admin"}
	admin.Name = "name"
	admin.Gender = "female"
}
