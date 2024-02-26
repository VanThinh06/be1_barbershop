package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetWards(ctx context.Context, req *barber.GetWardsRequest) (*barber.GetWardsResponse, error) {

	res, err := server.Store.GetWards(ctx, int16(req.DistrictId))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetWardsResponse{
		Wards: convertWards(res),
	}
	return rsp, nil
}