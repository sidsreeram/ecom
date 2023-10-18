package usecase

import (
	"fmt"
	"math/rand"
	"net/smtp"
	"strconv"
	"sync"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
)

// type OtpUseCase struct {
// 	// otpRepo interfaces.OtpRepository
// 	cfg config.Config
// }

// func NewOtpUseCase(cfg config.Config) services.OtpUseCase {
// 	return &OtpUseCase{
// 		// otpRepo: repo,
// 		cfg: cfg,
// 	}
// }

// func (c *OtpUseCase) SendOtp(ctx context.Context, phno helperstruct.OTPData) error {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	twilioAccountSID := os.Getenv("TWILIOAUTHTOKEN")
// 	twilioAuthToken := os.Getenv("TWILIOACCOUNTSID")
// 	twilioServiceID := os.Getenv("TWILIOSERVICEID")
// 	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Username: twilioAccountSID,
// 		Password: twilioAuthToken,
// 	})

// 	params := &openapi.CreateVerificationParams{}
// 	params.SetTo("+91" + phno.PhoneNumber)
// 	log.Fatal(twilioAccountSID)
// 	params.SetChannel("sms")
// 	_, err = client.VerifyV2.CreateVerification(twilioServiceID, params)
// 	return err
// }

// func (c *OtpUseCase) ValidateOtp(otpDetails helperstruct.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error) {
// 	err := godotenv.Load(".env")
// 	if err != nil {
// 		log.Fatal("Error loading .env file")
// 	}
// 	twilioAccountSID := os.Getenv("TWILIOAUTHTOKEN")
// 	twilioAuthToken := os.Getenv("TWILIOACCOUNTSID")
// 	twilioServiceID := os.Getenv("TWILIOSERVICEID")
// 	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
// 		Username: twilioAccountSID,
// 		Password: twilioAuthToken,
// 	})
// 	params := &openapi.CreateVerificationCheckParams{}
// 	params.SetTo("+91" + otpDetails.User.PhoneNumber)
// 	params.SetCode(otpDetails.Code)
// 	resp, err := client.VerifyV2.CreateVerificationCheck(twilioServiceID, params)
// 	return resp, err
// }

var otpMutex sync.Mutex

// GenerateOTPEMAIL generates an OTP for email.
func GenerateOTPEMAIL() string {
	otpMutex.Lock()
	defer otpMutex.Unlock()

	// Generate a random 4-digit OTP
	otp := strconv.Itoa(rand.Intn(9000) + 1000)
	return otp
}
func SendEmailOTP(email helperstruct.UserReq, otp string) error {
	auth := smtp.PlainAuth(
		"",
		"sidx141202@gmail.com", // Replace with your Gmail email
		"yhlm wzqg wobt deww",  // Replace with your Gmail app password
		"smtp.gmail.com",
	)
	subject := "OTP Verification"
	body := "Your OTP code is: " + otp
	msg := "Subject: " + subject + "\r\n\r\n" + body

	err := smtp.SendMail(
		"smtp.gmail.com:587",
		auth,
		"sidx141202@gmail.com",     // Replace with your Gmail email
		[]string{email.Email}, // Use email from the struct
		[]byte(msg),
	)

	if err != nil {
		fmt.Println("Failed to send email:", err)
	}

	return err
}
