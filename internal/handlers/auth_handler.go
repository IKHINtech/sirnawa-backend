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
	"github.com/IKHINtech/sirnawa-backend/internal/middleware"
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

func getValidVerificationCode(userID, code string) (*models.UserVerification, error) {
	var verification models.UserVerification
	err := database.DB.Where("user_id = ? AND code = ? AND is_used = false AND expires_at > ?", userID, code, time.Now()).
		First(&verification).Error
	if err != nil {
		return nil, err
	}
	return &verification, nil
}

func markVerificationCodeUsed(id string) error {
	return database.DB.Model(&models.UserVerification{}).Where("id = ?", id).
		Update("is_used", true).Error
}

func getOnlyUserByUserID(e string) (*models.User, error) {
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

func createUserVerification(userID, code string, expiredAt time.Time) error {
	db := database.DB
	return db.Create(&models.UserVerification{
		UserID:    userID,
		Code:      code,
		ExpiresAt: expiredAt,
	}).Error
}

func getActiveVerificationCode(userID string) (*models.UserVerification, error) {
	db := database.DB
	var userVerification models.UserVerification
	if err := db.Where("user_id = ? AND is_used = false", userID).First(&userVerification).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &userVerification, nil
}

func activateUser(userID string) error {
	return database.DB.Model(&models.User{}).Where("id = ?", userID).
		Update("is_active", true).Error
}

func getUserByUserID(e string) (*models.User, error) {
	db := database.DB.
		Preload("Resident").
		Preload("UserRTs").
		Preload("UserRTs.Rt").
		Preload("Resident.ResidentHouses").
		Preload("Resident.ResidentHouses.House").
		Preload("Resident.ResidentHouses.House.Rw").
		Preload("Resident.ResidentHouses.House.Rt").
		Preload("Resident.ResidentHouses.House.Block").
		Preload("Resident.ResidentHouses.House.HousingArea")

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

	if !usermodels.IsActive {
		return h.Forbidden(c, []string{"User is not verified"})
	}

	if !CheckPasswordHash(pass, usermodels.Password) {
		return h.Forbidden(c, []string{"Invalid identity or password"})
	}

	userData = response.UserResponse{
		ID:    usermodels.ID,
		Email: usermodels.Email,
		Role:  usermodels.Role.ToString(),
	}

	// Create Access Token
	accessString, activeUntil, err := utils.GenerateAccessToken(*usermodels)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	// Create Refresh Token
	refreshString, err := utils.GenerateRefreshToken(userData.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create refresh token", err.Error()})
	}

	utils.SetRefreshTokenCookie(c, refreshString)

	accessToken := dto.Token{
		Token:        accessString,
		RefreshToken: refreshString,
		ExpiresIn:    activeUntil,
	}

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
	userResponse := response.UserToResponse(userData)
	return h.Ok(c, userResponse, "Success get user data", nil)
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
		Email:    &input.Email,
		Password: string(hashedPassword),
		Role:     models.RoleWarga,
	}

	// Save user to database
	db := database.DB
	if err := db.Create(&user).Error; err != nil {
		return h.InternalServerError(c, []string{err.Error(), "Failed to create user"})
	}

	respUser := response.UserResponse{
		ID:    user.ID,
		Email: user.Email,
		Role:  string(user.Role),
	}
	return h.Created(c, respUser, "User registered successfully")
}

