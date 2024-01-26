package client

type User struct {
	Name, Idetifer, Occupation string
}

// Create a new client
func NewUser(name, idetifer, occupation string) User {
	return User{
		Name:       name,
		Idetifer:   idetifer,
		Occupation: occupation,
	}
}
