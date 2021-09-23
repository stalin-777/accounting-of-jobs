package postgres

import (
	"context"

	"github.com/jackc/pgx/v4/pgxpool"
	"github.com/juju/errors"
)

func NewPgxConnPool(host string, port uint16, user, pass, db string) (*pgxpool.Pool, error) {

	pgxconfig, err := pgxpool.ParseConfig("")
	if err != nil {
		return nil, errors.Annotate(err, "Unable to initialize config")
	}
	pgxconfig.ConnConfig.Host = host
	pgxconfig.ConnConfig.Port = port
	pgxconfig.ConnConfig.Database = db
	pgxconfig.ConnConfig.User = user
	pgxconfig.ConnConfig.Password = pass

	connPool, err := pgxpool.ConnectConfig(context.Background(), pgxconfig)
	if err != nil {
		return nil, errors.Annotate(err, "Failed to connect database instance")
	}

	return connPool, nil
}
