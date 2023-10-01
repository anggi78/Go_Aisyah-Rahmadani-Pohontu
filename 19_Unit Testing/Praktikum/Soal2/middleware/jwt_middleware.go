package middleware

import (
	"project/constants"
	"time"

	"github.com/dgrijalva/jwt-go"

	echojwt "github.com/labstack/echo-jwt/v4"
	"github.com/labstack/echo/v4"
)

func CreateToken(userID int) (string, error) {
	claims := jwt.MapClaims{}
	claims["userID"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 1).Unix()

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(constants.SECRET_JWT))
}

func JWTMiddleware() echo.MiddlewareFunc {
	return echojwt.WithConfig(echojwt.Config{
		SigningKey:    []byte(constants.SECRET_JWT),
		SigningMethod: "HS256",
	})
}