package api

import (
	"fmt"
	"testing"
)

func Test(t *testing.T)  {
	fmt.Println(TestisPrime(133))
	TestisPrime(133)
}

func TestisPrime(n int) bool {
	if n > 2{
		iCheck := false
		for i := 2; i < n/2; i++ {
		if n % i != 0 {
			iCheck = true
		}
	return iCheck
	}
	}
	return false
	}
