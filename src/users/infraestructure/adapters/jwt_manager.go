package adapters

import (
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v4"
)

type JWTManager struct {
    secretKey string
}

func NewJWTManager(secretKey string) *JWTManager {
    return &JWTManager{secretKey: secretKey}
}

func (m *JWTManager) GenerateToken(username string, role string) (string, error) {
    claims := jwt.MapClaims{ 
        "username": username,
        "role":     role, 
        "exp":      time.Now().Add(time.Hour * 1).Unix(),
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    return token.SignedString([]byte(m.secretKey))
}

func (m *JWTManager) ValidateToken(tokenString string) (bool, map[string]interface{}, error) {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
            return nil, fmt.Errorf("unexpected signing method")
        }
        return []byte(m.secretKey), nil
    })

    if err != nil || !token.Valid {
        return false, nil, err
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if !ok {
        return false, nil, fmt.Errorf("invalid token claims")
    }

    return true, claims, nil
}