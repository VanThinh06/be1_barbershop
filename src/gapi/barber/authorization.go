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

func (server *Server) authorizeBarber(ctx context.Context) (*token.BarberPayload, error) {
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

func (server *Server) authorizeCustomer(ctx context.Context) (*token.CustomerPayload, error) {
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

func (server *Server) isAdministrator(payload *token.BarberPayload) bool {
	if payload.Barber.BarberRoleType != string(utilities.Administrator) {
		return false
	}
	return true
}

func (server *Server) authorizeBarberAndCustomer(ctx context.Context) error {
	_, err := server.authorizeBarber(ctx)
	if err != nil {
		_, err := server.authorizeCustomer(ctx)
		if err != nil {
			return fmt.Errorf("unauthenticated")
		}
	}
	return nil
}

func (server *Server) checkCreateAccountBarberPermission(ctx context.Context, payload *token.BarberPayload, roleId int32) error {
	var err error = fmt.Errorf("no permission to access")
	if payload.Barber.BarberRoleType != string(utilities.Administrator) || payload.Barber.BarberRole != int32(utilities.Manager) {
		return err
	}

	if payload.Barber.BarberRole == int32(utilities.Manager) {
		if roleId != int32(utilities.Barber) && roleId != int32(utilities.OtherStaff) {
			return err
		}
	}
	return nil
}

