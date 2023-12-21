package service

import (
	"github.com/gin-gonic/gin"
	"go_study/gin_chat/models"
)

func GetUserList(c *gin.Context) {

	data := models.GetUserList()

	c.JSON(200, gin.H{
		"code": 200,
		"msg":  "success",
		"data": gin.H{
			"name": data,
		},
	})
}
