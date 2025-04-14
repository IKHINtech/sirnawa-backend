package handlers

import (
	"errors"
	"net/mail"
	"time"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"github.com/IKHINtech/sirnawa-backend/internal/database"
	"github.com/IKHINtech/sirnawa-backend/internal/dto"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/request"
	"github.com/IKHINtech/sirnawa-backend/internal/dto/response"
	"github.com/IKHINtech/sirnawa-backend/internal/models"
	"github.com/IKHINtech/sirnawa-backend/pkg/helpers"
	"github.com/IKHINtech/sirnawa-backend/pkg/utils"
	"github.com/IKHINtech/sirnawa-backend/pkg/validators"
	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

// CheckPasswordHash compare password with hash
func CheckPasswordHash(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil
}

func getUserByUserID(e string) (*models.User, error) {
	db := database.DB
	var user models.User
	if err := db.Where("id = ?", e).First(&user).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &user, nil
}

func isEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// Login get user and password godoc
// @Summary Login
// @Description Login
// @Tags Auth
// @Accept json
// @Produce json
// @Param login body request.LoginInput true "Login"
// @Success 200 {object} utils.ResponseData
// @Router /auth/login [post]
func Login(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}

	input := new(request.LoginInput)

	var userData response.UserResponse
	if err := c.BodyParser(&input); err != nil {
		return h.BadRequest(c, []string{"Error on login request", err.Error()})
	}

	identity := input.Email
	pass := input.Password
	usermodels, err := new(models.User), *new(error)

	if isEmail(identity) {
		usermodels, err = helpers.GetUserByEmail(identity)
	}

	if usermodels == nil {
		return h.Forbidden(c, []string{"User not found"})
	}

	if err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Internal server error"})
	}

	userData = response.UserResponse{
		ID:    usermodels.ID,
		Name:  usermodels.Name,
		Email: usermodels.Email,
		Role:  usermodels.Role,
	}

	if !CheckPasswordHash(pass, usermodels.Password) {
		return h.Forbidden(c, []string{"Invalid identity or password"})
	}

	// Create Access Token
	accessString, activeUntil, err := utils.GenerateAccessToken(*usermodels)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	accessToken := dto.Token{
		Token:     accessString,
		ExpiresIn: activeUntil,
	}

	// Create Refresh Token
	refreshString, err := utils.GenerateRefreshToken(userData.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create refresh token", err.Error()})
	}

	utils.SetRefreshTokenCookie(c, refreshString)

	response := dto.ResponseLogin{
		User:        userData,
		AccessToken: accessToken,
	}

	return h.Ok(c, response, "Success login", nil)
}

// GetUser get user godoc
// @Summary Get User
// @Description Get User
// @Tags Auth
// @Accept json
// @Produce json
// @Security Bearer
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/me [get]
func Me(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}
	userID := c.Locals("user_id").(string)
	userData, err := getUserByUserID(userID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on get user data", err.Error()})
	}
	if userData == nil {
		return h.NotFound(c, []string{"User not found"})
	}
	return h.Ok(c, userData, "Success get user data", nil)
}

// Register handles user registration godoc
// @Summary Register
// @Description Register
// @Tags Auth
// @Accept json
// @Produce json
// @Param register body dto.RegisterInput true "Register"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/register [post]
func Register(c *fiber.Ctx) error {
	input := new(dto.RegisterInput)
	h := &utils.ResponseHandler{}
	if err := c.BodyParser(&input); err != nil {
		return h.BadRequest(c, []string{err.Error(), "Invalid Input"})
	}

	// Validate input fields
	err := validators.ValidateStruct(input)
	if err != nil {
		return h.BadRequest(c, []string{err.Error(), "Invalid Input"})
	}
	if err := c.BodyParser(&input); err != nil {
		return h.BadRequest(c, []string{err.Error(), "Invalid Input"})
	}

	// Check if email is valid
	if !isEmail(input.Email) {
		return h.BadRequest(c, []string{"Email is not valid"})
	}

	// Check if email already exists
	existingUserByEmail, _ := helpers.GetUserByEmail(input.Email)
	if existingUserByEmail != nil {
		return h.BadRequest(c, []string{"Email already exists"})
	}

	// Hash the password
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(input.Password), bcrypt.DefaultCost)
	if err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Failed to hash password"})
	}

	// Create new user
	user := models.User{
		Email:    input.Email,
		Password: string(hashedPassword),
	}

	// Save user to database
	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Failed to create user"})
	}
	return h.Created(c, user, "User registered successfully")
}

// RefreshToken handles refresh token request godoc
// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/refresh-token [get]
func RefreshToken(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}

	// Ambil refresh token dari cookie
	refreshToken := c.Cookies("refresh_token") // Ganti "refresh_token" dengan nama cookie yang sesuai

	// Pastikan token ada
	if refreshToken == "" {
		return h.Unauthorized(c, []string{"Refresh token is missing"})
	}
	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (interface{}, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.CFG.JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		return h.Forbidden(c, []string{"Invalid or expired refresh token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	usermodels, err := getUserByUserID(userID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on get user data", err.Error()})
	}

	if usermodels == nil {
		return h.Unauthorized(c, []string{"User not found"})
	}

	userData := response.UserResponse{
		ID:    usermodels.ID,
		Name:  usermodels.Name,
		Email: usermodels.Email,
		Role:  usermodels.Role,
	}
	// Generate new access token
	accessString, accessTime, err := utils.GenerateAccessToken(*usermodels)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	responseL := dto.ResponseLogin{
		User: userData,
		AccessToken: dto.Token{
			Token:     accessString,
			ExpiresIn: accessTime,
		},
	}

	return h.Ok(c, responseL, "Token refreshed successfully", nil)
}

// Logout handles logout request godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @AAccept json
// @Produce json
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/logout [delete]
func Logout(c *fiber.Ctx) error {
	// Menghapus refresh token dari cookie HTTP-only
	c.Cookie(&fiber.Cookie{
		Name:     "refresh_token",            // Nama cookie yang disetel saat login
		Value:    "",                         // Hapus nilai cookie
		Expires:  time.Now().Add(-time.Hour), // Set waktu kedaluwarsa di masa lalu untuk menghapus cookie
		HTTPOnly: true,                       // Pastikan cookie hanya dapat diakses oleh server
		Secure:   true,                       // Hanya kirim cookie di koneksi HTTPS
		SameSite: "lax",                      // Perlindungan CSRF
	})

	// Kembalikan respons sukses tanpa token
	return c.JSON(fiber.Map{
		"message": "Successfully logged out",
	})
}
