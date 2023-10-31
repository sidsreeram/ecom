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

type userUseCase struct {
	userRepo interfaces.UserRepository
}

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

func (c *userUseCase) VerifyOTP(otp string) (string, error) {
	id, res := c.userRepo.VerifyOTP(otp)
	log.Printf("absdhasd")
	if !res {
		return "", errors.New("error in verifying otp")
	}
	log.Println(id)

	claims := jwt.MapClaims{
		"id":  id,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}
	fmt.Println(ss)
	return ss, nil

}

func (c *userUseCase) IsSignIn(phno string) (bool, error) {
	isSignin, err := c.userRepo.IsSignIn(phno)
	return isSignin, err
}

func userStoreOTP(c *userUseCase, userEmail string, otp string) error {
	store := c.userRepo.StoreOTP(userEmail, otp)
	if store {
		return nil
	} else {
		return errors.New("error in storing otp")
	}
}
func (c *userUseCase) AddAddress(id int, address helperstruct.Address) error {
	err := c.userRepo.AddAddress(id, address)
	return err
}
func (c *userUseCase) UpdateAddress(id, addressId int, address helperstruct.Address) error {
	err := c.userRepo.UpdateAddress(id, addressId, address)
	return err
}
func (c *userUseCase) ViewProfile(id int) (response.UserData, error) {
	response, err := c.userRepo.ViewProfile(id)
	return response, err
}
func (c *userUseCase) UpdateProfile(id int, updatingdetails helperstruct.UserReq) (response.UserData, error) {
	updatedProfile, err := c.userRepo.UpdateProfile(id, updatingdetails)
	return updatedProfile, err
}
func (c *userUseCase) ChangePassword( user helperstruct.Email) error {

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
func (c *userUseCase) VerfiyForChangePassword(otp string, id int, passwords helperstruct.UpdatePassword) error {
	id, res := c.userRepo.VerifyOTP(otp)
	if !res {
		return errors.New("error in verifying otp")
	}
	// orginalPassword, err := c.userRepo.FindPassword(id)
	// if err != nil {
	// 	return err
	// }

	// // err = bcrypt.CompareHashAndPassword([]byte(orginalPassword), []byte(passwords.OldPassword))
	// // if err != nil {
	// // 	return err
	// // }

	hash, err := bcrypt.GenerateFromPassword([]byte(passwords.NewPassword), 10)
	if err != nil {
		return err
	}
	newPassword := string(hash)

	err = c.userRepo.UpdatePassword(id, newPassword)
	return err
}
