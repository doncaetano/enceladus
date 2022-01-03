package dtos

type UserDTO struct {
	Id        string `json:"id"`
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	IsActive  bool   `json:"isActive"`
	CreatedAt string `json:"createdAt"`
	UpdatedAt string `json:"updatedAt"`
}

type CreateUserDTO struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
}

type Repo interface {
	CreateUser(data *CreateUserDTO) (*UserDTO, error)
	FindByEmail(email string) (*UserDTO, error)
	FindById(id string) (*UserDTO, error)
}
