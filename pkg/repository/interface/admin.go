package interfaces

import (
	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
)

type AdminRepository interface {
	IsSuperAdmin(createrId int) (bool, error)
	CreateAdmin(admin helperstruct.CreateAdmin) (response.AdminData, error)
	AdminLogin(email string) (domain.Admins, error)
	BlockUser(body helperstruct.BlockData, adminId int) error
	UnblockUser(id int) error
}
