// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.27.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.adminDealExistsStmt, err = db.PrepareContext(ctx, adminDealExists); err != nil {
		return nil, fmt.Errorf("error preparing query AdminDealExists: %w", err)
	}
	if q.adminDeleteDealStmt, err = db.PrepareContext(ctx, adminDeleteDeal); err != nil {
		return nil, fmt.Errorf("error preparing query AdminDeleteDeal: %w", err)
	}
	if q.adminDeleteUserStmt, err = db.PrepareContext(ctx, adminDeleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query AdminDeleteUser: %w", err)
	}
	if q.adminGetDealForUpdateStmt, err = db.PrepareContext(ctx, adminGetDealForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query AdminGetDealForUpdate: %w", err)
	}
	if q.adminGetPitchRequestStmt, err = db.PrepareContext(ctx, adminGetPitchRequest); err != nil {
		return nil, fmt.Errorf("error preparing query AdminGetPitchRequest: %w", err)
	}
	if q.adminUpdateDealStmt, err = db.PrepareContext(ctx, adminUpdateDeal); err != nil {
		return nil, fmt.Errorf("error preparing query AdminUpdateDeal: %w", err)
	}
	if q.adminUpdateUserStmt, err = db.PrepareContext(ctx, adminUpdateUser); err != nil {
		return nil, fmt.Errorf("error preparing query AdminUpdateUser: %w", err)
	}
	if q.adminUserExistsStmt, err = db.PrepareContext(ctx, adminUserExists); err != nil {
		return nil, fmt.Errorf("error preparing query AdminUserExists: %w", err)
	}
	if q.adminViewDealsStmt, err = db.PrepareContext(ctx, adminViewDeals); err != nil {
		return nil, fmt.Errorf("error preparing query AdminViewDeals: %w", err)
	}
	if q.adminViewUsersStmt, err = db.PrepareContext(ctx, adminViewUsers); err != nil {
		return nil, fmt.Errorf("error preparing query AdminViewUsers: %w", err)
	}
	if q.countFilteredDealsStmt, err = db.PrepareContext(ctx, countFilteredDeals); err != nil {
		return nil, fmt.Errorf("error preparing query CountFilteredDeals: %w", err)
	}
	if q.createDealStmt, err = db.PrepareContext(ctx, createDeal); err != nil {
		return nil, fmt.Errorf("error preparing query CreateDeal: %w", err)
	}
	if q.createNewUserStmt, err = db.PrepareContext(ctx, createNewUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateNewUser: %w", err)
	}
	if q.createPitchRequestStmt, err = db.PrepareContext(ctx, createPitchRequest); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePitchRequest: %w", err)
	}
	if q.deletePitchRequestStmt, err = db.PrepareContext(ctx, deletePitchRequest); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePitchRequest: %w", err)
	}
	if q.filterDealsStmt, err = db.PrepareContext(ctx, filterDeals); err != nil {
		return nil, fmt.Errorf("error preparing query FilterDeals: %w", err)
	}
	if q.forgotPasswordStmt, err = db.PrepareContext(ctx, forgotPassword); err != nil {
		return nil, fmt.Errorf("error preparing query ForgotPassword: %w", err)
	}
	if q.getDealsByIdStmt, err = db.PrepareContext(ctx, getDealsById); err != nil {
		return nil, fmt.Errorf("error preparing query GetDealsById: %w", err)
	}
	if q.getDealsBySalesRepStmt, err = db.PrepareContext(ctx, getDealsBySalesRep); err != nil {
		return nil, fmt.Errorf("error preparing query GetDealsBySalesRep: %w", err)
	}
	if q.getDealsByStatusStmt, err = db.PrepareContext(ctx, getDealsByStatus); err != nil {
		return nil, fmt.Errorf("error preparing query GetDealsByStatus: %w", err)
	}
	if q.getPitchRequestForUpdateStmt, err = db.PrepareContext(ctx, getPitchRequestForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetPitchRequestForUpdate: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUserForUpdateStmt, err = db.PrepareContext(ctx, getUserForUpdate); err != nil {
		return nil, fmt.Errorf("error preparing query GetUserForUpdate: %w", err)
	}
	if q.pitchRequestExistStmt, err = db.PrepareContext(ctx, pitchRequestExist); err != nil {
		return nil, fmt.Errorf("error preparing query PitchRequestExist: %w", err)
	}
	if q.updatePassWordStmt, err = db.PrepareContext(ctx, updatePassWord); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePassWord: %w", err)
	}
	if q.updatePitchRequestStmt, err = db.PrepareContext(ctx, updatePitchRequest); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePitchRequest: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	if q.viewPitchRequestsStmt, err = db.PrepareContext(ctx, viewPitchRequests); err != nil {
		return nil, fmt.Errorf("error preparing query ViewPitchRequests: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.adminDealExistsStmt != nil {
		if cerr := q.adminDealExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminDealExistsStmt: %w", cerr)
		}
	}
	if q.adminDeleteDealStmt != nil {
		if cerr := q.adminDeleteDealStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminDeleteDealStmt: %w", cerr)
		}
	}
	if q.adminDeleteUserStmt != nil {
		if cerr := q.adminDeleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminDeleteUserStmt: %w", cerr)
		}
	}
	if q.adminGetDealForUpdateStmt != nil {
		if cerr := q.adminGetDealForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminGetDealForUpdateStmt: %w", cerr)
		}
	}
	if q.adminGetPitchRequestStmt != nil {
		if cerr := q.adminGetPitchRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminGetPitchRequestStmt: %w", cerr)
		}
	}
	if q.adminUpdateDealStmt != nil {
		if cerr := q.adminUpdateDealStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminUpdateDealStmt: %w", cerr)
		}
	}
	if q.adminUpdateUserStmt != nil {
		if cerr := q.adminUpdateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminUpdateUserStmt: %w", cerr)
		}
	}
	if q.adminUserExistsStmt != nil {
		if cerr := q.adminUserExistsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminUserExistsStmt: %w", cerr)
		}
	}
	if q.adminViewDealsStmt != nil {
		if cerr := q.adminViewDealsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminViewDealsStmt: %w", cerr)
		}
	}
	if q.adminViewUsersStmt != nil {
		if cerr := q.adminViewUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing adminViewUsersStmt: %w", cerr)
		}
	}
	if q.countFilteredDealsStmt != nil {
		if cerr := q.countFilteredDealsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing countFilteredDealsStmt: %w", cerr)
		}
	}
	if q.createDealStmt != nil {
		if cerr := q.createDealStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createDealStmt: %w", cerr)
		}
	}
	if q.createNewUserStmt != nil {
		if cerr := q.createNewUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createNewUserStmt: %w", cerr)
		}
	}
	if q.createPitchRequestStmt != nil {
		if cerr := q.createPitchRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPitchRequestStmt: %w", cerr)
		}
	}
	if q.deletePitchRequestStmt != nil {
		if cerr := q.deletePitchRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePitchRequestStmt: %w", cerr)
		}
	}
	if q.filterDealsStmt != nil {
		if cerr := q.filterDealsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing filterDealsStmt: %w", cerr)
		}
	}
	if q.forgotPasswordStmt != nil {
		if cerr := q.forgotPasswordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing forgotPasswordStmt: %w", cerr)
		}
	}
	if q.getDealsByIdStmt != nil {
		if cerr := q.getDealsByIdStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDealsByIdStmt: %w", cerr)
		}
	}
	if q.getDealsBySalesRepStmt != nil {
		if cerr := q.getDealsBySalesRepStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDealsBySalesRepStmt: %w", cerr)
		}
	}
	if q.getDealsByStatusStmt != nil {
		if cerr := q.getDealsByStatusStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getDealsByStatusStmt: %w", cerr)
		}
	}
	if q.getPitchRequestForUpdateStmt != nil {
		if cerr := q.getPitchRequestForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPitchRequestForUpdateStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUserForUpdateStmt != nil {
		if cerr := q.getUserForUpdateStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserForUpdateStmt: %w", cerr)
		}
	}
	if q.pitchRequestExistStmt != nil {
		if cerr := q.pitchRequestExistStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing pitchRequestExistStmt: %w", cerr)
		}
	}
	if q.updatePassWordStmt != nil {
		if cerr := q.updatePassWordStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePassWordStmt: %w", cerr)
		}
	}
	if q.updatePitchRequestStmt != nil {
		if cerr := q.updatePitchRequestStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePitchRequestStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
		}
	}
	if q.viewPitchRequestsStmt != nil {
		if cerr := q.viewPitchRequestsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing viewPitchRequestsStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                           DBTX
	tx                           *sql.Tx
	adminDealExistsStmt          *sql.Stmt
	adminDeleteDealStmt          *sql.Stmt
	adminDeleteUserStmt          *sql.Stmt
	adminGetDealForUpdateStmt    *sql.Stmt
	adminGetPitchRequestStmt     *sql.Stmt
	adminUpdateDealStmt          *sql.Stmt
	adminUpdateUserStmt          *sql.Stmt
	adminUserExistsStmt          *sql.Stmt
	adminViewDealsStmt           *sql.Stmt
	adminViewUsersStmt           *sql.Stmt
	countFilteredDealsStmt       *sql.Stmt
	createDealStmt               *sql.Stmt
	createNewUserStmt            *sql.Stmt
	createPitchRequestStmt       *sql.Stmt
	deletePitchRequestStmt       *sql.Stmt
	filterDealsStmt              *sql.Stmt
	forgotPasswordStmt           *sql.Stmt
	getDealsByIdStmt             *sql.Stmt
	getDealsBySalesRepStmt       *sql.Stmt
	getDealsByStatusStmt         *sql.Stmt
	getPitchRequestForUpdateStmt *sql.Stmt
	getUserStmt                  *sql.Stmt
	getUserForUpdateStmt         *sql.Stmt
	pitchRequestExistStmt        *sql.Stmt
	updatePassWordStmt           *sql.Stmt
	updatePitchRequestStmt       *sql.Stmt
	updateUserStmt               *sql.Stmt
	viewPitchRequestsStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                           tx,
		tx:                           tx,
		adminDealExistsStmt:          q.adminDealExistsStmt,
		adminDeleteDealStmt:          q.adminDeleteDealStmt,
		adminDeleteUserStmt:          q.adminDeleteUserStmt,
		adminGetDealForUpdateStmt:    q.adminGetDealForUpdateStmt,
		adminGetPitchRequestStmt:     q.adminGetPitchRequestStmt,
		adminUpdateDealStmt:          q.adminUpdateDealStmt,
		adminUpdateUserStmt:          q.adminUpdateUserStmt,
		adminUserExistsStmt:          q.adminUserExistsStmt,
		adminViewDealsStmt:           q.adminViewDealsStmt,
		adminViewUsersStmt:           q.adminViewUsersStmt,
		countFilteredDealsStmt:       q.countFilteredDealsStmt,
		createDealStmt:               q.createDealStmt,
		createNewUserStmt:            q.createNewUserStmt,
		createPitchRequestStmt:       q.createPitchRequestStmt,
		deletePitchRequestStmt:       q.deletePitchRequestStmt,
		filterDealsStmt:              q.filterDealsStmt,
		forgotPasswordStmt:           q.forgotPasswordStmt,
		getDealsByIdStmt:             q.getDealsByIdStmt,
		getDealsBySalesRepStmt:       q.getDealsBySalesRepStmt,
		getDealsByStatusStmt:         q.getDealsByStatusStmt,
		getPitchRequestForUpdateStmt: q.getPitchRequestForUpdateStmt,
		getUserStmt:                  q.getUserStmt,
		getUserForUpdateStmt:         q.getUserForUpdateStmt,
		pitchRequestExistStmt:        q.pitchRequestExistStmt,
		updatePassWordStmt:           q.updatePassWordStmt,
		updatePitchRequestStmt:       q.updatePitchRequestStmt,
		updateUserStmt:               q.updateUserStmt,
		viewPitchRequestsStmt:        q.viewPitchRequestsStmt,
	}
}
