package postgres

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func Migrate(host string, port uint16, user, pass, db string) error {

	migrator, err := migrate.New(
		"file://postgres/migration",
		fmt.Sprintf("postgres://%s:%s@%s:%v/%s?sslmode=disable", user, pass, host, port, db))
	if err != nil {
		return err
	}
	err = migrator.Up()
	if err != nil {
		if err != migrate.ErrNoChange {
			return err
		}
	}

	return nil
}
