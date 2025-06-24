package client

import "context"

type Dao interface {
	Get(ctx context.Context, val any, fields ...any) any
	Set(ctx context.Context, val any, fields ...any) any
	Update(ctx context.Context, val any, fields ...any) any
	Delete(ctx context.Context, val any, fields ...any) any
}
