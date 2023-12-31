package types

import user "Say-Hi/user/handler"

type Handler struct {
	User
	Notification
	Message
	ChatHistory
}

type User struct {
	RegisterHandler user.RegisterHandler
}

type Notification struct {
}

type Message struct {
}

type ChatHistory struct {
}
