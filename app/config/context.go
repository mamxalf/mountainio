package config

import (
	"context"
	"time"
)

func DBContextTimeout(duration time.Duration) (context.Context, context.CancelFunc) {
	return context.WithTimeout(context.Background(), duration*time.Second)
}
