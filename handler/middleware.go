package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/utils"
	"net/http"
)

const (
	UserIDHeaderKey = "X-POKER-USER-ID"
)

func UseAuth(contextManager utils.ContextManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID := c.GetHeader(UserIDHeaderKey)
		if userID == "" {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"message": "unauthorized",
			})
			return
		}
		ctx = contextManager.SetUserID(ctx, userID)

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
