package requests

type RefreshTokenReq struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
	IpAddress    string
	UserAgent    string
}
