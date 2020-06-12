package server

import (
    "net/http"
    "log"
    "github.com/devguardio/oha/db"
    "encoding/base64"
    "crypto/rand"
	oidc "github.com/coreos/go-oidc"
	"os"
	"golang.org/x/oauth2"
	"github.com/gin-gonic/gin"
	"context"
    "fmt"
    "net/url"
    "encoding/json"
    "errors"
    "strings"
)

type Authenticator struct {
	Provider *oidc.Provider
	Config   oauth2.Config
	Ctx      context.Context
}

var oauthconfig oauth2.Config;
var authenticator Authenticator;

func SetupAuth(g *gin.Engine) {

	ctx := context.Background()
	provider, err := oidc.NewProvider(ctx, "https://" + os.Getenv("AUTH0_DOMAIN") + "/")
	if err != nil {
        panic(err)
    }

    oauthconfig = oauth2.Config{
        ClientID:     os.Getenv("AUTH0_CLIENT_ID"),
        ClientSecret: os.Getenv("AUTH0_CLIENT_SECRET"),
        RedirectURL:  os.Getenv("AUTH0_CALLBACK_URL"),
        Endpoint: 	  provider.Endpoint(),
        Scopes:       []string{oidc.ScopeOpenID, "profile", "email", "user_metadata", "app_metadata", "picture", "user_profile"},
    };

    authenticator = Authenticator{
		Provider: provider,
		Config:   oauthconfig,
		Ctx:      ctx,
	};

    db.ExternalTokenVerifier = func(access_token string) (map[string]interface{}, error){
        var oc = oauthconfig.Client(ctx, &oauth2.Token{
            AccessToken: access_token,
        });
        r, err := oc.Get("https://" + os.Getenv("AUTH0_DOMAIN") + "/userinfo");
        if err != nil {
            return nil, err;
        }

        defer r.Body.Close()
        var profile  map[string]interface{};
        json.NewDecoder(r.Body).Decode(&profile);

        if r.StatusCode != 200 {
            return nil, errors.New(fmt.Sprintf("userinfo endpoint returned %v %v ", r.StatusCode, profile));
        }

        return profile, nil;
    }



    /*
    var kk = []byte(os.Getenv("SESSION_KEY"));
    if len(kk) < 1 {
        panic("missing SESSION_KEY env")
    }
    var kkhash = sha512.Sum512([]byte(os.Getenv("SESSION_KEY")));

    store := cookie.NewStore(kkhash[:]);
    store.Options(sessions.Options{
          MaxAge:     604800,
          Secure:     true,
          HttpOnly:   true,
          SameSite:   http.SameSiteLaxMode,
    });
    g.Use(sessions.Sessions("session", store))
    */

    g.GET("/_ui/login",             LoginHandler);
    g.GET("/_ui/logout",            LogoutHandler);
    g.GET("/_ui/oauth/callback",    CallbackHandler);
}

func CurrentUserUi(c *gin.Context) *db.User {
    user, err := CurrentUser(c);
    if err != nil {
        log.Println(err);
        ClearCookies(c);
        c.Redirect(http.StatusTemporaryRedirect, "/_ui/login");
        return nil;
    }
    return user;
}

func CurrentUserApi(c *gin.Context) *db.User {
    user, err := CurrentUser(c);
    if err != nil {
        c.JSON(http.StatusUnauthorized, &Error{
            Code:       "UnAuthorized",
            Message:    err.Error(),
        });
        return nil;
    }
    return user;
}

func CurrentUser(c *gin.Context) (*db.User, error) {


    var token string;
    var err   error;

    maybe_auth_header := c.GetHeader("Authorization")
    if maybe_auth_header != "" {
        maybe_auth_header_l := strings.Split(c.GetHeader("Authorization"), "Bearer ")
        if len(maybe_auth_header_l) == 2 {
            token = maybe_auth_header_l[1]
        }
    }

    if token == "" {
        token, err = c.Cookie("access_token");
        if err != nil { return  nil, err}
    }

    user,err := db.UserForToken(token);
    if err != nil {
        return nil, err;
    }

    return user, nil;
}

