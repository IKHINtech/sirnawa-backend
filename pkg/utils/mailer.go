package utils

import (
	"fmt"
	"strconv"

	"github.com/IKHINtech/sirnawa-backend/internal/config"
	"gopkg.in/gomail.v2"
)

type MailRequest struct {
	To      string
	Subject string
	Body    string
}

func SendEmail(mail MailRequest) error {
	host := config.AppConfig.EMAIL_HOST
	portStr := config.AppConfig.EMAIL_PORT
	username := config.AppConfig.EMAIL_USERNAME
	password := config.AppConfig.EMAIL_PASSWORD
	fromName := config.AppConfig.EMAIL_FROM_NAME

	port, err := strconv.Atoi(portStr)
	if err != nil {
		return fmt.Errorf("invalid SMTP port: %v", err)
	}

	fmt.Println("Email Config:", host, port, username, password)

	m := gomail.NewMessage()
	m.SetHeader("From", fmt.Sprintf("%s <%s>", fromName, username))
	m.SetHeader("To", mail.To)
	m.SetHeader("Subject", mail.Subject)
	m.SetBody("text/html", mail.Body)

	d := gomail.NewDialer(host, port, username, password)

	if err := d.DialAndSend(m); err != nil {
		return fmt.Errorf("failed to send email: %w", err)
	}

	return nil
}

func GenerateVerificationEmailBody(code string) string {
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Verifikasi Email</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);">
				<h2 style="color: #333333;">Kode Verifikasi Email</h2>
				<p>Halo,</p>
				<p>Gunakan kode verifikasi berikut untuk menyelesaikan proses verifikasi email Anda:</p>
				<h1 style="color: #007BFF; letter-spacing: 4px;">%s</h1>
				<p>Kode ini berlaku selama 10 menit.</p>
				<p>Jika Anda tidak merasa melakukan permintaan ini, abaikan email ini.</p>
				<br>
				<p>Salam,</p>
				<p><strong>RT App</strong></p>
			</div>
		</body>
		</html>
	`, code)
}

func GenerateForgotPasswordOTPEmailBody(otp string) string {
	return fmt.Sprintf(`
		<!DOCTYPE html>
		<html>
		<head>
			<meta charset="UTF-8">
			<title>Reset Password - OTP</title>
		</head>
		<body style="font-family: Arial, sans-serif; background-color: #f9f9f9; padding: 20px;">
			<div style="max-width: 600px; margin: auto; background-color: #ffffff; padding: 20px; border-radius: 8px; box-shadow: 0 0 10px rgba(0, 0, 0, 0.05);">
				<h2 style="color: #333333;">Kode OTP Reset Password</h2>
				<p>Halo,</p>
				<p>Gunakan kode OTP berikut untuk mengatur ulang password akun Anda:</p>
				<h1 style="color: #DC3545; letter-spacing: 4px;">%s</h1>
				<p>Kode ini hanya berlaku selama 10 menit.</p>
				<p>Jika Anda tidak merasa melakukan permintaan ini, abaikan email ini.</p>
				<br>
				<p>Salam,</p>
				<p><strong>RT App</strong></p>
			</div>
		</body>
		</html>
	`, otp)
}

// code := "123456" // misalnya generate pakai random string/OTP generator
// body := GenerateVerificationEmailBody(code)
//
// err := utils.SendEmail(utils.MailRequest{
// 	To:      "user@email.com",
// 	Subject: "Kode Verifikasi Email Anda",
// 	Body:    body,
// })
