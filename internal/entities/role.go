package entities

type Role struct {
	BaseEntityWithSoftDelete
	Name        string       `gorm:"type:varchar(255);unique"`
	Users       []User       `gorm:"many2many:user_roles"`
	Permissions []Permission `gorm:"many2many:role_permissions"`
}
