package utilities

type EnumGender int

const (
	Male EnumGender = iota + 1
	Female
	Other
)

type EnumRole int

const (
	Customer EnumRole = iota + 1
	Barber
	Receptionist
	Admin
	OtherStaff
)

type EnumRoleType string

const (
	User          EnumRoleType = "User"
	Staff         EnumRoleType = "Staff"
	Administrator EnumRoleType = "Administrator"
)
