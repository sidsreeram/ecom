package repository

import (
	"context"

	"github.com/ECOMMERCE_PROJECT/pkg/common/helperstruct"
	"github.com/ECOMMERCE_PROJECT/pkg/common/response"
	"github.com/ECOMMERCE_PROJECT/pkg/domain"
	interfaces "github.com/ECOMMERCE_PROJECT/pkg/repository/interface"
	"gorm.io/gorm"
)
type userDatabase struct {
	DB *gorm.DB
}

func NewUserRepository(DB *gorm.DB) interfaces.UserRepository {
	return &userDatabase{DB}
}

func (c *userDatabase) UserSignUp(ctx context.Context, user helperstruct.UserReq) (response.UserData, error) {
	var userData response.UserData
	insertQuery := `INSERT INTO users (name,email,mobile,password)VALUES($1,$2,$3,$4) 
					RETURNING id,name,email,mobile`
	err := c.DB.Raw(insertQuery, user.Name, user.Email, user.Mobile, user.Password).Scan(&userData).Error
	return userData, err
}
func (c *userDatabase) UserLogin(ctx context.Context, email string) (domain.Users, error) {
	var userData domain.Users
	err := c.DB.Raw("SELECT * FROM users WHERE email=?", email).Scan(&userData).Error
	return userData, err
}

func (c *userDatabase) IsSignIn(phno string) (bool, error) {
	quwery := "select exists(select 1 from users where mobile=?)"
	var isSignIng bool
	err := c.DB.Raw(quwery, phno).Scan(&isSignIng).Error
	return isSignIng, err

}

func (c *userDatabase) OtpLogin(phno string) (int, error) {
	var id int
	query := "SELECT id FROM users WHERE mobile=?"
	err := c.DB.Raw(query, phno).Scan(&id).Error
	return id, err
}