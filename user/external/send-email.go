package external

import (
	"Say-Hi/user/constants"
	"Say-Hi/user/contracts"
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"net/http"
	"time"
)

type EmailService struct {
	MaxRetries int
	BaseDelay  time.Duration
	MaxDelay   time.Duration
}

func NewEmailService(maxRetries int, baseDelay, maxDelay time.Duration) *EmailService {
	return &EmailService{
		MaxDelay:   maxDelay,
		MaxRetries: maxRetries,
		BaseDelay:  baseDelay,
	}
}

func generateOTP() string {
	otp := rand.Intn(900000) + 100000
	return fmt.Sprintf("%06d", otp)
}

func createHttpRequest(url, method, action string, body io.Reader) (*http.Request, error) {
	req, err := http.NewRequest(method, url, body)
	if err != nil {
		return nil, err
	}
	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("action", action)
	return req, nil
}

func (e *EmailService) SendEmailWithExponentialBackoff(email, name string) (*string, error) {

	for attempt := 0; attempt < e.MaxRetries; attempt++ {
		if otp, err := e.send(email, name); err == nil {
			// Email sent successfully
			return otp, nil
		}

		fmt.Printf("Attempt %d failed: \n", attempt+1)

		// Calculate backoff duration using exponential backoff formula
		delay := time.Duration(e.BaseDelay*(1<<attempt)) * time.Second
		if delay > e.MaxDelay*time.Second {
			delay = e.MaxDelay * time.Second
		}

		fmt.Printf("Retrying in %v...\n", delay)
		time.Sleep(delay)
	}

	return nil, fmt.Errorf("failed after %d attempts", e.MaxRetries)
}

func (e *EmailService) send(email, name string) (*string, error) {
	req := contracts.SendEmailRequest{
		OTP:   generateOTP(),
		Email: email,
		Name:  name,
	}

	body, err := json.Marshal(req)
	if err != nil {
		return nil, err
	}

	request, err := createHttpRequest(constants.SendEmailUrl, http.MethodPost, constants.VerifyEmailAction, bytes.NewBuffer(body))
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK && response.StatusCode != http.StatusCreated {
		return nil, errors.New("invalid email")
	}

	return &req.OTP, nil
}
