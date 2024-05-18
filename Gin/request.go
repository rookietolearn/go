package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main1() {
	ginServer := gin.Default()
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{"msg": "hello"})
	})
	ginServer.POST("post_hello", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "post_hello"})
	})
	//http://localhost:8888/get_query?userId=123&pwd=345和
	ginServer.GET("/get_query", func(context *gin.Context) {
		userId := context.Query("userId")
		pwd := context.Query("pwd")
		context.JSON(http.StatusOK, gin.H{"userId": userId,
			"pwd": pwd})

	})
	//context.Param
	//http://localhost:8888/get_param/123/455
	ginServer.GET("/get_param/:userId/:pwd", func(context *gin.Context) {
		userId := context.Param("userId")
		pwd := context.Param("pwd")
		context.JSON(http.StatusOK, gin.H{"userId": userId,
			"pwd": pwd})

	})
	//Form参数
	ginServer.POST("post_form", func(context *gin.Context) {
		username := context.PostForm("username")
		pwd := context.PostForm("pwd")
		context.JSON(http.StatusOK, gin.H{"userId": username,
			"pwd": pwd})
	})
	//无匹配的路由
	ginServer.NoRoute(func(context *gin.Context) {
		context.JSON(http.StatusNotFound, gin.H{"msg": "404"})
	})
	//Any 方法用于注册一个处理器（handler），该处理器可以匹配所有 HTTP 方法（GET、POST、PUT、DELETE 等），并在请求中调用相应的处理函数。
	ginServer.Any("/any", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "any_enter"})
	})

	//路由组
	user := ginServer.Group("/user")
	user.GET("/get", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "user_get"})
	})
	user.POST("/post", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{"msg": "user_post"})
	})
	ginServer.Run(":8888")
}
