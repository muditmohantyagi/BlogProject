package controller

import (
	"net/http"

	"blog.com/dto"
	"blog.com/model"
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

	result, err := model.FindUserByEmail(InputDTO.Email)
	if err != nil {

		response := lib.Error("SQL error", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if result != 0 {

		response := lib.Error("User error", "user alreay exists", lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	password, err := helper.PwdEncription(InputDTO.Password)

	if err != nil {
		response := lib.Error("encripton error", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	mobile_no, err := helper.ConvertStoI(InputDTO.Mobile)
	if err != nil {

		response := lib.Error("Ivalid mobile no", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	var user model.User

	user.Name = InputDTO.Name
	user.Email = InputDTO.Email
	user.Mobile = mobile_no
	user.Password = password
	db := model.GoConnect()
	if result := db.Create(&user); result.Error != nil {

		response := lib.Error("SQL error", result.Error.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := lib.Success(true, "ok", user)
	c.JSON(http.StatusOK, response)
	// c.JSON(200, gin.H{
	// 	"message": user,
	// })
}
