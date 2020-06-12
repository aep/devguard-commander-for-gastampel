package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
)

func init() {
	up := func(db orm.DB) error {
        _, err := db.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`);
        if err != nil { return err}
        return nil
	}

	down := func(db orm.DB) error {
		return nil
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200614183204_init", up, down, opts)
}
