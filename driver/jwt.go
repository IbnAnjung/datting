package driver

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Jwt struct {
	issuer         string
	signKey        string
	expireDuration int64
}

type jwtClaims struct {
	jwt.RegisteredClaims
	claims interface{}
}

func NewJwt(
	issuer string,
	signKey string,
	expireDuration int64,
) Jwt {
	return Jwt{
		issuer:         issuer,
		signKey:        signKey,
		expireDuration: expireDuration,
	}
}

func (j Jwt) GenerateToken(claims interface{}) (signToken string, err error) {
	jwtclaims := jwtClaims{
		jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Duration(j.expireDuration) * time.Minute)),
			IssuedAt:  jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Issuer:    j.issuer,
		},
		claims,
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwtclaims)
	signToken, err = token.SignedString([]byte(j.signKey))

	return
}

func (j Jwt) ParseToken(tokenString string) (claims interface{}, err error) {
	jwtClaims := jwtClaims{}
	token, err := jwt.ParseWithClaims(tokenString, &jwtClaims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(j.signKey), nil
	})

	if token.Valid {
		return jwtClaims.claims, err
	}

	return nil, err
}
