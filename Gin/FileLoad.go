package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func FileLoadSingle() {
	ginServer := gin.Default()
	ginServer.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("f1")
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "上传失败"})
			return
		}
		dst := fmt.Sprintf("D:/A_workspace/fileTemp/%s", file.Filename)
		err = context.SaveUploadedFile(file, dst)
		if err != nil {
			context.JSON(http.StatusInternalServerError, gin.H{"msg": "存储失败"})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("'%s' uploaded!", file.Filename),
		})
	})
	ginServer.Run(":8888")
}

func FileUploadMulti() {
	ginServer := gin.Default()
	ginServer.POST("/uploadMulti", func(context *gin.Context) {
		form, _ := context.MultipartForm()
		files := form.File["file123"] //"file"是上传时的key的名字
		for i, file := range files {
			dst := fmt.Sprintf("D:/A_workspace/fileTemp/%s_%d", file.Filename, i)
			context.SaveUploadedFile(file, dst)
		}
		context.JSON(http.StatusOK, gin.H{
			"message": fmt.Sprintf("%d files uploaded!", len(files)),
		})
	})
	ginServer.Run(":8888")
}
