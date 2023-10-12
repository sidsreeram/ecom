package repository

import (
	"fmt"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)

type adminDatabase struct {
	DB *gorm.DB
}

func NewAdminRepository(DB *gorm.DB) interfaces.AdminRepository {
	return &adminDatabase{DB}
}

func (c *adminDatabase) IsSuperAdmin(createrId int) (bool, error) {
	var isSuper bool
	query := "SELECT is_super_admin FROM admins WHERE id=$1"
	err := c.DB.Raw(query, createrId).Scan(&isSuper).Error
	return isSuper, err
}

func (c *adminDatabase) CreateAdmin(admin helperstruct.CreateAdmin) (response.AdminData, error) {
	var adminData response.AdminData
	query := `INSERT INTO admins (name,email,password,is_super_admin,created_at)
								  VALUES($1,$2,$3,$4,NOW())
								  RETURNING id,name,email,is_super_admin`

	err := c.DB.Raw(query, admin.Name, admin.Email, admin.Password, admin.IsSuper).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) AdminLogin(email string) (domain.Admins, error) {
	var adminData domain.Admins
	err := c.DB.Raw("SELECT * FROM admins WHERE email=?", email).Scan(&adminData).Error
	return adminData, err
}

func (c *adminDatabase) BlockUser(body helperstruct.BlockData, adminId int) error {
	
	tx := c.DB.Begin()

	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1)", body.UserId).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user")
	}

	if err := tx.Exec("UPDATE users SET is_blocked = true WHERE id = ?", body.UserId).Error; err != nil {
		tx.Rollback()
		return err
	}
	
	if err := tx.Exec("INSERT INTO user_infos (users_id, reason_for_blocking, blocked_at, blocked_by) VALUES (?, ?, NOW(), ?)", body.UserId, body.Reason, adminId).Error; err != nil {
		tx.Rollback()
		return err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	
	return nil

}

func (c *adminDatabase) UnblockUser(id int) error {
	tx := c.DB.Begin()

	var isExists bool
	if err := tx.Raw("SELECT EXISTS(SELECT 1 FROM users WHERE id = $1 AND is_blocked=true)", id).Scan(&isExists).Error; err != nil {
		tx.Rollback()
		return err
	}
	if !isExists {
		tx.Rollback()
		return fmt.Errorf("no such user to unblock")
	}
	if err := tx.Exec("UPDATE users SET is_blocked = false WHERE id=$1", id).Error; err != nil {
		tx.Rollback()
		return err
	}
	query := "UPDATE user_infos SET reason_for_blocking=$1,blocked_at=NULL,blocked_by=$2 WHERE users_id=$3"
	if err := tx.Exec(query, "", 0, id).Error; err != nil {
		tx.Rollback()
		return err
	}
	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return err
	}
	return nil
}
