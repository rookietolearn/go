package main

import (
	"GO_code/entity"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"time"
)

var DB *gorm.DB

func DBinit() {
	dsn := "root:1230@(127.0.0.1:3306)/sequence?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("连接数据库失败, error=" + err.Error())
	}
	// 连接成功
	fmt.Println(db)
	//把其DB类型实例赋给定义好的全局变量
	DB = db
}
func CreateSingle() {
	user := entity.User{
		Username:   "wy",
		Password:   "567899",
		CreateTime: time.Now(),
	}
	err := DB.Create(&user).Error
	fmt.Println(err)
	fmt.Println(user.ID)
}
func CreateArr() {
	var userList []entity.User
	for i := 0; i < 10; i++ {
		userList = append(userList, entity.User{
			Username:   fmt.Sprintf("wy%d", i),
			Password:   "2345",
			CreateTime: time.Now(),
		})
	}
	DB.Create(&userList)
}
