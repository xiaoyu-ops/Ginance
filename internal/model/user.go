package model

import (
	"gorm.io/gorm"
)

// user定义用户表结构
type User struct {
	// gorm.Model 中自动包含了 ID，createdat，updatedat，deletedat
	gorm.Model

	// 'gorm:"unique" 保证用户名不重复
	Username string `gorm:"unique;not null" json:"username"`

	// 密码之后会加密存，现在先预留
	Password string `gorm:"not null" json:"-"`

	// 邮箱
	Email string `gorm:"unique" json:"email"`
}
