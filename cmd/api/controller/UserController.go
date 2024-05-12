package controller

import "github.com/gin-gonic/gin"

type UserController struct{}

func (con UserController) RegisterUser(c *gin.Context) {
	c.JSON(200, gin.H{
		"message": "pong",
	})
}
