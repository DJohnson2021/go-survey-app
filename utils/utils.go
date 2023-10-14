package utils

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load .env file
	if err := godotenv.Load("../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}
}

func generateRandomBytes(n int) ([]byte, error) {
	b := make([]byte, n)
	_, err := rand.Read(b)
	if err != nil {
		return nil, err
	}
	return b, nil
}

func generateJWTKey() (string, error) {
	key, err := generateRandomBytes(32) // 32 bytes for HS256, you can use 64 bytes for HS512
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(key), nil
}

func GetJWTSecret() ([]byte, error) {
	jwtSecret := os.Getenv("JWT_SECRET")
	if jwtSecret == "" {
		return nil, errors.New("missing JWT_SECRET in environment variables")
	}

	return []byte(jwtSecret), nil
}

func GetAdminNamesAndEmails() ([]string, []string, error) {
	adminNames := strings.Split(os.Getenv("ADMIN_NAMES"), ",")
	adminEmails := strings.Split(os.Getenv("ADMIN_EMAILS"), ",")
	if len(adminNames) == 0 || len(adminEmails) == 0 {
		return nil, nil, errors.New("missing ADMIN_NAMES or ADMIN_EMAILS in environment variables")
	}

	return adminNames, adminEmails, nil
}
