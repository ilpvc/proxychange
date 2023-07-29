package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
)

func main() {
	S := gin.Default()
	S.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "测试一下"})
		fmt.Println(&S)
	})
	err := S.Run(":8080")
	if err != nil {
		fmt.Println("服务启动失败")
	}
}
