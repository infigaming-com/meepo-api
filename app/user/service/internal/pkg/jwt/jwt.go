package jwt

import (
	"errors"
	"strings"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenPrefix = "Bearer "
)

var (
	ErrInvalidToken  = errors.New("invalid token")
	ErrTokenExpired  = errors.New("token expired")
	ErrInvalidClaims = errors.New("invalid claims")
)

type UserInfo struct {
	UserId string `json:"user_id"`
}

type Claims struct {
	UserInfo `json:"user_info"`
	jwt.RegisteredClaims
}

func CreateToken(userInfo UserInfo, secret string, expiry time.Duration) (string, error) {
	claims := &Claims{
		UserInfo: userInfo,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(expiry)),
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	signedToken, err := token.SignedString([]byte(secret))
	if err != nil {
		return "", err
	}

	return TokenPrefix + signedToken, nil
}

func ParseToken(token string, secret string) (*Claims, error) {
	splitToken := strings.Split(token, " ")
	if len(splitToken) != 2 || splitToken[0] != TokenPrefix {
		return nil, ErrInvalidToken
	}
	token = splitToken[1]

	jwtToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(secret), nil
	})
	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return nil, ErrTokenExpired
		}
		return nil, err
	}
	if !jwtToken.Valid {
		return nil, ErrInvalidToken
	}

	claims, ok := jwtToken.Claims.(*Claims)
	if !ok {
		return nil, ErrInvalidClaims
	}

	return claims, nil
}
