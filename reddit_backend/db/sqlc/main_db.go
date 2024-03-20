package db

import (
	"context"
	"github.com/rs/zerolog/log"
	"github.com/jackc/pgx/v5/pgxpool"


	_ "github.com/lib/pq"
)
// "github.com/jackc/pgx/v5"

// type Handlers struct {
// 	queries *Queries
// }

// func NewHandlers(queries *Queries) *Handlers {
// 	return &Handlers{queries: queries}
// }

// Store defines all functions to execute db queries and transactions
type Handler interface {
	Querier
	VerifyEmail(ctx context.Context, arg *CreateVerifyEmailParams) (CreateVerifyEmailTxResult, error)
}

// SQLStore provides all functions to execute SQL queries and transactions
type Handlers struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewStore creates a new store
func NewHandlers(connPool *pgxpool.Pool) Handler {
	return &Handlers{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

func InitialiseDB(DB_URL string) (*pgxpool.Pool, error) {
	// conn, err := pgx.Connect(context.Background(), "postgresql://"+user+":"+password+"@"+host+"/postgres?sslmode=disable")
	// conn, err := pgx.Connect(context.Background(), DB_URL)
	conn, err := pgxpool.New(context.Background(), DB_URL)

	if err != nil {
		log.Error().Err(err).Msg("Cannot connect to db")
	}

	// Check the connection
	if err := conn.Ping(context.Background()); err != nil {
		log.Error().Err(err).Msg("Cannot ping database")
	}

	return conn, nil
}

func CloseDB(conn *pgxpool.Pool) {
	if conn != nil {
		conn.Close()
	}
}
