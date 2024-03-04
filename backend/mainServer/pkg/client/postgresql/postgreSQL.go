package postgresql

import (
	"context"
	"fmt"
	"time"

	"github.com/Heatdog/ElDocManager/backend/mainServer/internal/config"
	repeatable "github.com/Heatdog/ElDocManager/backend/mainServer/pkg/utils"

	"github.com/jackc/pgconn"
	"github.com/jackc/pgx/v4"
	"github.com/jackc/pgx/v4/pgxpool"
)

type Client interface {
	Exec(ctx context.Context, sql string, args ...interface{}) (pgconn.CommandTag, error)
	Query(ctx context.Context, sql string, args ...interface{}) (pgx.Rows, error)
	QueryRow(ctx context.Context, sql string, args ...interface{}) pgx.Row
	Begin(ctx context.Context) (pgx.Tx, error)
	Close()
}

func NewClient(ctx context.Context, sc config.PostgreStorageConfig, maxAttemps int) (Client, error) {
	dsn := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", sc.Username, sc.Password, sc.Host, sc.Port, sc.Database)
	var conn *pgxpool.Pool
	var err error
	repeatable.DoWithAttemps(func() error {
		ctx, cancel := context.WithTimeout(ctx, 5*time.Second)
		defer cancel()

		conn, err = pgxpool.Connect(ctx, dsn)
		if err != nil {
			return err
		}
		return nil
	}, maxAttemps, time.Duration(maxAttemps*int(time.Second)))

	return conn, nil
}
