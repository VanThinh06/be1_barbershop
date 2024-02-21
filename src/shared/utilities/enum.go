package utilities

type EnumGender int32

const (
	Male EnumGender = iota + 1
	Female
	Other
)

type EnumRole int32

const (
	NoRole EnumGender = iota
	Customer
	Barber
	Receptionist
	Manager
	Admin
	SuperAdmin
	OtherStaff
)

type EnumRoleType string

const (
	Staff         EnumRoleType = "Staff"
	Administrator EnumRoleType = "Administrator"
)
