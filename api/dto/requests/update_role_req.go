package requests

type UpdateRoleReq struct {
	Name string `json:"name" validate:"required,max=255"`
}
