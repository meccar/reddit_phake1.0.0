package db

import (
	"context"
	"log"

	"github.com/jackc/pgx/v5"

	_ "github.com/lib/pq"
)

type Handlers struct {
	queries *Queries
}

func NewHandlers(queries *Queries) *Handlers {
	return &Handlers{queries: queries}
}

func InitialiseDB(DB_URL string) (*pgx.Conn, error) {
	// conn, err := pgx.Connect(context.Background(), "postgresql://"+user+":"+password+"@"+host+"/postgres?sslmode=disable")
	conn, err := pgx.Connect(context.Background(), DB_URL)

	if err != nil {
		log.Fatal("Cannot connect to db:", err)
	}

	// Check the connection
	if err := conn.Ping(context.Background()); err != nil {
		log.Fatal("Cannot ping database:", err)
	}

	return conn, nil
}

func CloseDB(conn *pgx.Conn) {
	if conn != nil {
		conn.Close(context.Background())
	}
}
