package user

type User interface {
	GetName() string
	GetUserName() string
	GetPassword() string
	GetEmail() string
	GetPhone() string
	GetRole() string
}
