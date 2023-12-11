package utils

import (
	"errors"
	"strconv"
	"strings"
)

func GetRemainingString(fullString, subString string) string {
	index := strings.Index(fullString, subString)

	if index != -1 {
		return fullString[index+len(subString):]
	}

	return ""
}

func GetValueAtIndex(arr []string, index int) (string, error) {
	// Kiểm tra xem chỉ số có nằm trong phạm vi hợp lý hay không
	if index >= 0 && index < len(arr) {
		return arr[index], nil
	}

	// Trả về lỗi nếu chỉ số nằm ngoài phạm vi
	return "", errors.New("Chỉ số nằm ngoài phạm vi của mảng")
}

func ConvertStringToInt(str string) (int, error) {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return value, nil
}
