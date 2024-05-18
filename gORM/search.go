package main

import (
	"GO_code/entity"
	"github.com/gin-gonic/gin"
	"net/http"
)

func search() {
	ginServer := gin.Default()
	user := entity.User{}
	ginServer.POST("/search", func(context *gin.Context) {
		err := context.ShouldBind(&user)
		//Take不会走where SQL直接根据ID或者数据库的排序规则直接返回一条记录
		err = DB.Take(&user, "password=?", user.Password).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
		} else {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK,
				"msg":  "操作成功",
				"data": user,
			})
		}
	})
	ginServer.Run()
}

func Find() {
	ginServer := gin.Default()
	ginServer.POST("/find", func(context *gin.Context) {
		var users []entity.User
		userQuery := entity.User{}
		err := context.ShouldBind(&userQuery)
		//Take不会走where SQL直接根据ID或者数据库的排序规则直接返回一条记录
		query := DB.Where("password=?", userQuery.Password)
		if userQuery.Username != "" {
			query.Where("username=?", userQuery.Username)
		}
		err = DB.Find(&users, query).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
		} else {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK,
				"msg":  "操作成功",
				"data": users,
			})
		}
	})
	ginServer.Run()
}

// Save save是全字段更新，找到主键然后整条记录更新
func Save() {
	ginServer := gin.Default()
	ginServer.POST("/save", func(context *gin.Context) {

		userQuery := entity.User{}
		userInfo := entity.User{}
		err := context.ShouldBind(&userQuery)
		DB.Take(&userInfo, userQuery.ID)
		userInfo.Username = userQuery.Username
		userInfo.Password = userQuery.Password
		//如果结构体实例的主键字段是零值，即尚未在数据库中创建过对应记录，则 Save 方法会将该结构体作为新记录插入数据库。
		//如果结构体实例的主键字段已经有值，即数据库中已存在对应的记录，则 Save 方法会更新该记录的其他字段值。
		err = DB.Save(&userInfo).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
		} else {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK,
				"msg":  "操作成功",
				"data": userQuery,
			})
		}
	})
	ginServer.Run()
}

// 仅修改的字段
func Select() {
	ginServer := gin.Default()
	ginServer.POST("/select", func(context *gin.Context) {

		userQuery := entity.User{}
		err := context.ShouldBind(&userQuery)

		//query := DB.Where("username", userQuery.Username)

		err = DB.Select("username").Save(&userQuery).Error
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "服务器错误"})
		} else {
			context.JSON(http.StatusOK, gin.H{"code": http.StatusOK,
				"msg":  "操作成功",
				"data": userQuery,
			})
		}
	})
	ginServer.Run()
}
