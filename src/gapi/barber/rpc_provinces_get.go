package gapi

import (
	"barbershop/src/pb/barber"
	"context"
	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetProvinces(ctx context.Context, _ *barber.GetProvincesRequest) (*barber.GetProvincesResponse, error) {

	res, err := server.Store.GetProvinces(ctx)
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetProvincesResponse{
		Provinces: convertProvinces(res),
	}
	return rsp, nil
}