package token

import (
	"errors"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type JWTMaker struct {
	secretKey string
}

// CreateToken implements Maker.
func (j *JWTMaker) CreateToken(username string, role string, duration time.Duration) (string, error) {
	payload, err := NewPayload(username, role, duration)
	if err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"id":         payload.Id,
		"username":   payload.Username,
		"role":       payload.Rolename,
		"issued_at":  payload.IssuedAt,
		"expired_at": payload.ExpiredAt,
	})
	return token.SignedString([]byte(j.secretKey))
}

// VerifyToken implements Maker.
func (j *JWTMaker) VerifyToken(token string) (*Payload, error) {
	claim := &Payload{}
	_, err := jwt.ParseWithClaims(token, claim, func(t *jwt.Token) (interface{}, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, errors.New("unexpected signing method")
		}
		return []byte(j.secretKey), nil
	})
	if err != nil {
		return nil, errors.New("could not parse token")
	}
	err = claim.Valid()
	if err != nil {
		return nil, err
	}

	return claim, nil
}

func NewJWTMaker(secretKey string) (Maker, error) {
	return &JWTMaker{secretKey}, nil
}
