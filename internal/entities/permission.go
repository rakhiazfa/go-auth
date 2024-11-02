package entities

type Permission struct {
	BaseEntityWithSoftDelete
	Name       string `gorm:"type:varchar(255)"`
	ServiceKey string `gorm:"type:varchar(255)"`
	Method     string `gorm:"type:varchar(50)"`
	Path       string `gorm:"type:varchar(255)"`
	Roles      []Role `gorm:"many2many:role_permissions"`
}
