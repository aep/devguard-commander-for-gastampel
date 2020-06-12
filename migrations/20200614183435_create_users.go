package main

import (
	"github.com/go-pg/pg/v9/orm"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
    dbd "github.com/devguardio/commander/db"
)

func init() {
	up := func(db orm.DB) error {
        err := orm.CreateTable(db, &dbd.User{}, nil);
        if err != nil { return err }
        err = orm.CreateTable(db, &dbd.Auth{}, nil);
		return err
	}

	down := func(db orm.DB) error {
        err := orm.DropTable(db, &dbd.User{}, nil);
        if err != nil { return err }
        err = orm.DropTable(db, &dbd.Auth{}, nil);
		return err
	}

	opts := migrations.MigrationOptions{}

	migrations.Register("20200614183435_create_users", up, down, opts)
}
