package gapi

import (
	"barbershop/src/shared/token"
	"barbershop/src/shared/utilities"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/metadata"
)

const (
	authorizationHeader = "authorization"
	authorizationBearer = "bearer"
)

func (server *Server) AuthorizeUser(ctx context.Context) (*token.Payload, error) {
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
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %v", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyToken(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return payload, nil
}

func (server *Server) AuthorizeCustomer(ctx context.Context) (*token.CustomerPayload, error) {
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
		return nil, fmt.Errorf("invalid authorization header format")
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return nil, fmt.Errorf("unsupported authorization type: %v", authType)
	}

	accessToken := fields[1]
	payload, err := server.tokenMaker.VerifyTokenCustomer(accessToken)
	if err != nil {
		return nil, fmt.Errorf("invalid token: %v", err)
	}

	return payload, nil
}

func (server *Server) IsAdministrator(ctx context.Context, payload *token.Payload) error {
	if payload.Barber.BarberRoleType != string(utilities.Administrator) {
		return fmt.Errorf("PermissionDenied")
	}
	return nil
}
