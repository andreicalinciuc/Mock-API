package service

import (
	"context"
	"github.com/andreicalinciuc/mock-api/model"
)

// Data godoc
type Data interface {
	Find(ctx context.Context, id int64) (*model.Data, error)
	FindByUserName(ctx context.Context, username string) (*model.Data, error)
	Login(ctx context.Context, session, username, password string) (*model.Data, string, error)
	FindAll(ctx context.Context, page, limit uint64) ([]model.Data, uint64, error)
	Store(ctx context.Context, u *model.Data) error
	Update(ctx context.Context, u *model.Data) error
	Delete(ctx context.Context, id int64) error
}
