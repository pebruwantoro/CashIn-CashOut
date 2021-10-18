package models

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID           uint   `gorm:"type:int; primaryKey"`
	Username     string `gorm:"type:varchar(16); unique; not null" json:"username" form:"username"`
	First_Name   string `gorm:"type:varchar(20); not null" json:"first_name" form:"first_name"`
	Last_Name    string `gorm:"type:varchar(20); not null" json:"last_name" form:"last_name"`
	Email        string `gorm:"type:varchar(50); unique; not null" json:"email" form:"email"`
	Password     string `gorm:"not null" json:"password" form:"password"`
	Phone_Number string `gorm:"type:varchar(15); not null" json:"phone_number" form:"phone_number"`
	Age          uint   `gorm:"type:varchar(15); not null" json:"age" form:"age"`
	CreatedAt    time.Time
	UpdatedAt    time.Time
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}
