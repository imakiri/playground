package app

import (
	"context"
	"errors"
	"github.com/google/uuid"
	"github.com/imakiri/playground/core"
	"golang.org/x/crypto/bcrypt"
)

func (e *App) Detect(u uuid.UUID, image []byte) ([]byte, error) {
	if !checkPermissionForUUID(e.gate, u, core.FN_Detect) {
		return nil, core.Status_AccessDenied
	}

	response, _ := e.services.FaceDetection.Detect(context.Background(), &core.DetectionRequest{Img: image})
	if err := response.GetErr(); err != nil {
		return nil, errors.New(err.String())
	}

	return response.GetImg().GetData(), nil
}

func (e *App) CreateUser(u uuid.UUID, login string, password string, avatar []byte, name string) error {
	if !checkPermissionForUUID(e.gate, u, core.FN_CreateUser) {
		return core.Status_AccessDenied
	}

	var c core.ContainerCreateUser
	var err error

	c.Request.Login = login
	c.Request.Avatar = avatar
	c.Request.Name = name
	c.Request.PassHash, err = bcrypt.GenerateFromPassword([]byte(password+e.app.Salt), e.app.HashCost)
	if err != nil {
		return err
	}

	err = e.data.CreateUser(&c)
	return err
}

func (e *App) Login(login string, password string) (uuid.UUID, error) {
	var c core.ContainerGetUserPassHash
	var err error

	c.Request.Login = login
	err = e.data.GetPassHash(&c)
	if err != nil {
		return uuid.New(), core.Status_AccessDenied
	}

	err = bcrypt.CompareHashAndPassword(c.Response.PassHash, []byte(password))
	if err != nil {
		return uuid.New(), core.Status_AccessDenied
	}

	return uuid.New(), nil
}
