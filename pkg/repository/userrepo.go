package repository

import (
	"context"
	"fmt"
	"log"
	"time"

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

// ---------------------- StoreOTP ---------------------- //

func (c *userDatabase) StoreOTP(userEmail string, otp string) bool {
	var userData domain.Users

	c.DB.Where("email=?", userEmail).First(&userData)

	OTP := domain.OTP{
		UserEmail: userData.Email,
		UserID:    int(userData.ID),
		Code:      otp,
		ExpiresAt: time.Now().Add(10 * time.Minute),
	}
	err := c.DB.Create(&OTP).Error
	if err != nil {
		return false
	}
	return err == nil
}

//--------------------- VerifyOTP ---------------------- //

func (c *userDatabase) VerifyOTP(otp string) (int, bool) {
	var otpDB domain.OTP

	// Use a parameterized query to prevent SQL injection.
	query := `SELECT * FROM otps WHERE code = $1`
	err := c.DB.Raw(query, otp).Scan(&otpDB).Error
	log.Println(otpDB.UserEmail)

	// Check for errors in the query execution.
	if err != nil {
		fmt.Printf(err.Error())
		// Handle the error, log it, or return an appropriate error message.
		// For example, you could return the error message:
		return 0, false
	}

	// Check if the OTP is expired.
	if otpDB.ExpiresAt.Before(time.Now()) {
		// Handle the expired OTP, log it, or return an appropriate error message.
		log.Printf("adasdas")
		return 0, false
	}

	// Delete the OTP from the database.
	deleteErr := c.DB.Delete(&otpDB).Error
	if deleteErr != nil {
		// Handle the error in the deletion process, log it, or return an appropriate error message.
		return 0, false
	}

	return otpDB.UserID, true
}

// func (c *userDatabase) OtpLogin(phno string) (int, error) {
// 	var id int
// 	query := "SELECT id FROM users WHERE mobile=?"
// 	err := c.DB.Raw(query, phno).Scan(&id).Error
// 	return id, err
// }
// func (c *userDatabase) insertOTP(emailID, otp string) error {
//     // Prepare the SQL statement for inserting OTP
// 	var id int
//     query := "INSERT INTO temp_otp (email_id, otp) VALUES ($1, $2)"

// 	err:=c.DB.Raw(query,emailID).Scan(&id).Error
//     // Execute the SQL query

//     return err
// }

func (c *userDatabase) AddAddress(id int, address helperstruct.Address) error {

	//Check if the new address is being set as default
	if address.IsDefault { //Change the default address into false
		changeDefault := `UPDATE addresses SET is_default = $1 WHERE users_id=$2 AND is_default=$3`
		err := c.DB.Exec(changeDefault, false, id, true).Error

		if err != nil {
			return err
		}
	}
	//Insert the new address
	insert := `INSERT INTO addresses (users_id,house_number,street,city, district,landmark,pincode,is_default)
		VALUES($1,$2,$3,$4,$5,$6,$7,$8)`
	err := c.DB.Exec(insert, id, address.
		House_number,
		address.Street,
		address.City,
		address.District,
		address.Landmark,
		address.Pincode,
		address.IsDefault).Error
	return err
}
func (c *userDatabase) UpdateAddress(id, addressId int, address helperstruct.Address) error {
	//Check if the new address is being set as default
	if address.IsDefault { //Change the default address into false
		changeDefault := `UPDATE addresses SET is_default = $1 WHERE id=$2 AND is_default=$3`
		err := c.DB.Exec(changeDefault, true, addressId, false).Error

		if err != nil {
			return err
		}
	}
	//Update the address
	update := `UPDATE addresses SET 
		house_number=$1,street=$2,city=$3, district=$4,landmark=$5,pincode=$6,is_default=$7 WHERE users_id=$8 AND id=$9`
	err := c.DB.Exec(update,
		address.House_number,
		address.Street,
		address.City,
		address.District,
		address.Landmark,
		address.Pincode,
		address.IsDefault,
		id,
		addressId).Error
	return err
}
func (c *userDatabase) ViewProfile(id int) (response.UserData, error) {
	var profileData response.UserData
	findProfile := `SELECT name,email,mobile FROM users WHERE id=?`
	err := c.DB.Raw(findProfile, id).Scan(&profileData).Error
	fmt.Printf("viewProfile")
	return profileData, err
}

func (c *userDatabase) UpdateProfile(id int,updatedetails helperstruct.UserReq) (response.UserData,error){
var profile response.UserData
query:=`UPDATE users SET name=$1,email=$2,mobile=$3 WHERE id=$4 RETURNING name,email,mobile`
err:=c.DB.Raw(query,updatedetails.Name,updatedetails.Email,updatedetails.Mobile,id).Scan(&profile).Error
fmt.Printf("updated Profile")
return profile,err
}
func (c *userDatabase) FindPassword(id int) (string, error) {
	var orginalPassword string
	err := c.DB.Raw("SELECT password FROM users WHERE id=?", id).Scan(&orginalPassword).Error
	return orginalPassword, err
}

func (c *userDatabase) UpdatePassword(id int, newPassword string) error {
	updatePassword := `UPDATE users SET password=$1 WHERE id=$2`
	err := c.DB.Exec(updatePassword, newPassword, id).Error
	return err
}
// func (c*userDatabase) Incrementwalllet(id int,money int)error{
// 	query:=`UPDATE users SET wallet = wallet+$1 WHERE id =$2`
// 	err:=c.DB.Exec(query,money,id).Error
// 	return err
// }
func (c *userDatabase) Decrementwallet(id int, money int) error {
	var checkmoney int
	checkwallet := `SELECT wallet FROM users WHERE id=$1`
	err := c.DB.Raw(checkwallet, id).Scan(&checkmoney).Error
	if err != nil {
		return fmt.Errorf("error in checking wallet amount: %v", err)
	}

	if checkmoney < money {
		return fmt.Errorf("insufficient funds in your wallet")
	}

	query := `UPDATE users SET wallet = wallet - $1 WHERE id=$2`
	err = c.DB.Exec(query, money, id).Error
	return err
}
