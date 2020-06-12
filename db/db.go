package db

import "C"
import (
    "github.com/jinzhu/gorm"
    _ "github.com/jinzhu/gorm/dialects/postgres"
    "os"
    "log"
    "errors"
    "github.com/devguardio/carrier/go"
)

var DB *gorm.DB;


func Init() {
    var err error;
    DB, err = gorm.Open("postgres", os.Getenv("DATABASE_URL"));
    if err != nil {
        panic(err)
    }

    if os.Getenv("GIN_MODE") != "release" {
        DB.LogMode(true)
    }

    DB.Exec(`CREATE EXTENSION IF NOT EXISTS "uuid-ossp"`);
    DB.AutoMigrate(&User{})
    DB.AutoMigrate(&Sub{})
    DB.AutoMigrate(&Token{})
    DB.AutoMigrate(&Identity{})
    DB.AutoMigrate(&Network{})
    DB.AutoMigrate(&Device{})
    DB.AutoMigrate(&Conduit{})
    DB.AutoMigrate(&ConduitEvent{})
    DB.AutoMigrate(&AllowedWebhookHost{})
    DB.AutoMigrate(&Stream{})

    //TODO this sucks
    DB.Exec("ALTER TABLE devices ADD CONSTRAINT unique_device_per_user UNIQUE (user_id, identity);");
    DB.Exec(`INSERT INTO allowed_webhook_hosts(id, user_id, host)
        VALUES ('10df07da-ed08-4fbb-b889-4a53efa62321' , '00000000-0000-0000-0000-000000000000','hooks.zapier.com');
    `);

    DB.Exec("update conduits set current_state = 'unknown';");
}

func LoadDefaultSecretKit(user *User) (*carrier.SecretKit, *carrier.Error) {
    var identity Identity;
    dbc := DB.Where("user_id = ?",user.Id).First(&identity)
    if dbc.Error != nil {
        return nil, &carrier.Error{
            Code: "NoSecrets",
            Message: dbc.Error.Error(),
        }
    }

    var network  Network;
    dbc = DB.Where("user_id = ?",user.Id).First(&network)
    if dbc.Error != nil {
        return nil, &carrier.Error{
            Code: "NoSecrets",
            Message: dbc.Error.Error(),
        }
    }

    secrets, err := carrier.SecretKitFromStringParts(identity.Secret, network.Secret);
    if err != nil {
        return nil, err;
    }

    return &secrets, nil;

}

func ListVaultIdentities(user *User) ([]Identity, error) {
    var ii []Identity;
    dbc := DB.Where("user_id = ?",user.Id).Find(&ii)
    if dbc.Error != nil {
        return nil, dbc.Error;
    }
    for i := range ii {
        ii[i].Secret = "";
    }

    return ii, nil;
}

func ListDevices(user *User) ([]Device, error) {
    var ii []Device;
    dbc := DB.Where("user_id = ?",user.Id).Find(&ii)
    if dbc.Error != nil {
        return nil, dbc.Error;
    }
    return ii, nil;
}

func CreateDevice (user *User, identity string) (*Device, *carrier.Error) {
    var device Device;
    dbc := DB.Where(Device{
        Identity:   identity,
        UserId:     user.Id,
        CreatedBy:  "claim",
    }).FirstOrCreate(&device);
    if dbc.Error != nil {
        return nil, &carrier.Error{
            Code: "Internal",
            Message: dbc.Error.Error(),
        }
    }
    return &device, nil;
}

func ListVaultNetworks(user *User) ([]Network, error) {
    var ii []Network;
    dbc := DB.Where("user_id = ?",user.Id).Find(&ii)
    if dbc.Error != nil {
        return nil, dbc.Error;
    }
    for i := range ii {
        ii[i].Secret = "";
    }

    return ii, nil;
}

func ListStreams(user *User) ([]Stream, error) {
    var ii []Stream;
    dbc := DB.
        Joins("left join conduits on conduits.id = streams.conduit_id").
        Where("conduits.user_id = ?",user.Id).
        Find(&ii)
    if dbc.Error != nil {
        return nil, dbc.Error;
    }
    return ii, nil;
}

func ListConduits(user *User) ([]Conduit, error) {
    var ii []Conduit;
    dbc := DB.Where("user_id = ?",user.Id).Find(&ii)
    if dbc.Error != nil {
        return nil, dbc.Error;
    }
    return ii, nil;
}

