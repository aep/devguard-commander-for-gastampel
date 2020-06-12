package server;

import (
    "os/exec"
    "os"
    "github.com/devguardio/oha/db"
    "github.com/devguardio/oha/carrier"
    "time"
    "log"
)




var running = make(map[string]*exec.Cmd);

func InitConduitManager() {

    go func() {
        for {

            log.Println("managing conduits");

            var conduits []db.Conduit;
            dbc := db.DB.Find(&conduits)
            if dbc.Error != nil {
                panic(dbc.Error);
            }


            for _,conduit := range conduits {

                var user = db.User{Id: conduit.UserId};
                dbc := db.DB.First(&user)
                if dbc.Error != nil {
                    log.Println("conduit without user", dbc.Error, conduit);
                    continue;
                }
                skb, cerr := db.LoadDefaultSecretKit(&user);
                if cerr != nil {
                    log.Println("conduit without sk", cerr, conduit);
                    continue;
                }

                sk, cerr := carrier.SecretKitToString(skb);
                if cerr != nil {
                    log.Println("conduit without sk", cerr, conduit);
                    continue;
                }



                if running[conduit.Id.String()] == nil {
                    url := "ws://127.0.0.1:8080/conduits/" + conduit.Id.String() + "/agent/" + conduit.AgentKey;
                    cmd := exec.Command(os.Args[0], "conduit", url);
                    cmd.Env = append(os.Environ(),"SECRETKIT=" + sk)

                    cmd.Stdout = os.Stderr
                    cmd.Stderr = os.Stdout

                    log.Println("starting conduit", os.Args[0], "conduit", url);
                    err := cmd.Start()
                    if err != nil {
                        log.Println("failed to start conduit", err);
                    } else {
                        running[conduit.Id.String()] = cmd;
                        go func() {
                            cmd.Wait();
                            log.Println("conduit died: ", os.Args[0], "conduit", url);
                            running[conduit.Id.String()] = nil;
                        }();
                    }
                }
            }

            time.Sleep(time.Second * 10);
        }
    }();
}