// RefreshToken handles refresh token request godoc
// @Summary Refresh Token
// @Description Refresh Token
// @Tags Auth
// @Accept json
// @Produce json
// @Param data body request.RefreshTokenRequest true "Refresh Token"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/refresh-token [post]
func RefreshToken(c *fiber.Ctx) error {
	h := &utils.ResponseHandler{}
	var refreshToken string

	// Ambil refresh token dari cookie
	refreshTokenCookies := c.Cookies("refresh_token") // Ganti "refresh_token" dengan nama cookie yang sesuai

	// Pastikan token ada
	if refreshTokenCookies == "" {

		var req request.RefreshTokenRequest
		if err := c.BodyParser(&req); err != nil {
			return h.BadRequest(c, []string{"Invalid request body", err.Error()})
		}

		middleware.ValidateRequest(req)
		if req.RefreshToken == "" {
			return h.Unauthorized(c, []string{"Refresh token is missing"})
		}

		refreshToken = req.RefreshToken
	} else {
		refreshToken = refreshTokenCookies
	}
	// Parse and validate the refresh token
	token, err := jwt.Parse(refreshToken, func(token *jwt.Token) (any, error) {
		// Validate signing method
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fiber.ErrUnauthorized
		}
		return []byte(config.AppConfig.JWT_SECRET), nil
	})

	if err != nil || !token.Valid {
		return h.Forbidden(c, []string{"Invalid or expired refresh token"})
	}

	claims := token.Claims.(jwt.MapClaims)
	userID := claims["user_id"].(string)

	usermodels, err := getOnlyUserByUserID(userID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on get user data", err.Error()})
	}

	if usermodels == nil {
		return h.Unauthorized(c, []string{"User not found"})
	}

	userData := response.UserResponse{
		ID:    usermodels.ID,
		Email: usermodels.Email,
		Role:  usermodels.Role.ToString(),
	}
	// Generate new access token
	accessString, accessTime, err := utils.GenerateAccessToken(*usermodels)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	responseL := dto.ResponseLogin{
		User: userData,
		AccessToken: dto.Token{
			Token:        accessString,
			ExpiresIn:    accessTime,
			RefreshToken: refreshToken,
		},
	}

	return h.Ok(c, responseL, "Token refreshed successfully", nil)
}

// Logout handles logout request godoc
// @Summary Logout
// @Description Logout
// @Tags Auth
// @Accept json
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

// Forgor Password handles godoc
// @Summary Send Forgot Password Verification
// @Description Send Forgot Password Verification
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.VerifyCodeRequest true "Forgot Password Request"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/forgot-password-verification [post]
func ForgotPassword(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}

	var req request.VerifyCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return h.BadRequest(c, []string{"Invalid request body", err.Error()})
	}

	middleware.ValidateRequest(req)

	user, err := helpers.GetUserByEmail(req.Email)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to get User", err.Error()})
	}

	if user == nil {
		return h.NotFound(c, []string{"User not found"})
	}

	code := utils.GenerateVerificationCode(6) // e.g., "348112"
	expiresAt := time.Now().Add(10 * time.Minute)

	existingCode, err := getActiveVerificationCode(user.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error checking existing verification", err.Error()})
	}
	if existingCode != nil {
		return h.BadRequest(c, []string{"A verification code was already sent. Please wait until it expires."})
	}

	body := utils.GenerateForgotPasswordOTPEmailBody(code)
	err = utils.SendEmail(utils.MailRequest{
		To:      *user.Email,
		Subject: "Forgot Password OTP",
		Body:    body,
	})
	if err != nil {
		return h.BadRequest(c, []string{err.Error(), "Error on send verification email"})
	}

	err = createUserVerification(user.ID, code, expiresAt)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on create verification code", err.Error()})
	}
	return h.Ok(c, nil, "Verification email sent", nil)
}

