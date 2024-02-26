package customergapi

import (
	db "barbershop/src/db/sqlc"
	"barbershop/src/pb/customer"

	"google.golang.org/protobuf/types/known/timestamppb"
)

func convertCustomer(res db.Customer) *customer.Customer {
	return &customer.Customer{
		Phone:             res.Phone.String,
		Email:             res.Email,
		Gender:            int32(res.GenderID),
		Name:              res.Name,
		Avatar:            res.Avatar.String,
		CustomerId:        res.ID.String(),
		CreatedAt:         timestamppb.New(res.CreateAt),
		PasswordChangedAt: timestamppb.New(res.PasswordChangedAt),
		IsSocialAuth:      res.IsSocialAuth.Bool,
	}
}

