package gapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/barber"
	"context"
	"fmt"
	"strings"

	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/metadata"
	"google.golang.org/grpc/status"
)

func (server *Server) ListNearbyBarberShops(ctx context.Context, req *barber.ListNearbyBarberShopsRequest) (*barber.ListNearbyBarberShopsResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		accessToken, err := extractTokenFromHeader(ctx)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}

		_, err = server.tokenMaker.VerifyTokenCustomer(accessToken)
		if err != nil {
			return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
		}
	}

	arg := db.ListNearbyBarberShopsParams{
		CurrentLongitude: req.GetLongitude().GetValue(),
		CurrentLatitude:  req.GetLatitude().GetValue(),
		DistanceInMeters: req.DistanceInMeters,
	}

	res, err := server.Store.ListNearbyBarberShops(ctx, arg)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.ListNearbyBarberShopsResponse{
		BarberShops: ConvertListBarberShopsNearby(res),
	}
	return rsp, nil
}

func extractTokenFromHeader(ctx context.Context) (string, error) {
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

	return accessToken, nil
}
