package httputils

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"github.com/yobol/yobol/config"
)

var (
	ErrTokenInvalid = errors.New("couldn't handle the token")
)

func NewJWT() *JWT {
	return &JWT{
		secret: []byte(config.C.Server.Auth.JWTSecret),
	}
}

type JWT struct {
	secret []byte
}

func (j *JWT) CreateToken(baseClaims *BaseClaims) (string, error) {
	claims := &CustomizedClaims{
		BaseClaims: *baseClaims,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(config.C.Server.Auth.TokenMaxAge).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(j.secret)
}

func (j *JWT) ParseToken(tokenString string) (*CustomizedClaims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &CustomizedClaims{}, func(token *jwt.Token) (i interface{}, e error) {
		return j.secret, nil
	})
	if err != nil {
		if ve, ok := err.(*jwt.ValidationError); ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				return nil, jwt.ErrTokenMalformed
			} else if ve.Errors&jwt.ValidationErrorExpired != 0 {
				return nil, jwt.ErrTokenExpired
			} else if ve.Errors&jwt.ValidationErrorNotValidYet != 0 {
				return nil, jwt.ErrTokenNotValidYet
			} else {
				return nil, ErrTokenInvalid
			}
		}
	}
	if token != nil {
		if claims, ok := token.Claims.(*CustomizedClaims); ok && token.Valid {
			return claims, nil
		}
		return nil, ErrTokenInvalid
	} else {
		return nil, ErrTokenInvalid
	}
}

type BaseClaims struct {
	Username string `json:"username"`
}

type CustomizedClaims struct {
	jwt.StandardClaims

	BaseClaims
}

func (c *CustomizedClaims) GetUsername() string {
	return c.Username
}
