package entity

import "time"

type Exercise struct {
	ID          uint64 `gorm:"primary_key;auto_increment;UNIQUE" json:"id"`
	Title       string `json:"title" binding:"required" gorm:"type:varchar(40)"`
	Description string `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	Img         string `json:"img" binding:"max=200" gorm:"type:varchar(200)"`
	Email       string `json:"email" validate:"required,email" gorm:"type:varchar(40)"`
}

type Video struct {
	ID          uint64    `gorm:"primary_key;auto_increment" json:"id"`
	Title       string    `json:"title" binding:"min=2,max=200" validate:"is-cool" gorm:"type:varchar(200)"`
	Description string    `json:"description" binding:"max=200" gorm:"type:varchar(200)"`
	URL         string    `json:"url" binding:"required,url" gorm:"type:varchar(100)"`
	Author      Exercise  `json:"author" binding:"required" gorm:"foreignkey: PersonID"`
	PersonID    uint64    `json:"-"`
	CreatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"created_at"`
	UpdatedAt   time.Time `json:"-" gorm:"default:CURRENT_TIMESTAMP" json:"update_at"`
}
