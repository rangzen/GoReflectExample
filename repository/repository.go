package repository

import "context"

type Repository interface {
	DataA(ctx context.Context) string
	DataB(ctx context.Context) string
	DataC(ctx context.Context) string
}
