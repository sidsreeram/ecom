// Package usecase provides the business logic for user-related operations.
package usecase

import (
	"context"
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"github.com/ECOMMERCE_PROJECT/pkg/usecase/controller"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

// userUseCase implements the services.UserUseCase interface.
type userUseCase struct {
	userRepo interfaces.UserRepository
}

// NewUserUseCase creates a new instance of userUseCase.
func NewUserUseCase(repo interfaces.UserRepository) services.UserUseCase {
	return &userUseCase{
		userRepo: repo,
	}
}

// UserSignUp handles user registration by hashing the password and calling the repository.
func (c *userUseCase) UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error) {
	hash, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if err != nil {
		return response.UserData{}, err
	}
	user.Password = string(hash)
	userData, err := c.userRepo.UserSignUp(ctx, user)
	return userData, err
}

// UserLogin handles user login, checks credentials, generates OTP, and stores it.
func (c *userUseCase) UserLogin(ctx context.Context, user helperstruct.LoginReq) error {
	userData, err := c.userRepo.UserLogin(ctx, user.Email)
	if err != nil {
		return err
	}

	if user.Email == "" {
		return fmt.Errorf("Please provide valid EMAILID")
	}

	err = bcrypt.CompareHashAndPassword([]byte(userData.Password), []byte(user.Password))
	if err != nil {
		return fmt.Errorf("Password entered is incorrect")
	}

	if userData.IsBlocked {
		return fmt.Errorf("user is blocked")
	}

	otp := controller.GenerateOTP()
	controller.SendOTP(user, otp)
	log.Println(otp)
	err = userStoreOTP(c, user.Email, otp)
	if err != nil {
		return err
	}

	return nil
}

// VerifyOTP verifies the provided OTP, generates a JWT token upon success.
func (c *userUseCase) VerifyOTP(otp string) (string, error) {
	id, res := c.userRepo.VerifyOTP(otp)

	if !res {
		return "", errors.New("error in verifying otp")
	}

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	return ss, nil
}


// userStoreOTP stores the OTP in the repository.
func userStoreOTP(c *userUseCase, userEmail string, otp string) error {
	store := c.userRepo.StoreOTP(userEmail, otp)
	if store {
		return nil
	} else {
		return errors.New("error in storing otp")
	}
}

// AddAddress adds an address to the user's profile.
func (c *userUseCase) AddAddress(id int, address helperstruct.Address) error {
	err := c.userRepo.AddAddress(id, address)
	return err
}

// UpdateAddress updates an existing address in the user's profile.
func (c *userUseCase) UpdateAddress(id, addressId int, address helperstruct.Address) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}

// ViewProfile retrieves the user's profile information.
func (c *userUseCase) ViewProfile(id int) (response.UserData, error) {
	response, err := c.userRepo.ViewProfile(id)
	return response, err
}

// UpdateProfile updates the user's profile information.
func (c *userUseCase) UpdateProfile(id int, updatingdetails helperstruct.UserReq) (response.UserData, error) {
	updatedProfile, err := c.userRepo.UpdateProfile(id, updatingdetails)
	return updatedProfile, err
}

// ChangePassword initiates the process of changing the user's password by generating and storing a new OTP.
func (c *userUseCase) ChangePassword(user helperstruct.Email) error {
	if user.Email == "" {
		return fmt.Errorf("Please provide valid EMAILID")
	}

	otp := controller.GenerateOTP()
	controller.SendOTPforpassword(user, otp)
	log.Println("password", otp)
	err := userStoreOTP(c, user.Email, otp)
	if err != nil {
		return err
	}

	return nil
}

// VerfiyForChangePassword verifies the OTP and updates the password if successful.
func (c *userUseCase) VerfiyForChangePassword(otp string, id int, passwords helperstruct.UpdatePassword) error {
	id, res := c.userRepo.VerifyOTP(otp)
	if !res {
		return errors.New("error in verifying otp")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(passwords.NewPassword), 10)
	if err != nil {
		return err
	}
	newPassword := string(hash)

	err = c.userRepo.UpdatePassword(id, newPassword)
	return err
}
