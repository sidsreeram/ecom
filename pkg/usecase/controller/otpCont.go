package controller

import (
	"math/rand"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"gopkg.in/gomail.v2"
)

func GenerateOTP() string {
	b := make([]byte, 6)
	_, err := rand.Read(b)
	if err != nil {
		panic(err)
	}

	for i := range b {
		b[i] = byte(48 + rand.Intn(10))
	}

	return string(b)
}

func SendOTP(email helperstruct.LoginReq, otp string) error {
	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", "sidx141202@gmail.com")
	message.SetHeader("To", email.Email)
	message.SetHeader("Subject", "OTP Verification")

	// Set the message body
	message.SetBody("text/plain", "This will be expired in 10 minutes\nYour OTP is: "+otp)

	// Create an instance of the SMTP sender.
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "sidx141202@gmail.com", "yhlm wzqg wobt deww")

	// Send the email message.
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}

// var otpMutex sync.Mutex

// // GenerateOTPEMAIL generates an OTP for email.
// func GenerateOTPEMAIL() string {
// 	otpMutex.Lock()
// 	defer otpMutex.Unlock()

// 	// Generate a random 4-digit OTP
// 	otp := strconv.Itoa(rand.Intn(9000) + 1000)
// 	return otp
// }
// func SendEmailOTP(email helperstruct.LoginReq, otp string) error {
// 	auth := smtp.PlainAuth(
// 		"",
// 		"sidx141202@gmail.com", // Replace with your Gmail email
// 		"yhlm wzqg wobt deww",  // Replace with your Gmail app password
// 		"smtp.gmail.com",
// 	)
// 	subject := "OTP Verification"
// 	body := "Your OTP code is: " + otp
// 	msg := "Subject: " + subject + "\r\n\r\n" + body

// 	err := smtp.SendMail(
// 		"smtp.gmail.com:587",
// 		auth,
// 		"sidx141202@gmail.com", // Replace with your Gmail email
// 		[]string{email.Email},  // Use email from the struct
// 		[]byte(msg),
// 	)

// 	if err != nil {
// 		fmt.Println("Failed to send email:", err)
// 	}

// 	return err
// }
func SendOTPforpassword(email helperstruct.Email, otp string) error {
	// Create a new message.
	message := gomail.NewMessage()
	message.SetHeader("From", "sidx141202@gmail.com")
	message.SetHeader("To", email.Email)
	message.SetHeader("Subject", "OTP Verification")

	// Set the message body
	message.SetBody("text/plain", "This will be expired in 10 minutes\nYour OTP is: "+otp)

	// Create an instance of the SMTP sender.
	dialer := gomail.NewDialer("smtp.gmail.com", 587, "sidx141202@gmail.com", "yhlm wzqg wobt deww")

	// Send the email message.
	if err := dialer.DialAndSend(message); err != nil {
		return err
	}
	return nil
}
