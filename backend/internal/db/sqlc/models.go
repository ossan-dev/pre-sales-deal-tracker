// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"database/sql"
	"time"
)

type Deal struct {
	ID              int64
	PitchID         int64
	SalesRepName    string
	CustomerName    string
	ServiceToRender string
	// it is either ongoing or closed
	Status string
	// tag for the status can be, working on: survey, proposal, costing depending
	StatusTag           string
	CurrentPitchRequest string
	NetTotalCost        sql.NullString
	Profit              sql.NullString
	CreatedAt           time.Time
	UpdatedAt           sql.NullTime
	ClosedAt            sql.NullTime
	Awarded             bool
}

type PitchRequest struct {
	ID              int64
	SalesRepID      int64
	Status          string
	CustomerName    string
	PitchTag        string
	CustomerRequest string
	RequestDeadline time.Time
	AdminViewed     bool
	CreatedAt       time.Time
	UpdatedAt       sql.NullTime
}

type User struct {
	ID       int64
	Username string
	// Role of the user, e.g., sales rep, admin, manager
	Role      string
	FullName  string
	Email     string
	Password  string
	UpdatedAt sql.NullTime
	CreatedAt time.Time
}
