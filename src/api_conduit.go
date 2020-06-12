/*
 * Devguard Open Hardware API
 *
 * Although devguard does offer direct end to end encrypted connections, some users may prefer a third party to establish a connection for them, as it is done by data hungry cloud corporations.  The Open Hardware API can be called with any standard https client by passing an secretkit in the headers, and a devguard hosted webservice will establish a device connection on behalf of the user.  We still do not collect any data, however due to the nature of webservices, interception by third parties is possible and the OHA service comes with no waranty whatsoever. 
 *
 * API version: 1.0.0
 * Generated by: OpenAPI Generator (https://openapi-generator.tech)
 */

package server

import (
	"net/http"
    "log"

	"github.com/gin-gonic/gin"
    "github.com/devguardio/oha/db"
    "github.com/devguardio/oha/conduit"
    "github.com/gorilla/websocket"
    "github.com/davecgh/go-spew/spew"
    "github.com/cornelk/hashmap"

    "encoding/json"
)




type ConnectedAgent struct {
    Hup chan bool
}
var ConduitAgents hashmap.HashMap;



// SystemDiscoveryGet - List available remote streams
func ConduitAgentWs(c *gin.Context) {
    if  c.Param("conduit")  == "" || c.Param("agent")  == "" {
        c.JSON(http.StatusBadRequest, gin.H{})
        return;
    }


    var dbconduit db.Conduit;
    dbc := db.DB.
        Where("conduits.id = ? and conduits.agent_key = ?", c.Param("conduit"), c.Param("agent")).
        First(&dbconduit);

    if dbc.Error != nil {
        c.JSON(http.StatusUnauthorized, gin.H{})
        return;
    }

    if _, ok := ConduitAgents.Get(dbconduit.Id.String()); ok {
        c.JSON(http.StatusConflict, gin.H{})
        return;
    }

    var wsupgrader = websocket.Upgrader{
        ReadBufferSize:  1024,
        WriteBufferSize: 1024,
    }

    conn, err := wsupgrader.Upgrade(c.Writer, c.Request, nil)
    if err != nil {
        log.Printf("Failed to set websocket upgrade: %+v", err)
        return
    }

    db.DB.Model(&dbconduit).Update("current_state", "disconnected");
    defer db.DB.Model(&dbconduit).Update("current_state", "disconnected");



    hup := make(chan bool,100)
    go func() {
        for range hup {
            var streams[]db.Stream;
            dbc = db.DB.Where("streams.conduit_id = ? and streams.disabled=false", dbconduit.Id).Find(&streams);
            if dbc.Error != nil {
                log.Println(dbc.Error);
                return;
            }


            var streams_config []conduit.ConfigStream;

            for _, s := range streams {
                streams_config = append(streams_config, conduit.ConfigStream{
                    Path:               s.Path,
                    RestartDelay:       s.RestartDelay,
                    Webhook:            &conduit.Webhook {
                        Url:            s.WebhookUrl,
                    },
                });
            }

            msg, err := json.Marshal(conduit.Commander2Conduit {
                Config: &conduit.Config{
                    Streams: streams_config,
                },
            });
            if err != nil {
                panic(err);
            }
            conn.WriteMessage(websocket.TextMessage, msg)
        }
    }();

    ConduitAgents.Set(dbconduit.Id.String(), &ConnectedAgent {
        Hup: hup,
    });
    defer close(hup);
    defer func() {ConduitAgents.Del(dbconduit.Id.String());}();
    hup <- true


    for {
        _, s , err := conn.ReadMessage()
        if err != nil {
            break
        }

        var msg conduit.Conduit2Commander;
        err = json.Unmarshal(s, &msg);
        if err != nil {
            continue;
        }

        spew.Dump(msg);

        if msg.State != "" {
            db.DB.Model(&dbconduit).Update("current_state", msg.State);
        }
        if msg.LogEvent != nil && msg.LogEvent.Message != "" {
            db.DB.Create(&db.ConduitEvent{
                ConduitId:  dbconduit.Id,
                Severity:   msg.LogEvent.Severity,
                Code:       msg.LogEvent.Code,
                Message:    msg.LogEvent.Message,
            });
        }
        if msg.Stats != nil && msg.Stats.Connections != 0 {
            db.DB.Model(&dbconduit).Updates(db.Conduit{
                StatsCurrentConnections:    msg.Stats.Connections,
                StatsCurrentCpuUsage:       msg.Stats.Cpu,
                StatsCurrentMemUsage:       msg.Stats.Mem,
            });
        }
        if msg.WebhookUrlFailed != nil && msg.WebhookUrlFailed.Url != "" {
            db.DB.Exec(`update streams set disabled=true, last_error=? where webhook_url=? and conduit_id=?`,
                msg.WebhookUrlFailed.Reason,
                msg.WebhookUrlFailed.Url,
                dbconduit.Id,
            );
            hup <- true
        }
    }
}

