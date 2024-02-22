package postgresql

import (
	"ElDocManager/internal/config"
	"ElDocManager/pkg/utils"
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgx.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
}

func NewClient(ctx context.Context, sc config.PostgreStorageConfig) (conn *pgx.Conn, err error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Password, sc.Database)

	utils.DoWithAttemps(func() error {
		pgConfig, err := pgx.ParseConnectionString(dsn)
		if err != nil {
			return err
		}
		conn, err = pgx.Connect(pgConfig)
		if err != nil {
			return err
		}
		return nil
	}, sc.MaxAttemps, 5*time.Second)

	return conn, nil
}
