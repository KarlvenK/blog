package main

import (
	"blog/model"
	"blog/routes"
)

func main() {
	//引用数据库
	model.InitDb()

	routes.InitRouter()
}
