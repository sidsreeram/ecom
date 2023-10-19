package interfaces

import (
	"context"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
)

type AdminUsecase interface {
	CreateAdmin(ctx context.Context, admis helperstruct.CreateAdmin, createrId int) (response.AdminData, error)
	AdminLogin(admin helperstruct.LoginReq) (string, error)
	BlockUser(body helperstruct.BlockData, adminId int) error
	UnblockUser(id int) error
	FindUser(id int) (response.UserDetails, error)
	ListAllUsers() ([]response.UserDetails, error) 
}
