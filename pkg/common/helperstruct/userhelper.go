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

type OTPData struct {
	PhoneNumber string `json:"phnum,omitempty" validate:"required"`
}

type VerifyOtp struct {
	User *OTPData `json:"user,omitempty" validate:"required"`
	Code string   `json:"code,omitempty" validate:"required"`
}
type EMAILOTP struct{
	EmailId string `json:"email" binding:"required"`
}