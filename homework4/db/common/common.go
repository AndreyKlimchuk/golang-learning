package common

import (
	"context"
	"github.com/jackc/pgx/v4"

	"github.com/jackc/pgconn"
)

type Queryer interface {
	Exec(ctx context.Context, sql string, arguments ...interface{}) (commandTag pgconn.CommandTag, err error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
}