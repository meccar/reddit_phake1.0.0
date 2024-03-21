package mail

import (
	"fmt"
	"math/rand"
	"time"

	util "util"

	"github.com/rs/zerolog/log"
)

// EmailVerifier represents an email verification service.
type EmailVerifier struct {
}

type mailer interface {
	SendMail(subject, body string, to []string) error
}

func StartMail(receiver, secret_code string) {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Err(err).Msg("cannot load config")
	}
	mailer := NewGmailSender(config.EmailSenderName, config.EmailSenderAddress, config.EmailSenderPassword)
	Send(mailer.(*GmailSender),receiver, secret_code)

	// return mailer

}

// NewEmailVerifier creates a new instance of EmailVerifier.
func NewEmailVerifier() *EmailVerifier {
	return &EmailVerifier{}
}

// GenerateCode generates a random verification code.
func (ev *EmailVerifier) GenerateCode(length int) string {
	rand.Seed(time.Now().UnixNano())

	// Available characters for the verification code
	charset := "ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789"

	code := make([]byte, length)
	for i := 0; i < length; i++ {
		code[i] = charset[rand.Intn(len(charset))]
	}
	return string(code)
}

// SendVerificationEmail sends a verification email to the given email address.
func (ev *EmailVerifier) SendVerificationEmail(email, code string) error {
	// Simulate sending email
	fmt.Printf("Verification code %s has been sent to %s\n", code, email)
	return nil
}

// VerifyCode verifies if the provided code matches the expected verification code.
func (ev *EmailVerifier) VerifyCode(expectedCode, actualCode string) bool {
	return expectedCode == actualCode
}

func Send(mailer *GmailSender, receiver, secret_code string) {

	// Define email details
	subject := "Welcome to TAF Viet"
	// content := "This is a test email content."
	verifyUrl := fmt.Sprintf("http://ltr5lt-1212.csb.app/api/v1/verify_email?email=%s&secret_code=%s",
	receiver, secret_code)
	content := fmt.Sprintf(`
	Thank you for registering with us!<br/>
	Please <a href="%s">click here</a> to verify your email address.<br/>
	`, verifyUrl)
	// The verification code is %s.<br/>

	// testv1.0@yahoo.com
	// test.v1.0.0.0.0.0.0.0.0@gmail.com
	to := []string{receiver}

	// Send the email
	err := mailer.SendEmail(subject, content, to, nil, nil, nil)
	if err != nil {
		log.Error().Err(err).Msg("failed to send email")
	} else {
		log.Info().Msg("email sent successfully")
	}

}
