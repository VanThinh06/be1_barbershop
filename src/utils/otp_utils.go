package utils

import (
	"crypto/rand"
	"errors"
	"fmt"
	"math/big"

	"gopkg.in/gomail.v2"
)

func GenerateOTP() string {
	const otpLength = 6
	const digits = "0123456789"
	otp := make([]byte, otpLength)
	for i := range otp {
		num, _ := rand.Int(rand.Reader, big.NewInt(int64(len(digits))))
		otp[i] = digits[num.Int64()]
	}
	return string(otp)
}

func SendOTP(email, otp string) error {
	config, err := InitConfig(".")
	if err != nil {
		return errors.New("token is invalid")
	}

	m := gomail.NewMessage()
	m.SetHeader("From", config.AccountEmail)
	m.SetHeader("To", email)
	m.SetHeader("Subject", "Your OTP Code")
	m.SetBody("text/plain", fmt.Sprintf("Your OTP code is %s", otp))

	d := gomail.NewDialer("smtp.gmail.com", 587, config.AccountEmail, config.PasswordEmail)

	if err := d.DialAndSend(m); err != nil {
		return err
	}
	return nil
}
