package controller

import (
	"net/http"

	"blog.com/dto"
	"blog.com/pkg/helper"
	"blog.com/pkg/lib"
	"github.com/gin-gonic/gin"
)

type UserController struct{}

func (con UserController) RegisterUser(c *gin.Context) {
	var InputDTO dto.Register

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	helper.Trimmer(&InputDTO)
	c.JSON(200, gin.H{
		"message": InputDTO,
	})
}
