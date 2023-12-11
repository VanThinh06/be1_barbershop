package utils

type gender int
const (
	Male gender = iota + 1
	Female
	Other
)
func GetGender() gender {
	return Male
}

// role 
type role int
const (
	HairStylist role = iota + 1
	ServiceProvider
	Manager
)
func GetRole() role {
	return HairStylist
}