// Logout handles SendEmailVerification godoc
// @Summary Send Email Verification
// @Description Send Email Verification
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.VerifyCodeRequest true "Verification Code Request"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/send-email-verification [post]
func SendEmailVerification(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}

	var req request.VerifyCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return h.BadRequest(c, []string{"Invalid request body", err.Error()})
	}

	middleware.ValidateRequest(req)

	user, err := helpers.GetUserByEmail(req.Email)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to get User", err.Error()})
	}

	if user == nil {
		return h.NotFound(c, []string{"User not found"})
	}
	if user.IsActive {
		return h.BadRequest(c, []string{"User already activated"})
	}
	if user.Email == nil || *user.Email == "" {
		return h.BadRequest(c, []string{"User does not have a valid email address"})
	}
	code := utils.GenerateVerificationCode(6) // e.g., "348112"
	expiresAt := time.Now().Add(10 * time.Minute)

	existingCode, err := getActiveVerificationCode(user.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Error checking existing verification", err.Error()})
	}
	if existingCode != nil {
		return h.BadRequest(c, []string{"A verification code was already sent. Please wait until it expires."})
	}

	// Kirim email
	body := utils.GenerateVerificationEmailBody(code)
	err = utils.SendEmail(utils.MailRequest{
		To:      *user.Email,
		Subject: "Kode Aktivasi Akun Anda",
		Body:    body,
	})
	if err != nil {
		return h.BadRequest(c, []string{err.Error(), "Error on send verification email"})
	}

	err = createUserVerification(user.ID, code, expiresAt)
	if err != nil {
		return h.InternalServerError(c, []string{"Error on create verification code", err.Error()})
	}
	return h.Ok(c, nil, "Verification email sent", nil)
}

// Forgor Password Verify Code handles godoc
// @Summary Verify Forgot Password Verification
// @Description Verify Forgot Password Verification
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.VerifyEmailCodeRequest true "Forgot Password Code Vefification"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/send-forgot-password-verification [post]
func VerifyForgotPasswordCode(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}

	var req request.VerifyEmailCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return h.BadRequest(c, []string{"Invalid request body", err.Error()})
	}

	middleware.ValidateRequest(req)

	user, err := helpers.GetUserByEmail(req.Email)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to get User", err.Error()})
	}

	verification, err := getValidVerificationCode(user.ID, req.Code)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to query verification code", err.Error()})
	}

	if verification == nil {
		return h.BadRequest(c, []string{"Invalid or expired verification code"})
	}

	// Tandai kode sebagai digunakan
	err = markVerificationCodeUsed(verification.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to mark verification code as used", err.Error()})
	}

	// Create Access Token
	accessString, activeUntil, err := utils.GenerateAccessToken(*user)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to create access token", err.Error()})
	}

	accessToken := dto.Token{
		Token:     accessString,
		ExpiresIn: activeUntil,
	}

	return h.Ok(c, accessToken, "Verification code used", nil)
}

// VerifyEmailCode handles verification of the email code
// @Summary Verify Email Code
// @Description Verifies the email code and activates the user
// @Tags Auth
// @Accept json
// @Produce json
// @Param request body request.VerifyEmailCodeRequest true "Verification Code"
// @Success 200 {object} utils.ResponseData
// @Failure 400 {object} utils.ErrorResponse
// @Router /auth/verify-email-code [post]
func VerifyEmailCode(c *fiber.Ctx) error {
	h := utils.ResponseHandler{}

	var req request.VerifyEmailCodeRequest
	if err := c.BodyParser(&req); err != nil {
		return h.BadRequest(c, []string{"Invalid request body", err.Error()})
	}

	middleware.ValidateRequest(req)

	user, err := helpers.GetUserByEmail(req.Email)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to get User", err.Error()})
	}

	if user != nil && user.IsActive {
		return h.BadRequest(c, []string{"User is already active"})
	}

	verification, err := getValidVerificationCode(user.ID, req.Code)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to query verification code", err.Error()})
	}

	if verification == nil {
		return h.BadRequest(c, []string{"Invalid or expired verification code"})
	}

	// Tandai kode sebagai digunakan
	err = markVerificationCodeUsed(verification.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to mark verification code as used", err.Error()})
	}

	// Update user: set is_active = true
	err = activateUser(user.ID)
	if err != nil {
		return h.InternalServerError(c, []string{"Failed to activate user", err.Error()})
	}

	return h.Ok(c, nil, "Email verification successful", nil)
}
