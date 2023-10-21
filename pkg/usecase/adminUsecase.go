package usecase

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	services "github.com/ECOMMERCE_PROJECT/pkg/usecase/interface"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
)

type adminUseCase struct {
	adminRepo interfaces.AdminRepository
	
}

func NewAdminUsecase(adminRepo interfaces.AdminRepository) services.AdminUsecase {
	return &adminUseCase{
		adminRepo: adminRepo,
		
	}
}

func (c *adminUseCase) CreateAdmin(ctx context.Context, admin helperstruct.CreateAdmin, createrId int) (response.AdminData, error) {
	IsSuper, err := c.adminRepo.IsSuperAdmin(createrId)
	if err != nil {
		return response.AdminData{}, err
	}
	if !IsSuper {
		return response.AdminData{}, fmt.Errorf("not a super admin")
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(admin.Password), 10)
	if err != nil {
		return response.AdminData{}, err
	}
	admin.Password = string(hash)
	adminData, err := c.adminRepo.CreateAdmin(admin)

	return adminData, err
}

func (c *adminUseCase) AdminLogin(admin helperstruct.LoginReq) (string, error) {
	adminData, err := c.adminRepo.AdminLogin(admin.Email)
	if err != nil {
		return "", err
	}

	if adminData.Email == "" {
		return "", fmt.Errorf("no user found")
	}

	err = bcrypt.CompareHashAndPassword([]byte(adminData.Password), []byte(admin.Password))
	if err != nil {
		return "", err
	}

	if adminData.IsBlocked {
		return "", fmt.Errorf("user is blocked")
	}

	claims := jwt.MapClaims{
		"id":  adminData.ID,
		"exp": time.Now().Add(time.Hour * 72).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	ss, err := token.SignedString([]byte("secret"))
	if err != nil {
		return "", err
	}

	return ss, nil
}

func (c *adminUseCase) BlockUser(body helperstruct.BlockData, adminId int) error {
	err := c.adminRepo.BlockUser(body, adminId)
	return err
}

func (c *adminUseCase) UnblockUser(id int) error {
	err := c.adminRepo.UnblockUser(id)
	return err
}
func (c *adminUseCase) FindUser(id int) (response.UserDetails, error) {
	userDetails, err := c.adminRepo.FindUser(id)
	return userDetails, err
}

func (c *adminUseCase) ListAllUsers() ([]response.UserDetails, error) {
	var userList []response.UserDetails

	// You need to fetch a list of user IDs here. Replace this with your actual logic.
	userIDs, err := c.adminRepo.ListAllUsers()
	if err != nil {
		return []response.UserDetails{}, err
	}

	for _, userID := range userIDs {
		userDetails, err := c.FindUser(userID.UserID)
		if err != nil {
			log.Printf("Can't find user")
			continue
		}
		userList = append(userList, userDetails)
	}

	return userList, nil
}

