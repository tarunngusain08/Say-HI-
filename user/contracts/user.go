package contracts

type RegisterUser struct {
	Name     string `json:"name"`
	UserName string `json:"user_name"`
	Password string `json:"password"`
	Email    string `json:"email"`
}

type SendEmailRequest struct {
	Name  string `json:"name"`
	OTP   string `json:"otp"`
	Email string `json:"email"`
}

type VerifyEmailRequest struct {
	OTP   string `json:"otp"`
	Email string `json:"email"`
}
