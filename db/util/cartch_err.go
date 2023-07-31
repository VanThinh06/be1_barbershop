package util

func MsgForTag(tag string) string {
	switch tag {
	case "Email":
		return "Invalid email"
	case "password":
		return "Password must be more than 6 characters"
	}
	return ""
}
