package utils

import (
	"social-media-app/pkg/config"
	"time"

	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
)

type JwtWrapper struct {
	SecretKey string
	Issuer    string
	Expire    int
}

// JwtClaim adds email as a claim to the token
type JwtClaim struct {
	UserId    uuid.UUID
	IpAddress string
	jwt.RegisteredClaims
}

func (j *JwtWrapper) ValidateToken(tokenString string) bool {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtClaim{},
		func(token *jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return false
	}

	claims, _ := token.Claims.(*JwtClaim)

	if claims.ExpiresAt.Local().Unix() < time.Now().Local().Unix() {
		return false
	}

	return token.Valid
}

/*
func (j JwtWrapper) ParseToken(tokenString string) (claims JwtClaim, err error) {
	token, err := jwt.ParseWithClaims(
		tokenString,
		&JwtClaim{},
		func(token jwt.Token) (interface{}, error) {
			return []byte(j.SecretKey), nil
		},
	)
	if err != nil {
		return nil, err
	}
	claims, ok := token.Claims.(JwtClaim)
	if !ok {
		return nil, fmt.Errorf("claims not JwtClaim")
	}

	return claims, nil
}*/

func GenerateJwt(userID uuid.UUID, key, ip string) (string, error) {
	claims := &JwtClaim{
		UserId:    userID,
		IpAddress: ip,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(time.Hour * time.Duration(config.ReadValue().JwtExpTime))),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenSigned, err := token.SignedString([]byte(key))
	if err != nil {
		return tokenSigned, err
	}
	return tokenSigned, nil
}
