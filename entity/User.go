package entity

import "time"

type User struct {
	ID         int64     `gorm:"id" json:"ID"`
	Username   string    `gorm:"username" form:"username" json:"username"`
	Password   string    `gorm:"password" form:"password" json:"password"`
	CreateTime time.Time `gorm:"create_time" json:"createTime"`
}

func (u User) TableName() string {
	return "gorm_user"
}
