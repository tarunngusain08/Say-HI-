package user

const administrator = "admin"

type admin struct {
	name     string `json:"name"`
	userName string `json:"userName"`
	password string `json:"password"`
	email    string `json:"email"`
	phone    string `json:"phone"`
	role     string `json:"role"`
}

func NewAdmin(name, username, password, email, phone string) User {
	return &admin{
		name:     name,
		userName: username,
		password: password,
		email:    email,
		phone:    phone,
		role:     administrator,
	}
}

func (a *admin) GetName() string {
	return a.name
}

func (a *admin) GetUserName() string {
	return a.userName
}

func (a *admin) GetPassword() string {
	return a.password
}

func (a *admin) GetEmail() string {
	return a.email
}

func (a *admin) GetPhone() string {
	return a.phone
}

func (a *admin) GetRole() string {
	return a.role
}
