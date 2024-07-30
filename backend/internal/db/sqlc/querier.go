// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
)

type Querier interface {
	AdminDeleteDeal(ctx context.Context, id int64) error
	AdminDeleteUser(ctx context.Context, id int64) error
	AdminGetDealForUpdate(ctx context.Context, id int64) (Deal, error)
	AdminUpdateDeal(ctx context.Context, arg AdminUpdateDealParams) (Deal, error)
	AdminUpdateUser(ctx context.Context, arg AdminUpdateUserParams) (User, error)
	AdminViewUsers(ctx context.Context, arg AdminViewUsersParams) ([]User, error)
	CreateDeal(ctx context.Context, arg CreateDealParams) (Deal, error)
	CreateNewUser(ctx context.Context, arg CreateNewUserParams) (User, error)
	CreatePitchRequest(ctx context.Context, arg CreatePitchRequestParams) (PitchRequest, error)
	DeletePitchRequest(ctx context.Context, salesRepID int64) error
	GetDealsByAward(ctx context.Context, arg GetDealsByAwardParams) ([]Deal, error)
	GetDealsByCustomerAndService(ctx context.Context, arg GetDealsByCustomerAndServiceParams) ([]Deal, error)
	GetDealsByCustomerName(ctx context.Context, arg GetDealsByCustomerNameParams) ([]Deal, error)
	GetDealsByProfit(ctx context.Context, arg GetDealsByProfitParams) ([]Deal, error)
	GetDealsBySalesRep(ctx context.Context, arg GetDealsBySalesRepParams) ([]Deal, error)
	GetDealsBySalesRepAndAwarded(ctx context.Context, arg GetDealsBySalesRepAndAwardedParams) ([]Deal, error)
	GetDealsByServicesRendered(ctx context.Context, arg GetDealsByServicesRenderedParams) ([]Deal, error)
	GetDealsByStatus(ctx context.Context, status string) ([]Deal, error)
	GetPitchRequestForUpdate(ctx context.Context, id int64) (PitchRequest, error)
	GetUserForUpdate(ctx context.Context, id int64) (User, error)
	UpdatePitchRequest(ctx context.Context, arg UpdatePitchRequestParams) (PitchRequest, error)
	UpdateUser(ctx context.Context, arg UpdateUserParams) (User, error)
	ViewPitchRequests(ctx context.Context, arg ViewPitchRequestsParams) ([]PitchRequest, error)
}

var _ Querier = (*Queries)(nil)
