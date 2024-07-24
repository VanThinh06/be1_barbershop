package helpers

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
	vietnamRegex = regexp.MustCompile(`^(0[35789])[0-9]{8}$`)
	numberRegex  = regexp.MustCompile(`^\d+$`)
)

func ValidateString(value string, minLength, maxLength int) error {
	n := len(strings.TrimSpace(value))
	if n < minLength || n > maxLength {
		return fmt.Errorf("length must be between %d and %d characters", minLength, maxLength)
	}
	return nil
}

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

func ValidatePassword(password string) error {
	if err := ValidateString(password, 6, 50); err != nil {
		return err
	}

	return nil
}

func ValidatePhoneNumber(phoneNumber string) error {
	normalizedNumber := regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")

	if vietnamRegex.MatchString(normalizedNumber) {
		return nil
	}

	return errors.New("invalid phone number format")
}

func ValidateFullName(fullName string) error {
	if err := ValidateString(fullName, 1, 50); err != nil {
		return err
	}

	for _, char := range fullName {
		if !isVietnameseLetter(char) && !unicode.IsSpace(char) {
			return errors.New("invalid full name format")
		}
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
