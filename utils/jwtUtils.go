package utils

import (
	"log"
	"net/http"
	"os"
	"time"

	"github.com/Qushai121/topaz-be/dto"
	"github.com/Qushai121/topaz-be/entities"
	"github.com/golang-jwt/jwt/v5"
)

var (
	AccessTokenKey  = os.Getenv("ACCESS_TOKEN_KEY")
	RefreshTokenKey = os.Getenv("REFRESH_TOKEN_KEY")
)

func GenerateToken[T any](data T, duration time.Duration, secretKey string) (*string, *dto.ErrorDto[any]) {
	exp := time.Now().Add(duration)

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, entities.CustomClaims[any]{
		Data: data,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(exp),
		},
	})

	signedToken, errSigned := token.SignedString([]byte(secretKey))
	
	if errSigned != nil {
		log.Printf("Signed Token %v", errSigned.Error())
		return nil, dto.NewErrorDto[any](errSigned.Error(), http.StatusInternalServerError, nil)
	}

	return &signedToken, nil
}

func EncodeToken[T any](tokenStr string, secretKey string) (*entities.CustomClaims[T], *dto.ErrorDto[any]) {
	var claims = entities.CustomClaims[T]{}

	_, err := jwt.ParseWithClaims(tokenStr, &claims, func(t *jwt.Token) (interface{}, error) {
		return []byte(secretKey), nil
	}, jwt.WithValidMethods([]string{jwt.SigningMethodHS256.Alg()}))

	if err != nil {
		log.Printf("Check Token %v", err.Error())
		return nil, dto.NewErrorDto[any](err.Error(), http.StatusUnauthorized, nil)
	}


	return &claims, nil
}
