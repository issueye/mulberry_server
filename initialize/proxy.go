package initialize

import (
	"context"
	"mulberry/app/downstream/engine"
)

func InitEngine(ctx context.Context) {
	engine.Start(ctx)
}
