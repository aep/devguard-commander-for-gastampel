package main

import (
	"log"
	//server  "github.com/devguardio/commander/src"
    "github.com/devguardio/commander/db"
    "github.com/joho/godotenv"
    //"github.com/alecthomas/kong"
    "os"
    "net/url"
)

func main() {

    err := godotenv.Load()
    if err != nil {
        panic(err);
    }

    db.Init();

    //router := server.NewRouter()
    //log.Fatal(router.Run(":8080"))
}
