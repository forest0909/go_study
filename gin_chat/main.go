package main

import (
	"go_study/gin_chat/router"
	"go_study/gin_chat/utils"
)

func main() {
	utils.InitConfig()
	utils.InitDB()
	r := router.Router()
	r.Run(":8081") // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}
