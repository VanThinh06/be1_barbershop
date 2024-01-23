package utils

import (
	"errors"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetRemainingString(fullString, subString string) string {
	index := strings.Index(fullString, subString)
	if index != -1 {
		return fullString[index+len(subString):]
	}
	return ""
}

func GetValueAtIndex(arr []string, index int) (string, error) {
	if index >= 0 && index < len(arr) {
		return arr[index], nil
	}
	return "", errors.New("Chỉ số nằm ngoài phạm vi của mảng")
}

func ConvertStringToInt(str string) (int, error) {
	value, err := strconv.Atoi(str)
	if err != nil {
		return 0, err
	}
	return value, nil
}

func ConvertStringListToUUIDList(stringList []string) ([]uuid.UUID, error) {
	var uuidList []uuid.UUID
	for _, str := range stringList {
		uuidVal, err := uuid.Parse(str)
		if err != nil {
			return nil, err
		}
		uuidList = append(uuidList, uuidVal)
	}
	return uuidList, nil
}