func LoginHandler(c *gin.Context) {
    // Generate random state
    b := make([]byte, 32)
    _, err := rand.Read(b)
    if err != nil {
        c.JSON(http.StatusInternalServerError, Error{
            Code:       "Internal",
            Message:    err.Error(),
        });
        return
    }
    state := base64.StdEncoding.EncodeToString(b)


    ClearCookies(c);
    http.SetCookie(c.Writer, &http.Cookie{
        Name:       "auth_state",
        Value:      url.QueryEscape(state),
        MaxAge:     86400,
        Path:       "/",
    });
    c.Redirect(http.StatusTemporaryRedirect, authenticator.Config.AuthCodeURL(state) + "&prompt=login");
}


func ClearCookies(c *gin.Context) {
    http.SetCookie(c.Writer, &http.Cookie{
        Name:       "access_token",
        MaxAge:     -1,
        Path:       "/",
    });
    http.SetCookie(c.Writer, &http.Cookie{
        Name:     "auth_state",
        MaxAge:   -1,
        Path:       "/",
    });
}

func LogoutHandler(c *gin.Context) {
    ClearCookies(c);
    c.Redirect(http.StatusSeeOther, "/");
}


func CallbackHandler(c *gin.Context) {

	// Redirect again to set cookie with strict samesite policy
    didagain := c.Query("doubleredirect");
    if didagain == "" {
        again := c.Request.URL;
        q := again.Query();
        q.Add("doubleredirect", "true");
        again.RawQuery = q.Encode();

        c.Redirect(http.StatusSeeOther, again.String())
        return;
    }

    state, err := c.Cookie("auth_state");

    ClearCookies(c);

    if err != nil {
        log.Println(err);
        c.JSON(http.StatusInternalServerError, Error{
            Code:       "Session",
            Message:    "no login flow session state",
        });
		return
	}

    fmt.Println("cookie auth_state: ", state);

	if c.Query("state") != state {
        c.JSON(http.StatusInternalServerError, Error{
            Code:       "Session",
            Message:    "state parameter mismatch",
        });
		return
	}

	token, err := authenticator.Config.Exchange(c, c.Query("code"))
	if err != nil {
        log.Println(err);
        c.JSON(http.StatusUnauthorized, Error{
            Code:       "Unauthorized",
            Message:    err.Error(),
        });
        return
	}

    log.Println(token)

	rawIDToken, ok := token.Extra("id_token").(string)
	if !ok {
        c.JSON(http.StatusBadRequest, Error{
            Code:       "Unauthorized",
            Message:    "No id_token field in oauth2 token.",
        });
		return
	}

	oidcConfig := &oidc.Config{
		ClientID: os.Getenv("AUTH0_CLIENT_ID"),
	}

    _ , err = authenticator.Provider.Verifier(oidcConfig).Verify(c, rawIDToken)
	if err != nil {
        log.Println(err);
        c.JSON(http.StatusBadRequest, Error{
            Code:       "Unauthorized",
            Message:    err.Error(),
        });
		return
	}

    http.SetCookie(c.Writer, &http.Cookie{
        Name:       "access_token",
        Value:      url.QueryEscape(token.AccessToken),
        MaxAge:     604800,
        Secure:     true,
        HttpOnly:   true,

        // TODO strict isnt working. The cross-site redirect from oauth BEFORE setting the cookie makes the whole pageload cross-site
        //    we're going to have to do Origin verification manually
        SameSite:   http.SameSiteLaxMode,
        Path:       "/",
    });

	if err != nil {
        log.Println(err);
        ClearCookies(c);
        c.JSON(http.StatusInternalServerError, Error{
            Code:       "Internal",
            Message:    err.Error(),
        });
        return;
	}

	// Redirect to logged in page
    c.Redirect(http.StatusSeeOther, "/");
}

