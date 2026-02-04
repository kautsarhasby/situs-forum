package posts

import (
	"errors"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/model/posts"
)

func (h *Handler) CreateComment(c * gin.Context){
	ctx := c.Request.Context()

	var request posts.CreateCommentRequest
	if err:= c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		return
	}

	userID:= c.GetInt64("userID")
	postIDStr:= c.Param("postID")
	postID, err:= strconv.ParseInt(postIDStr,10,64)
	if err!= nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error" : errors.New("PostID on params invalid"),
		})
	}

	if err:=h.postSVC.CreateComment(ctx,postID,userID,request); err != nil{
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}

	c.Status(http.StatusCreated)

}