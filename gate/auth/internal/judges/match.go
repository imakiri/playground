package judges

import (
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/gate/auth/internal"
	"github.com/imakiri/playground/gate/auth/internal/storage"
	"google.golang.org/grpc/codes"
	"math/rand"
)

var random = func(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

func NewCookieJudge() (*CookieJudge, error) {
	var cj CookieJudge
	cj.randFunc = random
	cj.storage = new(storage.MemStorage)

	return &cj, nil
}

type CookieJudge struct {
	randFunc func(n int) string
	storage  core.Storage
}

func (cj *CookieJudge) AddAssertion(id core.Assertion, c core.Credentials) (core.Assertion, error) {
	switch _ := id.(type) {
	case internal.Assertion_ID:
		var r internal.Assertion_Rand
		var err error
		r = internal.Assertion_Rand(cj.randFunc(60))

		err = cj.storage.Write(r)
		if err != nil {
			return nil, err
		}

		return r, nil
	default:
		return nil, core.StatusCode(codes.InvalidArgument)
	}
}

func (cj *CookieJudge) CheckAssertion(assertion core.Assertion, c core.Credentials) (core.Assertion, error) {
	panic("implement me")
}

func (cj *CookieJudge) WithdrawAssertion(ass core.Assertion, c core.Credentials) error {
	panic("implement me")
}
