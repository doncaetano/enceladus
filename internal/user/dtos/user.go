package dtos

type UserDTO struct {
	id        string `json:"id"`
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	email     string `json:"lastName"`
	isActive  bool   `json:"isActive"`
	createdAt string `json:"createdAt"`
	updatedAt string `json:"updatedAt"`
}

type CreateUserDTO struct {
	firstName string `json:"firstName"`
	lastName  string `json:"lastName"`
	email     string `json:"lastName"`
}

type Repo interface {
	CreateUser()
}
