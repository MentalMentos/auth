package service

import (
	"context"
)

type Service interface {
	Create(ctx context.Context, username string, password string) (int64, error)
	//Delete(ctx context.Context, username string) error
	Get(ctx context.Context, id int64) (string, error)
	//Update(ctx context.Context, id string, username string, password string) error
}
