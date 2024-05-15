package dto

type CreatePost struct {
	Title string `binding:"required,min=5,max=1000"`
	Post  string `binding:"required,min=5,max=5000"`
}
