package main

import (
	"blog.com/model"
	"blog.com/route"
)

func main() {
	db := model.GoConnect()
	db.AutoMigrate(&model.User{}, &model.Post{})
	r := route.SetupRouter()
	r.Run()
}
