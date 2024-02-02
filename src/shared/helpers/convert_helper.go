package helpers

import (
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

func ConvertFloat64ToFloat32Pointer(number float64) *float32 {
	convertedNumber := float32(number)
	return &convertedNumber
}

func ContainsValueUUID(array []uuid.UUID, valueToCheck uuid.UUID) bool {
	for _, element := range array {
		if element == valueToCheck {
			return true
		}
	}
	return false
}
