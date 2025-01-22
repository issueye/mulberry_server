package initialize

import (
	"context"
	"mulberry/internal/app/downstream/engine"
)

func InitEngine(ctx context.Context) {
	engine.Start(ctx)
}
