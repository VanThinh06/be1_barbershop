package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/lib/pq"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

func (server *Server) GetDistricts(ctx context.Context, req *barber.GetDistrictsRequest) (*barber.GetDistrictsResponse, error) {

	res, err := server.Store.GetDistricts(ctx, int16(req.ProvinceId))
	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok {
			switch pqErr.Code.Name() {
			}
		}
		return nil, status.Errorf(codes.Internal, "internal")
	}

	rsp := &barber.GetDistrictsResponse{
		Districts: convertDistricts(res),
	}
	return rsp, nil
}