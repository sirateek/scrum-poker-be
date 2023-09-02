package utils

import "context"

type ContextKey string

const (
	UserIDContextKey ContextKey = "userID"
)

type ContextManager struct {
}

func (c *ContextManager) GetUserID(ctx context.Context) string {
	return ctx.Value(UserIDContextKey).(string)
}

func (c *ContextManager) SetUserID(ctx context.Context, userID string) context.Context {
	return context.WithValue(ctx, UserIDContextKey, userID)
}
