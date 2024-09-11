package example

type User struct {
	username string
	password string
	currency int
}

var Users = []User{
	{"Bob", "1234", 0},
	{"Rob", "4321", 1000},
	{"Alice", "lala", 15000},
}
