package posts

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/middleware"
	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

type postService interface{
	CreatePost (ctx context.Context,userId int64, request posts.CreatePostRequest) error
}


type Handler struct {
	*gin.Engine
	postSVC postService
}

func NewHandler(api *gin.Engine, postSvc postService) *Handler{
	return  &Handler{
		Engine: api,
		postSVC: postSvc,
	}
}

func (h *Handler) RegisterRoute(){
	route := h.Group("posts")
	route.Use(middleware.AuthMiddleware())
	route.POST("/create",h.CreatePost)
	
}