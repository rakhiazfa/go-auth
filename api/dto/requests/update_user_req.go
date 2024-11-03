package requests

type UpdateUserReq struct {
	Name     string `json:"name" validate:"required,max=255"`
	Username string `json:"username" validate:"required,username,max=255"`
	Email    string `json:"email" validate:"required,email,max=255"`
}
