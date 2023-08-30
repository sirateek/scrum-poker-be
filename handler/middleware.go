package handler

import (
	"context"
	"github.com/gin-gonic/gin"
)

func UseAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		ctx := context.Background()
		ctx = context.WithValue(ctx, "A", "B")
		c.Request = c.Request.WithContext(ctx)
		c.Set("A", "B")
		c.Next()
	}
}
