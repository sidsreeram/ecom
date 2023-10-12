package interfaces

import (
	"context"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type UserRepository interface {
	UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error) 
	UserLogin(ctx context.Context,email string)(domain.Users,error)
	IsSignIn(phno string) (bool, error)
	OtpLogin(phno string) (int, error)
}