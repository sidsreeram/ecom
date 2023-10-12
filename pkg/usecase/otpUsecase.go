package usecase

import (
	"context"
	"log"
	"os"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/config"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/joho/godotenv"
	"github.com/twilio/twilio-go"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
)

type OtpUseCase struct {
	// otpRepo interfaces.OtpRepository
	cfg config.Config
}

func NewOtpUseCase(cfg config.Config) services.OtpUseCase {
	return &OtpUseCase{
		// otpRepo: repo,
		cfg: cfg,
	}
}

func (c *OtpUseCase) SendOtp(ctx context.Context, phno helperstruct.OTPData) error {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	twilioAccountSID := "AC1ca3125a8ea1cb7a13e80550f2acd880"
	twilioAuthToken := "f0a09d942e63796bb035b2b8c535d922"
	twilioServiceID := "VAa53821b64c22cfe9140d5bd789a7f5ea"
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioAccountSID,
		Password: twilioAuthToken,
	})

	params := &openapi.CreateVerificationParams{}
	params.SetTo("+91" + phno.PhoneNumber)
	log.Fatal("hiiiii")
	log.Fatal(twilioAccountSID)
	params.SetChannel("sms")
	_, err = client.VerifyV2.CreateVerification(twilioServiceID, params)
	return err
}

func (c *OtpUseCase) ValidateOtp(otpDetails helperstruct.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error) {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	twilioAccountSID := os.Getenv("TWILIOACCOUNTSID")
	twilioAuthToken := os.Getenv("TWILIOAUTHTOKEN")
	twilioServiceID := os.Getenv("TWILIOSERVICEID")
	var client *twilio.RestClient = twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: twilioAccountSID,
		Password: twilioAuthToken,
	})
	params := &openapi.CreateVerificationCheckParams{}
	params.SetTo("+91" + otpDetails.User.PhoneNumber)
	params.SetCode(otpDetails.Code)
	resp, err := client.VerifyV2.CreateVerificationCheck(twilioServiceID, params)
	return resp, err
}
