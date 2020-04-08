package util

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"

    "laojgo/app"
)

var jwtSigningKey = []byte("2i#L7Hym@2#O1")

type Claims struct {
    UserId string
    jwt.StandardClaims
}

func GenerateJwt(user_id string) (ss string, err error) {
    c := Claims{
        user_id,
        jwt.StandardClaims {
            ExpiresAt : time.Now().Add(3 * time.Hour).Unix(),
            Issuer : app.Cfg.Name,
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    ss, err = token.SignedString(jwtSigningKey)

    return
}

func ParseJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return jwtSigningKey, nil
    })

    if token != nil {
        if c, ok := token.Claims.(*Claims); ok && token.Valid {
        	return c, nil
        } 
    }

    return nil, err
}