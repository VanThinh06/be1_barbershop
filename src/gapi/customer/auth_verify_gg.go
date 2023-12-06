package customergapi

import (
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	"google.golang.org/grpc/metadata"
)

func (server *Server) authVerifyJWTGG(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return "", fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(authHeader) < 2 {
		return "", fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return "", fmt.Errorf("unsupported authorization type: %v", authType)
	}

	accessToken := fields[1]
	url := "https://oauth2.googleapis.com/tokeninfo"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", fmt.Errorf("Lỗi khi cấu hình: %v", err)
	}
	query := req.URL.Query()
	query.Add("access_token", accessToken)
	req.URL.RawQuery = query.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return "", fmt.Errorf("Lỗi khi cấu hình: %v", err)
	}

	var response struct {
		Email string `json:"email"`
	}

	// Đọc nội dung của resp.Body
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("Lỗi khi cấu hình: %v", err)

	}

	// Phân tích phản hồi JSON vào biến response
	err = json.Unmarshal(body, &response)
	if err != nil {
		return "", fmt.Errorf("Lỗi khi cấu hình: %v", err)

	}
	return response.Email, nil

	// config, err := google.From(ctx)
	// if err != nil {
	// 	return "", fmt.Errorf("Lỗi khi cấu hình: %v", err)
	// }

	// token := oauth2.Token{
	// 	AccessToken: accessToken,
	// 	TokenType:   "Bearer",
	// }

	// client := config.Client(ctx, &token)

	// _, err = client.Get("https://www.googleapis.com/oauth2/v3/userinfo")
	// if err != nil {
	// 	return "", fmt.Errorf("Lỗi khi xác thực: %v", err)
	// }

}
