package api

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {
	require.Equal(t, TestisPrime(120), false)
}

func TestisPrime(n int) bool {
	if n > 2 {
		iCheck := true
		fmt.Printf("âj %d", int(math.Sqrt(float64(n))))
		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			if n%i == 0 {
				return false
			}
		}
		return iCheck
	}
	return false
}
