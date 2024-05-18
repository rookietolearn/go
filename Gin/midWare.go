package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

func timeCount(c *gin.Context) {
	fmt.Println("timeCount is running")
	start := time.Now()
	time.Sleep(1 * time.Second)
	cost := time.Since(start)
	fmt.Printf("%v\n", cost)
}
func index(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"msg": "index"})
}
