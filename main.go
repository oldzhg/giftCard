package main

import (
	"giftCard/server"
	"giftCard/utils"
)

func main() {
	utils.InitDB()
	// 初次部署，修改InitDB函数中的dsn数据库相关信息，然后运行main.go，会自动创建表
	// 然后注释掉下面三行，重新运行
	//utils.Db.AutoMigrate(&model.Category{})
	//utils.Db.AutoMigrate(&model.Contact{})
	//utils.Db.AutoMigrate(&model.Card{})
	server.Run()
}
