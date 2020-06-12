package main

import (
	"log"
	"os"

	"github.com/go-pg/pg/v9"
	migrations "github.com/robinjoseph08/go-pg-migrations/v2"
    "github.com/joho/godotenv"
)

const directory = "."

func main() {
    err := godotenv.Load("../.env")
    if err != nil {
        panic(err);
    }

    opt, err := pg.ParseURL(os.Getenv("DATABASE_URL"));
    if err != nil {
        panic(err)
    }

    db := pg.Connect(opt)

	err = migrations.Run(db, directory, os.Args)
	if err != nil {
		log.Fatalln(err)
	}
}
