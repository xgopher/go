package models

import (
	"time"
)

// User 用户表结构
type User struct {
	ID        uint `gorm:"primary_key" json:"-"`
	CreatedAt time.Time	`json:"-"`
	UpdatedAt time.Time	`json:"-"`
	Name      string `gorm:"size:255;unique_index"`  // 用户名
	Password  string `gorm:"type:char(32);not null" json:"-"` // 密码
	Salt      string `gorm:"type:char(6);not null" json:"-"`  // 加密因子
	NickName  string `gorm:"size:255;not null" json:"nick_name"`       // 昵称
}

// TableName 设置表名
func (model *User) TableName() string {
	return "users"
}
