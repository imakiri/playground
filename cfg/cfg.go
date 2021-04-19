package cfg

import (
	"github.com/imakiri/gorum/utils"
)

type TypeCfg string

const (
	YAML TypeCfg = "yaml"
	ENV  TypeCfg = "env"
)

type Service struct {
	yaml_path string
}

func (Service) GetPostgresCfg(t TypeCfg) (*ConfigDatabasePostgres, error) {
	var conf = new(ConfigDatabasePostgres)

	switch t {
	case YAML:
	case ENV:

	}

	return conf, nil
}

func (s Service) GetPostgresSecret(t TypeCfg) (*SecretDatabasePostgres, error) {
	var secr = new(SecretDatabasePostgres)

	switch t {
	case YAML:
		var err = utils.ReadYAML(s.yaml_path, secr)
		if err != nil {
			return nil, err
		}
	case ENV:

	}

	return secr, nil
}

func (Service) GetMongoCfg(t TypeCfg) (*ConfigDatabaseMongo, error) {
	var conf = new(ConfigDatabaseMongo)

	switch t {
	case YAML:
	case ENV:

	}

	return conf, nil
}

func (Service) GetMongoSecret(t TypeCfg) (*SecretDatabaseMongo, error) {
	var secr = new(SecretDatabaseMongo)

	switch t {
	case YAML:
	case ENV:

	}

	return secr, nil
}
