package db

import (
	"database/sql"
	"log"

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
// type Repo struct {
// 	connPool *pgxpool.Pool
// }

// // New creates a new Repo object
// func NewRepo(connPool *pgxpool.Pool) *Repo {
// 	return &Repo{
// 		connPool: connPool,
// 	}
// }

var conn *sql.DB

func dbMain() *Queries {
	// Initialize the database connection
	var err error

	conn, err := sql.Open(dbDriver, dbSource)
	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}
	defer conn.Close()

	// Check the connection
	pingErr := conn.Ping()
	if pingErr != nil {
		log.Fatal("Cannot ping database:", pingErr)
	}

	q := New(conn)
	return q
	// ctx := context.Background()
	// connPool, err := pgxpool.New(ctx, dbSource)
	// fmt.Printf("\n connPool main_db %v \n", connPool)
	// if err != nil {
	// 	log.Fatal("cannot connect to db:", err)
	// }

	// q := New(connPool)
	// return q
}
