package core

import (
	"context"
	"github.com/imakiri/playground/cfg"
	error2 "github.com/imakiri/playground/erres"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
)

// Error

type StatusCode codes.Code

func (e StatusCode) Error() string {
	return string(e)
}

//

type ServiceName string

func (s ServiceName) String() string {
	return string(s)
}

type Service interface {
	Name() ServiceName
}

// FunctionID

type FunctionID uint16

func (e FunctionID) FID() uint16 {
	return uint16(e)
}

const FID_Detect FunctionID = 10
const FID_CreateUser FunctionID = 11

const FID_AuthLogin FunctionID = 0
const FID_AuthLogout FunctionID = 1

//

type ActionID uint64
type Meta struct {
	Status error2.Error
}

func Connect(c *cfg.Data) (*pgx.Conn, error) {
	var db *pgx.Conn
	var err error

	if c.GetDSN() == "" {
		return nil, error2.E_InvalidArgument
	}

	db, err = pgx.Connect(context.Background(), c.GetDSN())
	if err != nil {
		return nil, err
	}

	err = db.Ping(context.Background())
	if err != nil {
		return nil, err
	}

	return db, nil
}

type Container interface {
	Type() []string
	Data() []byte
}

func NewContainer(t []string, d []byte) (*Contain, error) {
	var c Contain
	var err error

	c._type = t
	c.data = d

	return &c, err
}

type Contain struct {
	_type []string
	data  []byte
}

func (c Contain) Type() []string {
	return c._type
}

func (c Contain) Data() []byte {
	return c.data
}

func NewEmailContainer() {

}
