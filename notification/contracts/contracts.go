package contracts

type SendEmailRequest struct {
	Name  string `json:"name"`
	OTP   string `json:"otp"`
	Email string `json:"email"`
}
