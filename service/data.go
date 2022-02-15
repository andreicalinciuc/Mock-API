package service

import (
	"context"
	"github.com/andreicalinciuc/mock-api/model"
)

// Data godoc
type Data interface {
	Find(ctx context.Context, id int64) (*model.Data, error)
	FindAll(ctx context.Context, page, limit uint64) ([]model.Data, uint64, error)
	Store(ctx context.Context, u *model.Data) error
	Update(ctx context.Context, u *model.Data) error
	Delete(ctx context.Context, id int64) error
}
