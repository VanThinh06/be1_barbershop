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
	Barber
	Skinner
	Receptionist
	Manager
	Admin
)

type EnumRoleType string

const (
	Staff         EnumRoleType = "Staff"
	Administrator EnumRoleType = "Administrator"
)
