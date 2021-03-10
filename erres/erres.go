package erres

import (
	"strings"
	"time"
)

type Error string

const (
	E_InvalidArgument      Error = "invalid argument"
	E_TypeMismatch         Error = "type mismatch"
	E_AccessDenied         Error = "access denied"
	E_NotFound             Error = "not found"
	E_AlreadyExist         Error = "already exist"
	E_InternalServiceError Error = "internal service error"
	E_SerializationError   Error = "serialization error"
	E_DeserializationError Error = "deserialization error"
	E_JustError            Error = "error"
)

const error_template = "[(time)] (error message) (trace).(trace):(description):(description)"

var o_Format = "2006-01-02 15:04:05"

func SetDefaultFormat(format string) {
	o_Format = format
}

func Compare(err error, base error) bool {
	switch {
	case err == nil && base == nil:
		return true
	case err != nil && base != nil:
		return strings.Contains(err.Error(), base.Error())
	default:
		return false
	}
}

func (e Error) Error() string {
	var es = string(e)

	es = trimTrace(es)
	es = trimDescription(es)

	return es
}

func (e Error) Time(format string) Error {
	var t string

	if format == "" {
		t = time.Now().Format(o_Format)
	} else {
		t = time.Now().Format(format)
	}
	t = "[" + t + "] "

	return Error(t + string(e))
}

func (e Error) Route(s string) Error {
	var es = string(e)

	if strings.HasSuffix(es, ".") {
		return Error(es + s + ".")
	} else {
		return Error(es + " " + s + ".")
	}
}

func (e Error) Description(s string) Error {
	var es = string(e)

	es = trimTrace(es)

	if strings.HasSuffix(es, ":") {
		return Error(es + s + ":")
	} else {
		return Error(es + " " + s + ":")
	}
}

func trimTrace(s string) string {
	if strings.HasSuffix(s, ".") {
		return strings.TrimSuffix(s, ".")
	} else {
		return s
	}
}

func trimDescription(s string) string {
	if strings.HasSuffix(s, ":") {
		return strings.TrimSuffix(s, ":")
	} else {
		return s
	}
}
