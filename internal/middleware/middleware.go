package middleware

import (
	"errors"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/kautsarhasby/situs-forum/internal/configs"
	"github.com/kautsarhasby/situs-forum/pkg/jwt"
)

func AuthMiddleware() gin.HandlerFunc {
	secretKey:= configs.Get().Service.SecretJWT
	return func(ctx *gin.Context) {
		header:= ctx.Request.Header.Get("Authorization")
		header= strings.TrimSpace(header)

		if header == "" {
			ctx.AbortWithError(http.StatusUnauthorized,errors.New("missing token"))
		}

		userID, username, err := jwt.ValidateToken(header,secretKey)
		if err != nil {
			ctx.AbortWithError(http.StatusUnauthorized,err)
		}	
		ctx.Set("userID",userID)
		ctx.Set("username",username)
		ctx.Next()
	}
}