package interfaces

import (
	"context"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type UserRepository interface {
	UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error)
	UserLogin(ctx context.Context, email string) (domain.Users, error)
	IsSignIn(phno string) (bool, error)
	StoreOTP(userEmail string, otp string) bool
	VerifyOTP(otp string) (int, bool)
	AddAddress(id int, address helperstruct.Address) error
	UpdateAddress(id, addressId int, address helperstruct.Address) error
	ViewProfile(id int) (response.UserData, error)
	UpdateProfile(id int,updatedetails helperstruct.UserReq) (response.UserData,error)
	FindPassword(id int) (string, error)
	UpdatePassword(id int, newPassword string) error
//     Incrementwalllet(id int,money int)error
}
