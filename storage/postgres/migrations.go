package postgres

import (
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"github.com/pkg/errors"
)

func (db *DB) Migrate() error {

	str := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		db.DBConfig.DBHost, db.DBConfig.DBPort,
		db.DBConfig.DBUser, db.DBConfig.DBPass, db.DBConfig.DBName)

	m, err := migrate.New("file://./storage/migrations", str)
	if err != nil {
		db.Logger.Fatalf("%v", err)
		return errors.WithStack(err)
	}

	err = m.Up()
	if err != nil && err != migrate.ErrNoChange {
		db.Logger.Fatalf("%v", err)
		return errors.WithStack(err)
	}
	return nil
}
