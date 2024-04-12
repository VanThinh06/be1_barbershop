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
	EmployeeManager
	ServiceManager
	Admin
)

type EnumRoleType string

const (
	Staff         EnumRoleType = "Staff"
	Management    EnumRoleType = "Management"
	Administrator EnumRoleType = "Administrator"
)

var MapRoleType = map[int32]string{
	int32(NoRole):          string(NoRole),
	int32(Barber):          string(Staff),
	int32(EmployeeManager): string(Management),
	int32(ServiceManager):  string(Management),
	int32(Admin):           string(Administrator),
}

type EnumPermission int32

const (
	NoPermisson EnumGender = iota
	ViewEmployeeInformation
	ManageEmployee
	ViewServiceList
	ManageService
	ViewAppointmentList
	ManageAppointment
)