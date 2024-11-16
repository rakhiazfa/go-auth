package requests

type SignInReq struct {
	Username  string `json:"username" validate:"required"`
	Password  string `json:"password" validate:"required"`
	IpAddress string
	UserAgent string
}
