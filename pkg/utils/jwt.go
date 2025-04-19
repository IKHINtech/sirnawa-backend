package utils

import (
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v5"
)

// GenerateAccessToken membuat access token dengan durasi tertentu
func GenerateAccessToken(data models.User) (string, int64, error) {
	claims := jwt.MapClaims{
		"user_id": data.ID,
		"exp":     time.Now().Add(time.Hour * 24).Unix(), // Berlaku 1 hari
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString([]byte(config.AppConfig.JWT_SECRET))
	if err != nil {
		return "", 0, err
	}
	// Menambahkan waktu expired dalam format epoch (jumlah detik sejak 1970-01-01)
	expiredTime := time.Now().Add(time.Minute * 15).Unix()

	// Kembalikan token dan waktu expired dalam format epoch
	return signedToken, expiredTime, nil
}

// GenerateRefreshToken membuat refresh token dengan durasi tertentu
func GenerateRefreshToken(userID string) (string, error) {
	claims := jwt.MapClaims{
		"user_id": userID,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // Berlaku 7 hari
		"iat":     time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.AppConfig.JWT_SECRET))
}

func SetRefreshTokenCookie(c *fiber.Ctx, refreshString string) {
	// Set Refresh Token as HTTP-only Cookie
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",                    // Cookie name
		Value:    refreshString,                      // Cookie value (refresh token)
		HTTPOnly: true,                               // Makes the cookie inaccessible to JavaScript
		Secure:   false,                              // Only send the cookie over HTTPS (recommended for production)
		Expires:  time.Now().Add(time.Hour * 24 * 7), // Set expiration (7 day)
		SameSite: "None",                             // Helps mitigate CSRF
	})
}
