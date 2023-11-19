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
func (c *adminDatabase) FindUser(id int) (response.UserDetails, error) {
	var userdetails response.UserDetails
	query := `SELECT  users.name, users.email , users.mobile , users.is_blocked , info.bloked_by , info.blocked_at , info.reason__for_blocking FROM users as users FULL OUTER JOIN user_info as info on user.id = info.user_id `
	err := c.DB.Raw(query, id).Scan(&userdetails).Error
	if err != nil {
		return response.UserDetails{}, err
	}
	if userdetails.Email == "" {
		return response.UserDetails{}, fmt.Errorf("There is no such user")
	}
	return response.UserDetails{}, nil
}
func (c *adminDatabase) ListAllUsers() ([]response.UserDetails, error) {
	var userdetails []response.UserDetails

	query := `
		SELECT users.name, users.email, users.mobile, users.is_blocked, info.blocked_by, info.blocked_at, info.reason_for_blocking
		FROM users
		LEFT JOIN user_info as info ON users.id = info.user_id
	`

	err := c.DB.Raw(query).Scan(&userdetails).Error
	if err != nil {
		return []response.UserDetails{}, err
	}

	return userdetails, nil
}
func (c *adminDatabase) GetDashBoard() (response.DashBoard, error) {
	tx := c.DB.Begin()
	var dashBoard response.DashBoard
	getDasheBoard := `SELECT SUM(oi.quantity*oi.price)as Total_Revenue,
			SUM (oi.quantity)as Total_Products_Selled,
			COUNT(DISTINCT o.id)as Total_Orders FROM orders o
			JOIN order_items oi on o.id=oi.orders_id
			WHERE o.order_status_id=$1`
	if err := tx.Raw(getDasheBoard, 1).Scan(&dashBoard).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}

	getTotalUsers := `SELECT COUNT(id)AS TotalUsers FROM users`
	if err := tx.Raw(getTotalUsers).Scan(&dashBoard.TotalUsers).Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}

	if err := tx.Commit().Error; err != nil {
		tx.Rollback()
		return response.DashBoard{}, err
	}
	return dashBoard, nil
}
func (c *adminDatabase) ViewDailySalesReport() ([]response.SalesReport, error) {
	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id 
		WHERE o.order_status_id=1 AND DATE(o.order_date) = CURRENT_DATE`
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}

func (c *adminDatabase) ViewWeeklySalesReport() ([]response.SalesReport, error) {
	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id 
		WHERE o.order_status_id=1 AND EXTRACT(WEEK FROM o.order_date) = EXTRACT(WEEK FROM CURRENT_DATE) AND EXTRACT(YEAR FROM o.order_date) = EXTRACT(YEAR FROM CURRENT_DATE)`
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}

func (c *adminDatabase) ViewMonthlySalesReport() ([]response.SalesReport, error) {
	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id 
		WHERE o.order_status_id=1 AND EXTRACT(MONTH FROM o.order_date) = EXTRACT(MONTH FROM CURRENT_DATE) AND EXTRACT(YEAR FROM o.order_date) = EXTRACT(YEAR FROM CURRENT_DATE)`
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}


func (c *adminDatabase) ViewYearlySalesReport() ([]response.SalesReport, error) {
	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id 
		WHERE o.order_status_id=1 AND EXTRACT(YEAR FROM o.order_date) = EXTRACT(YEAR FROM CURRENT_DATE)`
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}



func (c *adminDatabase) ViewSalesReport() ([]response.SalesReport, error) {
	var sales []response.SalesReport
	getReports := `SELECT u.name,
		pt.type AS payment_type,
		o.order_date,
		o.order_total 
		FROM orders o JOIN users u ON u.id=o.user_id 
		JOIN payment_types pt ON o.payment_type_id= pt.id 
		WHERE o.order_status_id=1`
	err := c.DB.Raw(getReports).Scan(&sales).Error
	return sales, err
}
