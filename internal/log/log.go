package log

import (
	"context"
	"github.com/elastic/go-elasticsearch/v8"
	"github.com/jackc/pgx/v4"
	"go.uber.org/zap"
)

type Service struct {
	client *elasticsearch.Client
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
	var conf = elasticsearch.Config{
		Addresses:             nil,
		Username:              "",
		Password:              "",
		CloudID:               "",
		APIKey:                "",
		ServiceToken:          "",
		Header:                nil,
		CACert:                nil,
		RetryOnStatus:         nil,
		DisableRetry:          false,
		EnableRetryOnTimeout:  false,
		MaxRetries:            0,
		CompressRequestBody:   false,
		DiscoverNodesOnStart:  false,
		DiscoverNodesInterval: 0,
		EnableMetrics:         false,
		EnableDebugLogger:     false,
		DisableMetaHeader:     false,
		RetryBackoff:          nil,
		Transport:             nil,
		Logger:                nil,
		Selector:              nil,
		ConnectionPoolFunc:    nil,
	}

	service.client, err = elasticsearch.NewClient(conf)
	if err != nil {
		return nil, err
	}

	if service.Logger, err = zap.NewDevelopment(); err != nil {
		return nil, err
	}

	return service, nil
}
