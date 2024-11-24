package models

type User struct {
	name  string
	email string
}

func All() {
	// Here return all the users from the db
}

func CreateUser() (User, error) {
	// implement creation logic
	return User{}, nil
}
