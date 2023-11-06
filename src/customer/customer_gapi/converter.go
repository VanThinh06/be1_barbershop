package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCustomer(customer db.Customer) *pb.Customer {
	return &pb.Customer{
		Phone:             customer.Phone,
		Email:             customer.Email,
		Gender:            customer.Gender,
		Name:              customer.Name,
		Avatar:            customer.Avatar.String,
		CustomerId:        customer.CustomerID.String(),
		CreatedAt:         timestamppb.New(customer.CreateAt),
		UpdateAt:          timestamppb.New(customer.UpdateAt.Time),
		PasswordChangedAt: timestamppb.New(customer.PasswordChangedAt),
	}
}
