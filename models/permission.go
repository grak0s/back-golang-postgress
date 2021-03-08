package models

type Permission struct {
	Id   uint   `json:"id" gorm:"primaryKey"`
	Name string `json:"name"`
}
