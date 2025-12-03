package request

type CreateUser struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
}
