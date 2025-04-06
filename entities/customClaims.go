package entities

import "github.com/golang-jwt/jwt/v5"

type CustomClaims[T any] struct {
	Data T `json:"data"`
	jwt.RegisteredClaims
}
