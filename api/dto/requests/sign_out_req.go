package requests

type SignOutReq struct {
	RefreshToken string `json:"refreshToken" validate:"required"`
}
