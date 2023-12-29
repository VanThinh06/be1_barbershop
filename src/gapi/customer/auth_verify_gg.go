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

func (server *Server) authVerifyJWTGG(ctx context.Context) (*SocialEmail, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata")
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header")
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(authHeader) < 2 {
		return nil, fmt.Errorf("invalid authorization header")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type")
	}

	accessToken := fields[1]
	url := "https://www.googleapis.com/oauth2/v3/userinfo"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration")
	}
	query := req.URL.Query()
	query.Add("access_token", accessToken)
	req.URL.RawQuery = query.Encode()
	client := http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("invalid configuration")
	}
	

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, fmt.Errorf("invalid read response")

	}

	credentialEmail, err := UnmarshalSocialEmail(body)
	if err != nil {
		return nil, fmt.Errorf("invalid parses the response")
	}
	return &credentialEmail, nil
}

func UnmarshalSocialEmail(data []byte) (SocialEmail, error) {
	var r SocialEmail
	err := json.Unmarshal(data, &r)
	return r, err
}

func (r *SocialEmail) Marshal() ([]byte, error) {
	return json.Marshal(r)
}

type SocialEmail struct {
	Sub           string `json:"sub"`
	Name          string `json:"name"`
	GivenName     string `json:"given_name"`
	Picture       string `json:"picture"`
	Email         string `json:"email"`
	EmailVerified bool   `json:"email_verified"`
	Locale        string `json:"locale"`
}
