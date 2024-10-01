// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
)

type Querier interface {
	CreateUser(ctx context.Context, arg CreateUserParams) (int64, error)
	GetNumberOfAdminUsers(ctx context.Context, role string) (int64, error)
	GetUserByUsername(ctx context.Context, username string) (User, error)
}

var _ Querier = (*Queries)(nil)
