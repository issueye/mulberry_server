package initialize

import (
	"context"
	"mulberry/host/app/downstream/engine"
)

func InitEngine(ctx context.Context) {
	engine.Start(ctx)
}
