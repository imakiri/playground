package log

import (
	"context"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type Service struct {
	*zap.Logger
}

func (s Service) Log(_ context.Context, level pgx.LogLevel, msg string, data map[string]interface{}) {
	var fields []zap.Field
	if data != nil {
		fields = make([]zap.Field, len(data))
		for key, value := range data {
			fields = append(fields, zap.Any(key, value))
		}
	}

	switch level {
	case pgx.LogLevelError:
		s.Error(msg, fields...)
	case pgx.LogLevelInfo:
		s.Info(msg, fields...)
	}
}

func NewService() (*Service, error) {
	var service = new(Service)

	var err error
	if service.Logger, err = zap.NewDevelopment(); err != nil {
		return nil, err
	}

	return service, nil
}
