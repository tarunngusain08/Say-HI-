package contracts

type RegisterUser struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}
