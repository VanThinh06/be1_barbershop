package utils

type gender int
const (
	Male gender = iota + 1
	Female
	Other
)

// role 
type role int
const (
	HairStylist role = iota + 1
	ServiceProvider
	Manager
)

type statusAppoinment int
const (
	Pending statusAppoinment = iota + 1
  	Confirmed
  	Canceled
  	Completed
)