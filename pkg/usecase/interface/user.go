package interfaces

import (
	"context"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error)
	UserLogin(ctx context.Context, user helperstruct.LoginReq) error
	IsSignIn(phno string) (bool, error)
	VerifyOTP(otp string) (string, error)
	AddAddress(id int, address helperstruct.Address) error
	UpdateAddress(id, addressId int, address helperstruct.Address) error

}
