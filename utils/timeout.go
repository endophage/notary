package utils

import (
	"time"

	"golang.org/x/net/context"
)

// Timeout returns a context with the timeout set to the provided duration
func Timeout(ctx context.Context, after time.Duration) context.Context {
	ctx, _ = context.WithTimeout(ctx, after)
	return ctx
}

// Timeout30 returns the background context with a 30 second timeout set
func Timeout30() context.Context {
	return Timeout(context.Background(), 30*time.Second)
}
