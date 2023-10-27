package val

import (
	"errors"
	"fmt"
	"regexp"
)

// ValidateString kiểm tra độ dài của một chuỗi.
func ValidateString(value string, minLength, maxLength int) error {
	n := len(value)
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

	emailRegex := `^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`
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

// ValidatePhoneNumber kiểm tra tính hợp lệ của số điện thoại.
func ValidatePhoneNumber(phoneNumber string) error {
	normalizedNumber := regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")

	e164Regex := regexp.MustCompile(`^\+84\d{9,10}$`)
	if e164Regex.MatchString(normalizedNumber) {
		return nil
	}

	vietnamRegex := regexp.MustCompile(`^0\d{9}$`)
	if vietnamRegex.MatchString(normalizedNumber) {
		return nil
	}

	return errors.New("invalid phone number format")
}

// ValidateFullName kiểm tra tính hợp lệ của họ và tên.
func ValidateFullName(fullName string) error {
	if err := ValidateString(fullName, 3, 50); err != nil {
		return err
	}

	fullNameRegex := `^[a-zA-Z\s]+$`
	match, _ := regexp.MatchString(fullNameRegex, fullName)
	if !match {
		return errors.New("invalid full name format")
	}

	return nil
}

// ValidateNickname kiểm tra tính hợp lệ của nickname.
func ValidateNickname(nickname string) error {
	if err := ValidateString(nickname, 1, 20); err != nil {
		return err
	}

	nicknameRegex := "^[a-zA-Z0-9" + regexp.QuoteMeta(".#@-_") + "]+$"
	match, _ := regexp.MatchString(nicknameRegex, nickname)
	if !match {
		return errors.New("invalid nickname format")
	}

	return nil
}
