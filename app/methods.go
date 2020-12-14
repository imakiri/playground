package app

import (
	"context"
	"errors"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"golang.org/x/crypto/bcrypt"
)

func (e *App) Detect(image []byte) ([]byte, error) {
	response, _ := e.services.FaceDetection.Detect(context.Background(), &core.DetectionRequest{Img: image})
	if err := response.GetErr(); err != nil {
		return nil, errors.New(err.String())
	}

	return response.GetImg().GetData(), nil
}

func (e *App) CreateUser(login string, password string, avatar []byte, name string) error {
	var c data.DBMainCreateUser
	var err error

	c.Request.Login = login
	c.Request.Avatar = avatar
	c.Request.Name = name
	c.Request.PassHash, err = bcrypt.GenerateFromPassword([]byte(password+e.app.Salt), e.app.HashCost)
	if err != nil {
		return err
	}

	err = e.data.DBMainCreateUser(&c)
	return err
}
