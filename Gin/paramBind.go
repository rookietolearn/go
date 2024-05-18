package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type Book struct {
	Name   string `form:"name"`
	Author string `form:"author"`
	Page   int    `form:"page"`
}

func bindData() {
	ginServer := gin.Default()
	ginServer.GET("mid", timeCount, index)

	ginServer.GET("/book", func(context *gin.Context) {
		var book Book
		err := context.ShouldBindQuery(&book)
		if err != nil {
			context.JSON(http.StatusBadGateway, gin.H{"msg": "bad"})
		} else {
			context.JSON(http.StatusOK, gin.H{"msg": "yes",
				"book.name":   book.Name,
				"book.author": book.Author,
				"book.page":   book.Page,
			})
		}
	})
	ginServer.Run(":8888")
}
