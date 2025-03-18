package utils

import (
	"fmt"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)


func ExtractUserIDFromToken(authHeader string) (int) {

	var jwtSecret = []byte("find_my_doc")

	tokenParts := strings.Split(authHeader, " ")

	tokenString := tokenParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return 0
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		userIDFloat, ok := claims["user_id"].(float64)
		if !ok {
			return 0
		}
		return int(userIDFloat)
	}

	return 0
}

func ExtractRoleFromToken(authHeader string) (string) {

	var jwtSecret = []byte("find_my_doc")

	tokenParts := strings.Split(authHeader, " ")

	tokenString := tokenParts[1]

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})

	if err != nil {
		return ""
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		role, ok := claims["role"].(string)
		if !ok {
			return ""
		}
		return role
	}

	return ""
}