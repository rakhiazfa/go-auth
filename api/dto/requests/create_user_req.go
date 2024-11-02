package requests

type CreateUserReq struct {
	Avatar   *string `json:"avatar"`
	Name     string  `json:"name" validate:"required,max=255"`
	Username string  `json:"username" validate:"required,username,max=255"`
	Email    string  `json:"email" validate:"required,email,max=255"`
	Password string  `json:"password" validate:"required,min=8,max=255"`
}