func ListConduitEvents(user *User, conduit *Conduit) ([]ConduitEvent, error) {
    var ii []ConduitEvent;
        dbc := DB.
            Joins("left join conduits on conduits.id = conduit_events.conduit_id").
            Where("conduits.user_id = ?",user.Id).
            Where("conduits.id = ?",conduit.Id).
            Order("conduit_events.created_at DESC").
            Limit(20).
            Find(&ii)
    if dbc.Error != nil {
        return nil, dbc.Error;
    }
    return ii, nil;
}

func CreateUserDefaults(user *User) error {

    id_public, id_secret, err := carrier.CreateIdentity();
    if err != nil {return errors.New(err.Message)}

    var ni = Identity{
        UserId:     user.Id,
        Public:     id_public,
        Secret:     id_secret,
    };
    dbc := DB.Save(&ni);
    if dbc.Error != nil { return dbc.Error; }

    net_public, net_secret, err := carrier.CreateNetwork();
    if err != nil {return errors.New(err.Message)}

    var nn = Network{
        UserId:     user.Id,
        Public:     net_public,
        Secret:     net_secret,
    };
    dbc = DB.Save(&nn);
    if dbc.Error != nil { return dbc.Error; }


    var nc = Conduit {
        UserId:         user.Id,
        NetworkId:      nn.Id,
        PrincipalId:    ni.Id,
        CurrentState:   "stopped",
        TargetState:    "running",
    }
    dbc = DB.Save(&nc);
    if dbc.Error != nil { return dbc.Error; }


    var ce = ConduitEvent {
        ConduitId:      nc.Id,
        Severity:       "info",
        Message:        "created",
    }

    dbc = DB.Save(&ce);
    if dbc.Error != nil { return dbc.Error; }

    return nil;
}


var ExternalTokenVerifier func(access_token string) (map[string]interface{}, error);

func UserForToken(access_token string) (*User, error) {

    var user User;
    dbc := DB.
        Joins("left join subs on subs.user_id = users.id").
        Joins("left join tokens on tokens.sub_id = subs.id").
        Where("tokens.id = ?", access_token).
        Where("age(tokens.created_at) < interval '7::days'").
        First(&user);

    if dbc.Error == nil {
        return &user, nil;
    }

    // token not found

    dbc = DB.Exec("delete from tokens where age(created_at) > interval '7::days' ;");
    if dbc.Error != nil {
        log.Println(dbc.Error);
    }

    var profile , err = ExternalTokenVerifier(access_token);
    if err != nil {
        return nil, err;
    }

    substr, ok := profile["sub"].(string);
    if !ok || substr == "" {
        return nil, errors.New("missing sub field in userinfo");
    }

    email, ok := profile["email"].(string);
    if !ok || email == "" {
        return nil, errors.New("missing email field in userinfo");
    }
    name,    ok := profile["name"].(string);
    picture, ok := profile["picture"].(string);


    // check if we know the Sub
    var sub Sub;
    dbc = DB.Where("sub = ?", substr).First(&sub);
    if dbc.Error == nil {
        // sub exists, update user
        user = User {
            Id:         sub.UserId,
            Email:      email,
            Name:       name,
            Picture:    picture,
        }
        dbc = DB.Save(&user);
        if dbc.Error != nil { return nil, dbc.Error; }
    } else {
        // sub doesnt exist, create user and sub
        user = User {
            Email:      email,
            Name:       name,
            Picture:    picture,
        };
        dbc = DB.Create(&user);
        if dbc.Error != nil { return nil, dbc.Error; }

        err = CreateUserDefaults(&user);
        if err != nil {
            return nil, err;
        }

        sub = Sub {
            Sub: substr,
            UserId: user.Id,
        }
        dbc = DB.Create(&sub);
    }

    // we have the sub, store the token
    dbc = DB.Create(&Token{
        Id:     access_token,
        SubId:  sub.Id,
    });
    if dbc.Error != nil { return nil, dbc.Error; }

    return UserForToken(access_token);

}


func IsWebhookHostAllowed(user *User, host string) bool {
    var al AllowedWebhookHost;
    dbc := DB.
        Where("(user_id = ? OR user_id = '00000000-0000-0000-0000-000000000000')", user.Id).
        Where("host = ?", host).
        First(&al);
    if dbc.Error != nil { return false ; }

    return al.Host == host;
}
