package helperstruct

type UserReq struct {
	Name     string `json:"name" binding:"required"`
	Email    string `json:"email" binding:"required,email"`
	Mobile   string `json:"mobile" binding:"required"`
	Password string `json:"password"`
}

type LoginReq struct {
	Email    string `json:"email"  binding:"required"`
	Password string `json:"password" binding:"required"`
}

type OTP struct {
	Code string
}
type Address struct {
	House_number string `json:"house_number" `
	Street       string `json:"street" `
	City         string `json:"city " `
	District     string `json:"district " `
	Landmark     string `json:"landmark" `
	Pincode      int    `json:"pincode " `
	IsDefault    bool   `json:"isdefault" `
}
type UpdatePassword struct {
    OldPassword string `json:"oldpassword" `
    NewPassword string `json:"newpassword" `
}
type Email struct {
	Email string `json:"email"`
}