package entities

type user struct {
	Name   string
	Gender string
}

type Admin struct {
	user
	Level string
}
