package utilities

type EnumGender int32

const (
	Male EnumGender = iota + 1
	Female
	Other
)

type EnumRole int32

const (
	Customer EnumRole = iota + 1
	Barber
	Receptionist
	Manager
	Admin
	SuperAdmin
	OtherStaff
)

type EnumRoleType string

const (
	User          EnumRoleType = "User"
	Staff         EnumRoleType = "Staff"
	Administrator EnumRoleType = "Administrator"
)
