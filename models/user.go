package models

type User struct {
	Id       uint   `gorm:"primaryKey"`
	Name     string `gorm:"type:varchar(100)"`
	Email    string `gorm:"type:varchar(100);unique"`
	Password []byte `gorm:"type:varchar(100)"`
}
