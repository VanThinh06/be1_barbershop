package customergapi

import (
	"barbershop/src/shared/token"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) authorizeUser(ctx context.Context) (*token.CustomerPayload, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return nil, fmt.Errorf("missing metadata") // Return an error if metadata is missing in the context
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return nil, fmt.Errorf("missing authorization header") // Return an error if the "Authorization" header is missing
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(authHeader) < 2 {
		return nil, fmt.Errorf("invalid authorization header format") // Return an error if the authorization header format is invalid
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %v", authType) // Return an error if the authorization type is unsupported
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err) // Return an error if the token is invalid
	}

	return payload, nil
}
