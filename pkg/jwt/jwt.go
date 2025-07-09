package jwt

import (
	"strings"
	"time"

	"github.com/go-kratos/kratos/v2/errors"
	"github.com/golang-jwt/jwt/v5"
)

const (
	TokenPrefix = "Bearer "
)

var (
	ErrInvalidToken  = errors.New(400, "BAD_REQUEST", "invalid token")
	ErrTokenExpired  = errors.New(401, "UNAUTHORIZED", "token expired")
	ErrInvalidClaims = errors.New(400, "BAD_REQUEST", "invalid claims")
)

type UserInfo struct {
	UserId             int64  `json:"userId"`
	OperatorId         int64  `json:"operatorId"`
	CompanyOperatorId  int64  `json:"companyOperatorId"`
	RetailerOperatorId int64  `json:"retailerOperatorId"`
	SystemOperatorId   int64  `json:"systemOperatorId"`
	RealOperatorId     int64  `json:"realOperatorId,omitempty"`
	OperatorType       string `json:"operatorType,omitempty"`
	RoleId             int64  `json:"roleId"`
}

type Claims struct {
	UserInfo `json:"userInfo"`
	jwt.RegisteredClaims
}

func CreateToken(userInfo UserInfo, secret string, expiry time.Duration) (string, error) {
	claims := Claims{
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

	return signedToken, nil
}

func ParseToken(token string, secret string) (string, *Claims, error) {
	// works for token with and without prefix
	token = strings.TrimPrefix(token, TokenPrefix)

	// Parse the token
	parsedToken, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		// Validate the signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, ErrInvalidToken
		}
		return []byte(secret), nil
	})

	if err != nil {
		if errors.Is(err, jwt.ErrTokenExpired) {
			return "", nil, ErrTokenExpired
		}
		return "", nil, err
	}

	if !parsedToken.Valid {
		return "", nil, ErrInvalidToken
	}

	claims, ok := parsedToken.Claims.(*Claims)
	if !ok {
		return "", nil, ErrInvalidClaims
	}

	return token, claims, nil
}

func GetBearerToken(token string) string {
	if strings.HasPrefix(token, TokenPrefix) {
		return token
	}
	return TokenPrefix + token
}

func FromJwtClaims[T any](claims jwt.Claims, key string) T {
	var zero T
	claimsMap, ok := claims.(jwt.MapClaims)
	if !ok {
		return zero
	}

	return FromJwtClaimsMap[T](claimsMap, key)
}

func FromJwtClaimsMap[T any](claims jwt.MapClaims, key string) T {
	var zero T

	v, ok := claims[key]
	if !ok {
		return zero
	}

	switch any(zero).(type) {
	case string:
		if value, ok := v.(string); ok {
			return any(value).(T)
		}
	case bool:
		if value, ok := v.(bool); ok {
			return any(value).(T)
		}
	}
	return zero
}
