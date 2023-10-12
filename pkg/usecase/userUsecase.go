package usecase

import (
	"context"
	"fmt"
	"time"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type userUseCase struct {
	userRepo interfaces.UserRepository
}

// IsSignIn implements interfaces.UserUseCase.

func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

func (c *userUseCase) UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(ctx, user)
	return userData, err
}

func (c *userUseCase) UserLogin(ctx context.Context, user helperstruct.LoginReq) (string, error) {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	if err != nil {
		return "", err
	}

	if user.Email == "" {
		return "", fmt.Errorf("Please provide valid EMAILID")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return "", fmt.Errorf("Password entered is incorrect")
	}

	if userData.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}

	claims := jwt.MapClaims{
		"id":  userData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *userUseCase) IsSignIn(phno string) (bool, error) {
	isSignin, err := c.userRepo.IsSignIn(phno)
	return isSignin, err
}

func (c *userUseCase) OtpLogin(phno string) (string, error) {
	id, err := c.userRepo.OtpLogin(phno)
	if err != nil {
		return "", err
	}

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}