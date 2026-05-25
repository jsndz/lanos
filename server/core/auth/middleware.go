package auth

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Middleware struct {
	SessionService *SessionService
}

func NewMiddleware(sessionService *SessionService) *Middleware {
	return &Middleware{
		SessionService: sessionService,
	}
}

func (m *Middleware) AuthenticateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token, err := ctx.Cookie("session_token")
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "missing session token",
			})
			return
		}

		session, err := m.SessionService.GetSessionByToken(
			ctx.Request.Context(),
			token,
		)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "invalid session",
			})
			return
		}

		ctx.Set("user_id", session.UserID)
		ctx.Set("session_id", session.ID)

		ctx.Next()
	}
}
