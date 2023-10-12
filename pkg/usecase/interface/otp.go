package interfaces

import (
	"context"
	openapi "github.com/twilio/twilio-go/rest/verify/v2"
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
)

type OtpUseCase interface {
	SendOtp(ctx context.Context, phno helperstruct.OTPData) error
	ValidateOtp(otpDetails helperstruct.VerifyOtp) (*openapi.VerifyV2VerificationCheck, error)
}