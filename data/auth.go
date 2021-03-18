package data

import (
	"context"
	"github.com/aidarkhanov/nanoid"
	"github.com/imakiri/gorum/cfg"
	"github.com/imakiri/gorum/erres"
	_ "github.com/jackc/pgx/stdlib"
	"github.com/jmoiron/sqlx"
	"google.golang.org/grpc"
)

type (
	ModelAuthKey            string
	ModelAuthExpirationDate int64

	ModelAuthLogin    []byte
	ModelAuthPassword []byte

	ModelAuthObscureCookie struct {
		UserID
		Key            ModelAuthKey
		ExpirationDate ModelAuthExpirationDate
	}
	ModelAuthHashedLogpass struct {
		UserID
		Login    ModelAuthLogin
		Password ModelAuthPassword
	}

	ViewAuthObscureCookie struct {
		PemID          ModelUserPemID
		Key            ModelAuthKey
		ExpirationDate ModelAuthExpirationDate
	}
	ViewAuthHashedLogpass struct {
		PemID    ModelUserPemID
		Login    ModelAuthLogin
		Password ModelAuthPassword
	}
)

type ConfigAuth interface {
	Get4DataAuth(ctx context.Context, in *cfg.Request, opts ...grpc.CallOption) (*cfg.DataAuth, error)
}

type Auth struct {
	configCached *cfg.DataAuth
	config       ConfigAuth
	db           *sqlx.DB
}

func NewAuth(cs ConfigAuth) (*Auth, error) {
	var s Auth
	var err error

	if cs == nil {
		return nil, erres.E_InvalidArgument.SetTime("").SetDescription("ConfigAuth is nil")
	}

	s.config = cs
	s.configCached, err = s.config.Get4DataAuth(context.Background(), &cfg.Request{})
	if err != nil {
		return nil, err
	}

	s.db, err = sqlx.Connect("pgx", s.configCached.GetDSN())
	if err != nil {
		return nil, erres.E_ConnectionError.SetTime("").SetDescription(err.Error())
	}

	return &s, err
}

//

func (a Auth) AddCookie(uuid ModelUserUUID, cookie ViewAuthObscureCookie) error {
	var c ModelAuthObscureCookie

	if uuid == "" {
		c.UUID = ModelUserUUID(nanoid.New())
	}

	c.Key = cookie.Key
	c.PemID = cookie.PemID
	c.ExpirationDate = cookie.ExpirationDate

	var _, err = a.db.NamedQuery("INSERT INTO main.auth.cookie VALUES (:key, :uuid, :pemid, :expirationDate)", c)
	return err
}

func (a Auth) GetCookie(key ModelAuthKey, container *ViewAuthObscureCookie) error {
	return a.db.Get(container, "SELECT uuid, pemid, expiration_date FROM main.auth.cookie WHERE key = $1", key)
}

func (a Auth) DeleteCookie(uuid ModelUserUUID) error {
	var _, err = a.db.Exec("DELETE FROM main.auth.cookie WHERE uuid = $1", uuid)
	return err
}

//

func (a Auth) AddLogpass(uuid ModelUserUUID, logpass ViewAuthHashedLogpass) error {
	var l ModelAuthHashedLogpass

	if uuid == "" {
		l.UUID = ModelUserUUID(nanoid.New())
	}

	l.PemID = logpass.PemID
	l.Password = logpass.Password
	l.Login = logpass.Login

	var _, err = a.db.NamedQuery("INSERT INTO main.auth.logpass VALUES (:uuid, :login, :password, :pemid)", l)
	return err
}

func (a Auth) GetLogpass(login ModelAuthLogin, container *ViewAuthHashedLogpass) error {
	return a.db.Get(container, "SELECT uuid, pemid, password FROM main.auth.logpass WHERE login = $1", login)
}

func (a Auth) DeleteLogpass(uuid ModelUserUUID) error {
	var _, err = a.db.Exec("DELETE FROM main.auth.logpass WHERE uuid = $1", uuid)
	return err
}
