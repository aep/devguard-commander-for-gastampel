package server

import (
	"github.com/foolin/goview/supports/ginview"
	"github.com/gin-gonic/gin"
	"net/http"
    "github.com/devguardio/oha/db"
    "github.com/devguardio/oha/carrier"
)


func SetupUi(g *gin.Engine) {

    g.Static("/_ui/static/", "./static")
    g.Static("/_ui/api/",    "./static/api")

	g.HTMLRender = ginview.Default()

	g.GET("/_ui/", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}


		c.HTML(http.StatusOK, "index", gin.H{
            "user":    user,
		})
	})
	g.GET("/_ui/claim", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}


		c.HTML(http.StatusOK, "claim", gin.H{
            "user":    user,
		})
	})
	g.GET("/_ui/webhooks/new", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}

		c.HTML(http.StatusOK, "webhooks_new", gin.H{
            "user":     user,
		})
	})
	g.GET("/_ui/webhooks", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}

        streams, err  := db.ListStreams(user);
        if err != nil {
            panic(err);
        }


		c.HTML(http.StatusOK, "webhooks", gin.H{
            "user":    user,
            "streams":  streams,
		})
	})
	g.GET("/_ui/networks", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}


        conduits, err:= db.ListConduits(user);
        if err != nil {
            panic(err);
        }


        events, err:= db.ListConduitEvents(user, &conduits[0]);
        if err != nil {
            panic(err);
        }


		c.HTML(http.StatusOK, "network", gin.H{
            "user":     user,
            "conduit":  conduits[0],
            "events":   events,
		})
	})
	g.GET("/_ui/claim/is_valid_target", func(c *gin.Context) {

        target_string := c.Query("target");
        _, err := carrier.IdentityFromString(target_string);

        if err != nil {
            c.JSON(http.StatusBadRequest, gin.H{})
        } else {
            c.JSON(http.StatusOK, gin.H{})
        }

	})
	g.GET("/_ui/vault/", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}


        identities , err:= db.ListVaultIdentities(user);
        if err != nil {
            panic(err);
        }

        networks, err:= db.ListVaultNetworks(user);
        if err != nil {
            panic(err);
        }


        var vault []interface{};

        for _,v := range identities {
            vault = append(vault, gin.H{
                "Id":         v.Id,
                "Public":     v.Public,
                "CreatedAt":  v.CreatedAt,
                "Item":       "identity",
            });
        }

        for _,v := range networks {
            vault = append(vault, gin.H{
                "Id":         v.Id,
                "Public":     v.Public,
                "CreatedAt":  v.CreatedAt,
                "Item":       "network",
            });
        }


		c.HTML(http.StatusOK, "vault", gin.H{
            "user":    user,
            "vault":   vault,
		})
	})
	g.GET("/_ui/devices/", func(c *gin.Context) {
        user := CurrentUserUi(c);
        if user == nil {return}


        devices , err:= db.ListDevices(user);
        if err != nil {
            panic(err);
        }

		c.HTML(http.StatusOK, "devices", gin.H{
            "user":    user,
            "devices": devices,
		})
	})
}
