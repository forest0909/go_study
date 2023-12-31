package main

import (
	"go_study/gin_chat/models"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func main() {
	db, err := gorm.Open(mysql.Open("root:123456@tcp(120.46.66.248:33306)/ginchat?charset=utf8mb4&parseTime=True&loc=Local"), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	// 迁移 schema
	db.AutoMigrate(&models.UserBasic{})

	// Create
	user := &models.UserBasic{}
	user.Username = "zzl"
	db.Create(user)

	// Read
	db.First(user, 1) // 根据整型主键查找
	//db.First(user, "code = ?", "D42") // 查找 code 字段值为 D42 的记录

	// Update - 将 product 的 price 更新为 200
	db.Model(user).Update("Password", 123456)
	// Update - 更新多个字段
	//db.Model(user).Updates(UserBasic{Price: 200, Code: "F42"}) // 仅更新非零值字段
	//db.Model(user).Updates(map[string]interface{}{"Price": 200, "Code": "F42"})

	// Delete - 删除 product
	//db.Delete(&userbasic, 1)
}
