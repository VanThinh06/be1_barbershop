package util
import (
	"fmt"
	"math"
)

func MsgForTag(tag string) string {
	switch tag {
	case "Email":
		return "Invalid email"
	case "password":
		return "Password must be more than 6 characters"
	}
	return ""
}

func compute(fns func(float64, float64) float64) float64 {
	return fns(3, 4)
}

func main() {
	hypot := func(x, y float64) float64 {
		return math.Sqrt(x*x + y*y)
	}
	fmt.Println(hypot(5, 12))
	math.Pow(3, 4)

	fmt.Println(compute(hypot))
	fmt.Println(compute(math.Pow))
}
