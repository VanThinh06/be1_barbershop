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

func (server *Server) FindBarberShopsNearby(ctx context.Context, req *barber.FindBarberShopsNearbyRequest) (*barber.FindBarberShopsNearbyResponse, error) {

	_, err := server.AuthorizeUser(ctx)
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

	arg := db.FindBarberShopsNearbyLocationsParams{
		CurrentLongitude: req.GetLatitude().GetValue(),
		CurrentLatitude:  req.GetLatitude().GetValue(),
		DistanceInMeters: req.DistanceInMeters,
	}

	res, err := server.Store.FindBarberShopsNearbyLocations(ctx, arg)
	if err != nil {
		return nil, status.Errorf(codes.Internal, "failed to create barber shop ", err)
	}

	rsp := &barber.FindBarberShopsNearbyResponse{
		BarberShops: ConvertListBarberShopsNearby(res),
	}
	return rsp, nil
}
func extractTokenFromHeader(ctx context.Context) (string, error) {
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return "", fmt.Errorf("missing metadata") // Return an error if metadata is missing in the context
	}

	values := md.Get(authorizationHeader)
	if len(values) == 0 {
		return "", fmt.Errorf("missing authorization header") // Return an error if the "Authorization" header is missing
	}

	authHeader := values[0]
	fields := strings.Fields(authHeader)
	if len(authHeader) < 2 {
		return "", fmt.Errorf("invalid authorization header format") // Return an error if the authorization header format is invalid
	}

	authType := strings.ToLower(fields[0])
	if authType != authorizationBearer {
		return "", fmt.Errorf("unsupported authorization type: %v", authType) // Return an error if the authorization type is unsupported
	}

	// verifyTokenCustomer

	accessToken := fields[1]

	return accessToken, nil
}
