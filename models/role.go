package models

type Role struct {
	Id         uint         `json:"id"gorm:"primaryKey";"unique"`
	Name       string       `json:"name"`
	Permission []Permission `json:"permissions" gorm:"many2many:role_permissions"`
}
