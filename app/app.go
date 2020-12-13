package app

import (
	"context"
	"errors"
	"github.com/imakiri/playground/core"
)

func NewApp(s core.Settings) *App {
	return &App{settings: s}
}

type App struct {
	settings core.Settings
}

func (e App) Detect(image []byte) ([]byte, error) {
	response, _ := e.settings.Services.FaceDetection.Detect(context.Background(), &core.DetectionRequest{Img: image})
	if err := response.GetErr(); err != nil {
		return nil, errors.New(err.String())
	}

	return response.GetImg().GetData(), nil
}
