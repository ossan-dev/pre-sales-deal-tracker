package db

import "github.com/jackc/pgx/v5/pgxpool"

const (
	DefaultUserPassword = "vasDeal45"
)
// Store holds function definitions for interacting with the database
// executes queries and transactions
type Store interface {
	Querier
}

// SqlStore holds fields required to interact with database
type SqlStore struct {
	connPool *pgxpool.Pool
	*Queries
	defaultUserPassword string
}

// NewStore creates a new Store with dbpool
func NewStore(dbpool *pgxpool.Pool) Store {
	return &SqlStore {
		connPool: dbpool,
		Queries: New(dbpool),
	}
}