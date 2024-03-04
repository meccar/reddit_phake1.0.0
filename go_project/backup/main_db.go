package db

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5/pgxpool"
	_ "github.com/lib/pq"
)

const (
	dbDriver = "postgres"
	dbSource = "postgresql://postgres:secret@localhost:5432/postgres?sslmode=disable"
)

// type Handlers struct {
// 	Repo Repo
// }

// func NewHandlers(repo Repo) *Handlers {
// 	return &Handlers{Repo: repo}
// }

// Repo defines all functions to execute db queries
type Repo interface {
	Querier
	SubmitFormTx(ctx context.Context, arg SubmitFormTxParams) (SubmitFormTxResult, error)
}

// SQLRepo provides all functions to execute SQL queries
type SQLRepo struct {
	connPool *pgxpool.Pool
	*Queries
}

// NewRepo creates a new repo
func NewRepo(connPool *pgxpool.Pool) Repo {
	return &SQLRepo{
		connPool: connPool,
		Queries:  New(connPool),
	}
}

// var q Repo

func dbMain() Repo {
	// Initialize the database connection
	// conn, err = sql.Open(dbDriver, dbSource)
	// if err != nil {
	// 	log.Fatal("Cannot connect to db:", err)
	// }
	ctx := context.Background()
	connPool, err := pgxpool.New(ctx, dbSource)
	fmt.Printf("\n connPool main_db %v \n", connPool)
	if err != nil {
		log.Fatal("cannot connect to db:", err)
	}

	return NewRepo(connPool)
}
