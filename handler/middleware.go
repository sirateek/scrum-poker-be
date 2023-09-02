package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/sirateek/poker-be/utils"
)

const (
	UserIDHeaderKey = "X-POKER-USER-ID"
)

func UseAuth(contextManager utils.ContextManager) gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := c.Request.Context()

		userID := c.GetHeader(UserIDHeaderKey)
		if userID != "" {
			ctx = contextManager.SetUserID(ctx, userID)
		}

		c.Request = c.Request.WithContext(ctx)
		c.Next()
	}
}
