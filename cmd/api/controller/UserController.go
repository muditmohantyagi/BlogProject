package controller

import (
	"net/http"
	"strconv"

	"blog.com/dto"
	"blog.com/model"
	"blog.com/pkg/helper"
	"blog.com/pkg/lib"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
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
		lib.ELog.Error(err.Error())
		response := lib.Error("SQL error", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if result != 0 {
		lib.WLog.Warn("user alreay exists")
		response := lib.Error("User error", "user alreay exists", lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	password, err := helper.PwdEncription(InputDTO.Password)

	if err != nil {
		lib.ELog.Error(err.Error())
		response := lib.Error("encripton error", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	mobile_no, err := helper.ConvertStoI(InputDTO.Mobile)
	if err != nil {
		lib.ELog.Error(err.Error())
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
		lib.ELog.Error(result.Error.Error())
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

func (con UserController) Login(c *gin.Context) {
	var InputDTO dto.Login

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	helper.Trimmer(&InputDTO)
	user_data, err := model.FindUserDataByEmailId(InputDTO.Email)
	if err != nil {
		lib.ELog.Error(err.Error())
		response := lib.Error("SQL error", err.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if user_data == nil {

		response := lib.Error("MSG", "User not found", lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user_data.Password), []byte(InputDTO.Password)); err != nil {
		error_response := lib.Error("Ivalid password", "Password is incorrect", lib.EmptyObj{})
		lib.ILog.Info(err.Error())
		c.JSON(http.StatusBadRequest, error_response)
		return
	}
	/////

	JwtToken := lib.GenerateToken(strconv.Itoa(int((user_data.ID))), 1)

	response := lib.Success(true, "ok", JwtToken)
	c.JSON(http.StatusOK, response)
}
