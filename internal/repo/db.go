package repo

import (
	"context"
	"errors"
	"example.com/m/internal/model"
)

var (
	ErrURLNotFound = errors.New("url not found")
	ErrURLExists   = errors.New("url exists")
)

type NoteRepo interface {
	Create(ctx context.Context, note *model.User) (int64, error)
	Get(ctx context.Context, id int64) (*model.User, error)
}
