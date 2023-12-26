package interfaces

import (
	"context"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
)

type UserUseCase interface {
	UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error)
	UserLogin(ctx context.Context, user helperstruct.LoginReq) error
	VerifyOTP(otp string) (string, error)
	AddAddress(id int, address helperstruct.Address) error
	UpdateAddress(id, addressId int, address helperstruct.Address) error
	ViewProfile(id int)(response.UserData, error)
	UpdateProfile(id int,updatingdetails helperstruct.UserReq)(response.UserData,error)
	ChangePassword(user helperstruct.Email) error
	VerfiyForChangePassword(otp string, id int, passwords helperstruct.UpdatePassword) error
}
