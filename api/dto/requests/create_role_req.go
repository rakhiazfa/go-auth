package requests

type CreateRoleReq struct {
	Name string `json:"name" validate:"required,max=255"`
}
