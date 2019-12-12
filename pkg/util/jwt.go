package util

import (
    "time"

    jwt "github.com/dgrijalva/jwt-go"

    "laojgo/config"
)

var JwtSigningKey = []byte(config.App.JwtSigningKey)

type Claims struct {
    UserId string `json:"user_id"`
    jwt.StandardClaims
}

func GenerateJwt(user_id string) (ss string, err error) {
    c := Claims{
        user_id,
        jwt.StandardClaims {
            ExpiresAt : time.Now().Add(3 * time.Hour).Unix(),
            Issuer : "gin-blog",
        },
    }
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)
    ss, err = token.SignedString(JwtSigningKey)

    return
}

func ParseJwt(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token *jwt.Token) (interface{}, error) {
        return JwtSigningKey, nil
    })

    if token != nil {
        if c, ok := token.Claims.(*Claims); ok && token.Valid {
        	return c, nil
        } 
    }

    return nil, err
}