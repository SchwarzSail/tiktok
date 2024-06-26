package utils

import (
	"github.com/golang-jwt/jwt"
	"tiktok/pkg/constants"
	"time"
)

type Claims struct {
	ID       string
	UserName string
	jwt.StandardClaims
}

// GenerateAccessToken 签发access_token
func GenerateAccessToken(uid string, username string) (accessToken string, err error) {
	claims := &Claims{
		ID:       uid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			//TODO: 记得修改回去
			ExpiresAt: (time.Now().Add(24 * 30 * time.Hour)).Unix(),
			Issuer:    "tiktok",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(constants.JwtSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}

// GenerateRefreshToken 签发refresh_token
func GenerateRefreshToken(uid string, username string) (accessToken string, err error) {
	claims := &Claims{
		ID:       uid,
		UserName: username,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: (time.Now().Add(7 * 24 * time.Hour)).Unix(),
			Issuer:    "tiktok",
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	accessToken, err = token.SignedString([]byte(constants.JwtSecret))
	if err != nil {
		return "", err
	}
	return accessToken, nil
}
func ParseToken(token string) (*Claims, bool, error) {
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(constants.JwtSecret), nil
	})
	if tokenClaims != nil {
		if claims, ok := tokenClaims.Claims.(*Claims); ok && tokenClaims.Valid {
			return claims, true, nil
		}
	}
	return nil, false, err
}
