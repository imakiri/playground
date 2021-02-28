package internal

//func NewCookie(duration time.Duration, judge core.Judge) (*Cookie, error) {
//	var cookie = new(Cookie)
//	cookie.expirationTime = time.Now().Add(duration).UnixNano()
//	cookie.judges = []core.Judge{judge}
//	return cookie, nil
//}
//
//type Cookie struct {
//	judges         core.Judges
//	judgesExp      map[int]core.Judge
//	assertions     core.Assertions
//	assertionsExp  map[int]core.Assertion
//	currentLevel   int
//	wasVerified    bool
//	expirationTime int64
//}
//
//func (cookie Cookie) Judges() core.Judges {
//	return cookie.judges
//}
//
//func (cookie Cookie) Assertions() core.Assertions {
//	return cookie.assertions[1:]
//}
//
//func (cookie Cookie) ID() core.Assertion {
//	if cookie.wasVerified {
//		return cookie.assertions[0]
//	} else {
//		return nil
//	}
//}
//
//func (cookie Cookie) Level() int {
//	return len(cookie.assertions)
//}
//
//func (cookie Cookie) IsVerified() bool {
//	return cookie.wasVerified && time.Now().UnixNano() < cookie.expirationTime
//}
//
//func (cookie Cookie) ExtendWith(c core.Credentials) error {
//	return nil
//}
//
//func (cookie Cookie) RegisterAssertion(a core.Assertion) error {
//	var ass core.Assertion
//	var err error
//
//	switch a := a.(type) {
//	case Assertion_ID:
//		if cookie.Level() == 0 {
//			cookie.assertions = append(cookie.assertions, a)
//			return err
//		} else {
//			return errors.New("ID already registered")
//		}
//	case Assertion_ExpirationTime:
//		return err
//	case Assertion_Empty:
//		if cookie.Level() == 2 && !cookie.IsVerified() {
//			ass, err = cookie.judges[0].AddAssertion(a, cookie)
//			if err != nil {
//				return err
//			}
//
//			cookie.assertions = append(cookie.assertions, ass)
//			cookie.wasVerified = true
//			return err
//		} else {
//			return errors.New("incompatible Assertion")
//		}
//	default:
//		return errors.New("incompatible Assertion")
//	}
//}
//
//func (cookie Cookie) VerifyAssertion(a core.Assertion) error {
//	panic("implement me")
//}
//
//func (cookie Cookie) WithdrawAssertion() error {
//	panic("implement me")
//}
