package controller

import (
	"net/http"

	"blog.com/dto"
	"blog.com/model"
	"blog.com/pkg/helper"
	"blog.com/pkg/lib"
	"github.com/gin-gonic/gin"
)

type PostController struct{}

func (con PostController) CreatePost(c *gin.Context) {
	var InputDTO dto.CreatePost

	if errDTO := c.ShouldBindJSON(&InputDTO); errDTO != nil {
		msg := lib.ValidationError(errDTO)
		c.AbortWithStatusJSON(http.StatusBadRequest, msg)
		return
	}

	helper.Trimmer(&InputDTO)
	user_id := lib.GetUserID(c.GetHeader("Token"))
	var post model.Post
	post.Title = InputDTO.Title
	post.Post = InputDTO.Post
	post.UserId = user_id
	db := model.GoConnect()
	if result := db.Create(&post); result.Error != nil {
		lib.ELog.Error(result.Error.Error())
		response := lib.Error("SQL error", result.Error.Error(), lib.EmptyObj{})
		c.JSON(http.StatusBadRequest, response)
		return
	}
	response := lib.Success(true, "ok", "Post created successfully")
	c.JSON(http.StatusOK, response)
}
