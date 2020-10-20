package entity

import (
	"time"
)

type Video struct {
	ID          int64     `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=32,required" validate:"StartsWithCapital" gorm:"type:varchar(32)"` // custom validator
	Description string    `json:"description" binding:"max=128" gorm:"type:varchar(128)"`
	URL         string    `json:"url" binding:"url,required" gorm:"type:varchar(256);UNIQUE"`
	Author      Person    `json:"author" binding:"required" gorm:"foreignKey:PersonID"`
	PersonID    int64     `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP"`
}

type Person struct {
	ID        int64  `gorm:"primary_key;auto_increment" json:"id"`
	FirstName string `json:"firstname" binding:"required" gorm:"type:varchar(32)"`
	LastName  string `json:"lastname" binding:"required" gorm:"type:varchar(32)"`
	Age       int8   `json:"age" binding:"gte=1,lte=130"`
	Email     string `json:"email" binding:"email,required" gorm:"type:varchar(256)"`
}
