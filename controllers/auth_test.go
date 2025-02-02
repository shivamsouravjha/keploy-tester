package controllers

import (
	"fmt"
	"testing"

	"github.com/golang-jwt/jwt/v4"
)

// Test generated using Keploy
func TestGenerateJWT_ValidToken(t *testing.T) {
	username := "testuser"
	token, err := generateJWT(username)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	// Parse the token to verify claims
	parsedToken, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return jwtSecret, nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}

	claims, ok := parsedToken.Claims.(jwt.MapClaims)
	if !ok || !parsedToken.Valid {
		t.Fatalf("Token is invalid or claims are not of type jwt.MapClaims")
	}

	if claims["username"] != username {
		t.Errorf("Expected username claim to be '%s', got '%s'", username, claims["username"])
	}
}

// Test generated using Keploy
func TestGenerateJWT_SpecialCharactersUsername(t *testing.T) {
	username := "user!@#$%^&*()"
	tokenString, err := generateJWT(username)
	if err != nil {
		t.Fatalf("Expected no error, got %v", err)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	if err != nil {
		t.Fatalf("Failed to parse token: %v", err)
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		t.Fatalf("Token is invalid or claims are not jwt.MapClaims")
	}
	if claims["username"] != username {
		t.Errorf("Expected username claim to be '%s', got '%s'", username, claims["username"])
	}
}
