package router

import (
	"github.com/gin-gonic/gin"
	"go_study/gin_chat/service"
)

func Router() *gin.Engine {
	r := gin.Default()
	r.GET("/index", service.GetIndex)
	r.GET("/user/getUserList", service.GetUserList)

	return r
}
