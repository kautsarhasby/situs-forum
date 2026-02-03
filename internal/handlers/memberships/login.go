package memberships

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
)

func (h *Handler) Login(c *gin.Context){
	ctx:= c.Request.Context()

	var request memberships.LoginRequest
	if err:=c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H {
			"error" : err.Error(),
		})
		return 
	}

	accessToken ,err := h.membershipSVC.Login(ctx, request)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error" : err.Error(),
		})
	}

	response:=memberships.LoginResponse{
		AccessToken: accessToken,
	}
	c.JSON(http.StatusOK, response)
	
}