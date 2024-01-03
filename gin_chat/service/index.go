package service

import "github.com/gin-gonic/gin"

func GetIndex(c *gin.Context) {

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"name": "gin-chat",
		},
	})
}
