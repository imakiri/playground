package cfg

import (
	"github.com/imakiri/erres"
	"gopkg.in/yaml.v3"
	"io/ioutil"
	"os"
)

type TypeCfg string

const (
	YAML TypeCfg = "yaml"
	ENV  TypeCfg = "env"
)

type Service struct {
	path   string
	config *Config
	secret *Secret
}

func (s Service) readConfig() error {
	var raw, err = ioutil.ReadFile(s.path)
	if err != nil {
		return err
	}

	return yaml.Unmarshal(raw, s.config)
}

func (s Service) GetAvatarConfig() (*ConfigService, error) {
	var conf = new(ConfigService)

	*conf = s.config.Services.Avatar

	return conf, nil
}

func (s Service) GetPostgresConfig(name string) (*ConfigDatabasePostgres, error) {
	var conf = new(ConfigDatabasePostgres)
	var ok bool

	if *conf, ok = s.config.Connections.Postgres[name]; !ok {
		return nil, erres.NotFound
	}

	return conf, nil
}

func (s Service) GetPostgresSecret(name string) (*SecretDatabasePostgres, error) {
	var conf = new(SecretDatabasePostgres)
	var ok bool

	if *conf, ok = s.secret.Connections.Postgres[name]; !ok {
		return nil, erres.NotFound
	}

	return conf, nil
}

func (s Service) GetMongoConfig(name string) (*ConfigDatabaseMongo, error) {
	var conf = new(ConfigDatabaseMongo)
	var ok bool

	if *conf, ok = s.config.Connections.Mongo[name]; !ok {
		return nil, erres.NotFound
	}

	return conf, nil
}

func (s Service) GetMongoSecret(name string) (*SecretDatabaseMongo, error) {
	var conf = new(SecretDatabaseMongo)
	var ok bool

	if *conf, ok = s.secret.Connections.Mongo[name]; !ok {
		return nil, erres.NotFound
	}

	return conf, nil
}

func NewService(file string) (*Service, error) {
	var s = new(Service)
	s.config = new(Config)
	s.config.Connections.Postgres = make(map[string]ConfigDatabasePostgres)
	s.secret = new(Secret)

	var wd, err = os.Getwd()
	if err != nil {
		return nil, err
	}

	s.path = wd + "\\configs\\" + file

	if err = s.readConfig(); err != nil {
		return nil, err
	}

	return s, nil
}
