package app

import (
	"context"
	"errors"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"golang.org/x/crypto/bcrypt"
)

const hashCost = 10
const testKey = "testKey"

func NewApp(s core.Settings) (*App, error) {
	var a App
	var err error

	a.settings = s
	a.data, err = data.NewDB(s)
	if err != nil {
		return nil, err
	}

	return &a, nil
}

type App struct {
	settings core.Settings
	data     *data.DB
}

func (e App) Detect(image []byte) ([]byte, error) {
	response, _ := e.settings.Services.FaceDetection.Detect(context.Background(), &core.DetectionRequest{Img: image})
	if err := response.GetErr(); err != nil {
		return nil, errors.New(err.String())
	}

	return response.GetImg().GetData(), nil
}

func (e App) CreateUser(login string, password string, avatar []byte, name string) error {
	var c data.DBMainCreateUser
	var err error

	c.Request.Login = login
	c.Request.Avatar = avatar
	c.Request.Name = name
	c.Request.PassHash, err = bcrypt.GenerateFromPassword([]byte(password+e.settings.Salt), hashCost)
	if err != nil {
		return err
	}

	err = e.data.DBMainCreateUser(&c)
	return err
}
