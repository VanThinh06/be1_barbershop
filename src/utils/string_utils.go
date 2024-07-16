package utils

import (
	"strings"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5/pgtype"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func GenerateNickName(fullName string) string {
	newUUID := uuid.New()
	uuidPrefix := newUUID.String()[:8]
	parts := strings.Split(fullName, " ")
	nickName := parts[len(parts)-1]
	combinedNickName := strings.TrimSpace(nickName) + uuidPrefix
	return combinedNickName
}

func ConvertToTimestampProto(timestamp pgtype.Timestamp) *timestamppb.Timestamp {
	if !timestamp.Valid {
		return nil
	}
	return timestamppb.New(timestamp.Time)
}

func ExtractPublicIDFromURL(url string) string {
	publicID := url[strings.LastIndex(url, "/")+1 : strings.LastIndex(url, ".")]
	return publicID
}
