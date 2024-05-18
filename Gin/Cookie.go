package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func Cookie() {
	ginServer := gin.Default()
	ginServer.GET("/cookie", func(context *gin.Context) {
		cookie, err := context.Cookie(`gin_cookie`)
		cookie2, err := context.Cookie(`cookie`)
		if err != nil {
			cookie = "NotSet"
			context.SetCookie("gin_cookie", "test", 3600, "/", "localhost", false, true)
		}
		fmt.Printf("Cookie value: %s ,%s\n", cookie, cookie2)
	})
	ginServer.Run(":8888")
}
