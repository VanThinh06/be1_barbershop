package utils

import "regexp"

func ValidatePhoneNumber(phoneNumber string) bool {
	// Loại bỏ các ký tự không phải chữ số
	normalizedNumber := regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")

	// Kiểm tra số điện thoại theo chuẩn E.164
	e164Regex := regexp.MustCompile(`^\+84\d{9,10}$`)
	if e164Regex.MatchString(normalizedNumber) {
		return true
	}

	// Kiểm tra số điện thoại theo định dạng Việt Nam
	vietnamRegex := regexp.MustCompile(`^0\d{9}$`)
	return vietnamRegex.MatchString(normalizedNumber)
}
