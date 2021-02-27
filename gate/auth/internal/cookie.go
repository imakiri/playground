package internal

import (
	"github.com/imakiri/playground/core"
	"google.golang.org/grpc/codes"
	"time"
)

func NewCookie(duration time.Duration, storage core.Judge) (*Cookie, error) {
	var cookie = new(Cookie)
	cookie.expirationTime = time.Now().Add(duration).UnixNano()
	cookie.judges = []core.Judge{storage}
	cookie.level = 1
	return cookie, nil
}

type Cookie struct {
	judges         []core.Judge
	currentLevel   int8
	level          int8
	wasVerified    bool
	assertions     core.Assertions
	expirationTime int64
}

func (cookie Cookie) Judges() []core.Judge {
	return cookie.judges
}

func (cookie Cookie) ID() string {
	return cookie.assertions[0].Data().(string)
}

func (cookie Cookie) Level() int8 {
	return cookie.level
}

func (cookie Cookie) IsVerified() bool {
	return cookie.wasVerified && time.Now().UnixNano() < cookie.expirationTime
}

func (cookie Cookie) RegisterAssertion(a core.Assertion, j core.Judge) error {
	var ass core.Assertion
	var err error
	if a.Type() != Type_Assertion_ID {
		return core.StatusCode(codes.InvalidArgument)
	}
	id, ok := a.Data().(string)
	if !ok {
		return core.StatusCode(codes.InvalidArgument)
	}
	if id == "" {
		return core.StatusCode(codes.InvalidArgument)
	}
	cookie.assertions = append(cookie.assertions, a)

	ass, err = cookie.judges[0].AddAssertion(a, nil)
	if err != nil {
		return err
	}

	cookie.assertions = append(cookie.assertions, ass)
	cookie.wasVerified = true
	cookie.currentLevel = 1

	return err
}

func (cookie Cookie) VerifyAssertion(a core.Assertion, j core.Judge) error {
	panic("implement me")
}

func (cookie Cookie) WithdrawAssertion() error {
	panic("implement me")
}
