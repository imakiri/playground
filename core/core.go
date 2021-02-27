package core

import (
	"context"
	"github.com/imakiri/playground/admin/cfg"
	"github.com/jackc/pgx/v4"
	"google.golang.org/grpc/codes"
)

// Status

type StatusCode codes.Code

func (e StatusCode) Error() string {
	return string(e)
}

type Status string

func (e Status) Error() string {
	return string(e)
}

const Status_OK Status = "OK"
const Status_InvalidDSN Status = "InvalidDSN"
const Status_InvalidArgument Status = "InvalidArgument"
const Status_AccessDenied Status = "AccessDenied"
const Status_NotFound Status = "NotFound"
const Status_InternalServiceError Status = "InternalServiceError"
const Status_SerializationError Status = "SerializationError"
const Status_UnknownError Status = "UnknownError"
const Status_AlreadyExist Status = "AlreadyExist"

//

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
	Status Status
}

func Connect(c *cfg.Data) (*pgx.Conn, error) {
	var db *pgx.Conn
	var err error

	if c.GetDSN() == "" {
		return nil, Status_InvalidDSN
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
