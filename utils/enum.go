package utils

type Gender int

const (
	Male Gender = iota + 1
	Female
	Other
)

type Role int

const (
	HairStylist Role = iota + 1
	ServiceProvider
	Manager
)
