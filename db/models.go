package db

import (
    "github.com/satori/go.uuid"
    "time"
    "strings"
)

type Conduit struct {
    Id              uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt       time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    AgentKey        string      `gorm:";not null" sql:"DEFAULT:uuid_generate_v4()"`

    UserId          uuid.UUID   `gorm:";not null"`
    PrincipalId     uuid.UUID   `gorm:";not null"`
    NetworkId       uuid.UUID   `gorm:";not null"`

    CurrentState    string
    TargetState     string

    StatsCurrentConnections uint
    StatsCurrentCpuUsage    float64
    StatsCurrentMemUsage    float64
}

type Stream struct {
    Id              uuid.UUID   `json:"id"              gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt       time.Time   `json:"created_at"      gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    ConduitId       uuid.UUID   `json:"conduit_id"      gorm:";not null"`
    Identity        string      `json:"identity,omitempty"`
    Path            string      `json:"path"            gorm:";not null"`
    RestartDelay    uint        `json:"restart_delay"   gorm:";not null"`

    WebhookUrl      string      `json:"webhook_url,omitempty"`

    Disabled        bool        `json:"disabled"    sql:"DEFAULT:false"`
    LastError       string      `json:"last_error"`
}

type AllowedWebhookHost struct {
    Id              uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt       time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    UserId          uuid.UUID   `gorm:";not null"` // '00000000-0000-0000-0000-000000000000' means allowed for everyone
    Host            string      `gorm:";not null"`
}


type ConduitEvent struct {
    Id              uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt       time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    ConduitId       uuid.UUID   `gorm:";not null"`

    Severity        string
    Code            string
    Message         string
}

type Device struct {
    Id          uuid.UUID   `json:"id"          gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt   time.Time   `json:"created_at"  gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    UserId      uuid.UUID   `json:"user_id"     gorm:";not null"`
    Identity    string      `json:"identity"    gorm:";not null"`

    CreatedBy   string      `json:"created_by"  gorm:";not null"`
}

type Identity struct {
    Id          uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt   time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    UserId      uuid.UUID   `gorm:";not null"`
    Public      string      `gorm:";not null"`
    Secret      string      `gorm:";not null" json:"-"`
}

type Network struct {
    Id          uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt   time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    UserId      uuid.UUID   `gorm:";not null"`
    Public      string      `gorm:";not null"`
    Secret      string      `gorm:";not null" json:"-"`
}

type User struct {
    Id          uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    CreatedAt   time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`

    Email       string      `gorm:";not null"`
    Name        string
    Picture     string
}

func (user *User) Initials() string {
    if user.Name == "" {
        return "";
    }
    l := strings.Split(user.Name, " ")

    s := ""
    for _,e := range l {
        s += string(e[0]);
        if len(s) > 2 {
            break;
        }
    }
    return s;
}

type Sub struct {
    Id          uuid.UUID   `gorm:"type:uuid;primary_key;" sql:"DEFAULT:uuid_generate_v4()"`
    Sub         string      `gorm:";not null"`
    UserId      uuid.UUID   `gorm:";not null"`
    CreatedAt   time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`
}

type Token struct {
    Id          string      `gorm:";primary_key;not null"`
    SubId       uuid.UUID   `gorm:";not null"`
    CreatedAt   time.Time   `gorm:";not null" sql:"DEFAULT:(now() at time zone 'utc')"`
}



