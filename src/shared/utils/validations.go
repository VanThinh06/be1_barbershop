package utils

import (
	"errors"
	"fmt"
	"regexp"
	"strings"
	"unicode"
)

var (
	nicknameRegex = "^[a-zA-Z0-9" + regexp.QuoteMeta(".#@-_") + "]+$"
	emailRegex    = `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
	e164Regex     = regexp.MustCompile(`^\+84\d{9,10}$`)
	vietnamRegex  = regexp.MustCompile(`^0\d{9}$`)
	numberRegex   = regexp.MustCompile(`^\d+$`)
)

const (
	minNumberLength = 1
	maxNumberLength = 10
)

func ValidateNumber(number string) error {
	if err := ValidateString(number, minNumberLength, maxNumberLength); err != nil {
		return err
	}

	match := numberRegex.MatchString(number)
	if !match {
		return errors.New("invalid number format")
	}

	return nil
}

// ValidateString kiểm tra độ dài của một chuỗi.
func ValidateString(value string, minLength, maxLength int) error {
	n := len(strings.TrimSpace(value))
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d characters", minLength, maxLength)
	}
	return nil
}

// ValidateEmail kiểm tra tính hợp lệ của địa chỉ email.
func ValidateEmail(email string) error {
	if err := ValidateString(email, 1, 50); err != nil {
		return err
	}

	match, _ := regexp.MatchString(emailRegex, email)
	if !match {
		return errors.New("invalid email format")
	}

	return nil
}

// ValidatePassword kiểm tra độ dài của mật khẩu.
func ValidatePassword(password string) error {
	if err := ValidateString(password, 6, 50); err != nil {
		return err
	}

	return nil
}

func ValidateId(id string) error {
	if err := ValidateString(id, 36, 36); err != nil {
		return err
	}

	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	normalizedNumber := regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")
	if e164Regex.MatchString(normalizedNumber) {
		return nil
	}

	if vietnamRegex.MatchString(normalizedNumber) {
		return nil
	}

	return errors.New("invalid phone number format")
}

func ValidateFullName(fullName string) error {
	if err := ValidateString(fullName, 3, 50); err != nil {
		return err
	}

	for _, char := range fullName {
		if !isVietnameseLetter(char) && !unicode.IsSpace(char) {
			return errors.New("invalid full name format")
		}
	}

	return nil
}

func ValidateNickname(nickname string) error {
	if err := ValidateString(nickname, 1, 20); err != nil {
		return err
	}

	match, _ := regexp.MatchString(nicknameRegex, nickname)
	if !match {
		return errors.New("invalid nickname format")
	}

	return nil
}

func isVietnameseLetter(char rune) bool {
	vietnameseCharset := []*unicode.RangeTable{
		unicode.Letter,
	}

	for _, charset := range vietnameseCharset {
		if unicode.In(char, charset) {
			return true
		}
	}

	return false
}
