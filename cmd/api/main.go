package main

import (
	"blog.com/model"
	router "blog.com/route"
)

func main() {
	db := model.GoConnect()
	db.AutoMigrate(&model.User{})
	r := router.SetupRouter()
	r.Run()
}
