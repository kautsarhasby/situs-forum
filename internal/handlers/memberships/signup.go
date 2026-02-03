package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
)

func (h *Handler) SignUp(c *gin.Context){
	ctx:= c.Request.Context()

	var request memberships.SignUpRequest
	if err:=c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error" : err.Error(),
		})
		return 
	}

	err := h.membershipSVC.SignUp(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error" : err.Error(),
		})
	}

	c.Status(http.StatusCreated)
	
}