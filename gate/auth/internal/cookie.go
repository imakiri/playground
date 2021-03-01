package internal

import (
	"math/rand"
)

var defaultRandFunc = func(n int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Int63()%int64(len(letterBytes))]
	}
	return string(b)
}

//func NewCookie(resolver core.Resolver, randFunc *func(n int) string) (*Cookie, error) {
//	var c Cookie
//	var err error
//	c.resolver = resolver
//
//	if randFunc == nil {
//		c.randFunc = &defaultRandFunc
//	} else {
//		c.randFunc = randFunc
//	}
//
//	return &c, err
//}
//
//type Cookie struct {
//	resolver core.Resolver
//	id       *ID
//	rand     *Factor_Random
//	randFunc *func(n int) string
//	done     bool
//}
//
//func (c Cookie) GetResolvers() []core.Resolver {
//	return []core.Resolver{c.resolver}
//}
//
//func (c Cookie) GetID() core.ID {
//	return c.id
//}
//
//func (c *Cookie) Identify(factor core.Factor) (bool, error) {
//	if c.done {
//		return false, errors.New("done")
//	}
//
//	switch f := factor.(type) {
//	case ID:
//		if c.id == nil {
//
//		} else {
//			return false, errors.New("id is already identified")
//		}
//	case Factor_Random:
//		if b, err := c.resolver.Enrol(&f, c.id); b {
//			c.rand = &f
//			c.done = true
//			return b, err
//		} else {
//			return b, err
//		}
//	}
//
//	return false, errors.New("factor type mismatch")
//}
//
//func (c *Cookie) Check(factor core.Factor) (bool, error) {
//	if c.done {
//		return false, errors.New("done")
//	}
//
//	switch f := factor.(type) {
//	case Factor_Random:
//		if id, err := c.resolver.Verify(&f); err == nil {
//			*c.id = id
//			c.done = true
//			return true, err
//		} else {
//			return false, err
//		}
//	}
//
//	return false, errors.New("factor type mismatch")
//}
//
//func (c *Cookie) Withdraw(factor core.Factor) (bool, error) {
//	if !c.done {
//		return false, errors.New("must be done")
//	}
//
//}
//
//func (c *Cookie) random(n uint8) string {
//
//}
