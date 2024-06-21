package utils

type GenderEnum int32
const (
	NoGender GenderEnum = iota
	Male
	Female
	Other
)

type RoleEnum int32
const (
	NoRole GenderEnum = iota
	Barber
	Skinner
	EmployeeManager
	ServiceManager
	Admin
)

type RoleTypeEnum string
const (
	Staff         RoleTypeEnum = "Staff"
	Management    RoleTypeEnum = "Management"
	Administrator RoleTypeEnum = "Administrator"
)

type PermissionEnum int32
const (
	NoPermission PermissionEnum = iota
	ViewEmployeeInformation
	ManageEmployee
	ViewServiceList
	ManageService
	ViewAppointmentList
	ManageAppointment
)

var COMBO_SERVICE_ID = 2512
var COMBO_SERVICE_NAME = "c2o5m1b2o"