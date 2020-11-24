package app

import (
	jwt "github.com/dgrijalva/jwt-go"
	"github.com/imakiri/playground/core"
	"github.com/imakiri/playground/data"
	"golang.org/x/crypto/bcrypt"
	_ "golang.org/x/crypto/bcrypt"
	"image"
	"time"
)

const hashCost = 10
const testKey = "testKey"

//
type CheckAuthorization struct {
	Request struct {
		login string
		pass  string
	}
	Response struct {
		re bool
	}
}
type GenerateJWT struct {
	Request struct {
		login string
	}
	Response struct {
		tokenString string
	}
}
type ImageConverter struct {
	picB []byte
	pic  image.Image
}

//
func (e *ImageConverter) Compute() (err error) {
	return nil
}
func (e *CheckAuthorization) Compute() (err error) {
	c := data.NewRequest(data.RequestInternalMainGetUserPassHash{}).(*core.DataInternalMainGetUserPassHash)

	c.Request.Login = e.Request.login
	c.SQL()

	if !c.Package.Status.IsOK() {
		return c.Package.Status
	}

	if bcrypt.CompareHashAndPassword(c.Response.PassHash, []byte(e.Request.pass)) == nil {
		return nil
	} else {
		return NotAuthorizedError{}
	}
} // Returns error if it is not authorized
func (e *GenerateJWT) Compute() (err error) {
	var token *jwt.Token
	var claims jwt.MapClaims

	token = jwt.New(jwt.SigningMethodHS256)
	claims = token.Claims.(jwt.MapClaims)
	claims["authorized"] = true
	claims["user"] = e.Request.login
	claims["exp"] = time.Now().Add(30 * time.Minute).Unix()

	e.Response.tokenString, err = token.SignedString([]byte(testKey))
	if err != nil {
		return InternalServiceError{ERROR(err.Error())}
	}

	return
}
