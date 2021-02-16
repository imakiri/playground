package core

// Status

type Status string

func (e Status) Error() string {
	return string(e)
}

const Status_OK Status = "OK"
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

// Logger

type ActionID uint64
type Log struct {
	FuncID FunctionID
	Status Status
	Start  int64
	End    int64
}
type Trace struct {
	ActionID
	Logs []Log
}

func (e *Trace) AddLog(fid FunctionID, s Status, start int64, end int64) {
	e.Logs = append(e.Logs, Log{fid, s, start, end})
}

type Resume struct {
	ActionID
	Status
}

//
