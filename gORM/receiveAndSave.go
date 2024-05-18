package main

import (
	"GO_code/entity"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func receiveAndSave() {
	ginServer := gin.Default()
	var user entity.User
	ginServer.POST("/userinfo", func(context *gin.Context) {
		err := context.ShouldBind(&user)
		user.CreateTime = time.Now()
		err = DB.Create(&user).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
		} else {
			context.JSON(http.StatusOK, gin.H{"msg": "操作成功"})
		}
	})
	ginServer.Run()
}
