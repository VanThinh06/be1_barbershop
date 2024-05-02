package gapi

import (
	"barbershop/src/pb/barber"
	"context"

	"github.com/google/uuid"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (server *Server) GetBarberShopServices(ctx context.Context, req *barber.GetBarberShopServicesRequest) (*barber.GetBarberShopServicesResponse, error) {

	_, err := server.authorizeBarber(ctx)
	if err != nil {
		return nil, status.Errorf(codes.Unauthenticated, "unauthenticated")
	}

	serviceId, err := uuid.Parse(req.GetId())
	if err != nil {
		return nil, status.Errorf(codes.InvalidArgument, "barbershops don't exist")
	}

	res, err := server.Store.GetBarberShopService(ctx, serviceId)
	if err != nil {
		return nil, status.Error(codes.Internal, "internal")
	}

	rsp := &barber.GetBarberShopServicesResponse{
		BarberShopService: &barber.BarberShopServices{
			Id:                res.ID.String(),
			CategoryId:        int32(res.CategoryID),
			CategoryName:      res.CategoryName.String,
			BarberShopId:      res.BarberShopID.String(),
			GenderId:          int32(res.GenderID),
			Name:              res.Name,
			Timer:             int32(res.Timer),
			Price:             res.Price,
			Description:       res.Description.String,
			ImageUrl:          res.ImageUrl.String,
			ComboServices:     res.ComboServices,
			DiscountPrice:     &res.DiscountPrice.Float32,
			DiscountStartTime: timestamppb.New(res.DiscountStartTime.Time),
			DiscountEndTime:   timestamppb.New(res.DiscountEndTime.Time),
			IsActive:          res.IsActive,
		},
	}
	return rsp, nil
}
