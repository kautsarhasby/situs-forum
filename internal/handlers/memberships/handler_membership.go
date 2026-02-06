package memberships

import (
	"context"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/middleware"
	"github.com/kautsarhasby/situs-forum/internal/model/memberships"
)

type membershipService interface {
	SignUp(ctx context.Context, request memberships.SignUpRequest) error
	Login(ctx context.Context, request memberships.LoginRequest) (string, string, error)
	ValidateRefreshToken(ctx context.Context, userID int64, request memberships.RefreshTokenRequest) (string, error)
}

type Handler struct {
	*gin.Engine
	membershipSVC membershipService
}

func NewHandler(api *gin.Engine, membershipSvc membershipService) *Handler {
	return &Handler{
		Engine:        api,
		membershipSVC: membershipSvc,
	}
}

func (h *Handler) RegisterRoute() {
	route := h.Group("memberships")
	route.GET("/ping", h.Ping)
	route.POST("/sign-up", h.SignUp)
	route.POST("/login", h.Login)

	routeRefresh := h.Group("memberships")
	routeRefresh.Use(middleware.AuthRefreshMiddleware())
	routeRefresh.POST("/refresh", h.Refresh)
}
