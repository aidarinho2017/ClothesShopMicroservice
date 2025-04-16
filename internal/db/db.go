package db

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
)

type DB struct {
	Conn *pgx.Conn
}

func NewDB(dsn string) (*DB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		return nil, fmt.Errorf("No db connection %w", err)
	}

	fmt.Println("Connected to DB")
	return &DB{Conn: conn}, nil
}

func (db *DB) Close() {
	db.Conn.Close(context.Background())
	fmt.Println("connection closed")
}

const DATABASE_URL = "postgres://postgres:postgres@localhost:5432/postgres?sslmode=disable"
