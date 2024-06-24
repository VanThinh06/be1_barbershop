package utils

import (
	"strings"

	"github.com/google/uuid"
)

func GenerateNickName(fullName string) string {

	newUUID := uuid.New()
	uuidPrefix := newUUID.String()[:8]
	parts := strings.Split(fullName, " ")
	nickName := parts[len(parts)-1]
	combinedNickName := strings.TrimSpace(nickName) + uuidPrefix
	return combinedNickName
}
