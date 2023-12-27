package user

const RoleCustomer = "Customer"

type Customer struct {
	name     string `json:"name"`
	userName string `json:"userName"`
	password string `json:"password"`
	email    string `json:"email"`
	phone    string `json:"phone"`
	role     string `json:"role"`
}

func NewUser(name, username, password, email, phone string) User {
	return &Customer{
		name:     name,
		userName: username,
		password: password,
		email:    email,
		phone:    phone,
		role:     RoleCustomer,
	}
}

func (c *Customer) GetName() string {
	return c.name
}

func (c *Customer) GetUserName() string {
	return c.userName
}

func (c *Customer) GetPassword() string {
	return c.password
}

func (c *Customer) GetEmail() string {
	return c.email
}

func (c *Customer) GetPhone() string {
	return c.phone
}

func (c *Customer) GetRole() string {
	return c.role
}
