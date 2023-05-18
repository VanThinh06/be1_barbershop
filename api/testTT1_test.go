package api

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {

	require.Equal(t, TestisPrime(5), true)
	require.Equal(t, factorSum(2), 2)
	fmt.Println(diagonalDifference([][]int32{
		{1, 2, 3},
		{4, 5, 6},
		{9, 8, 9},
	}))

}

func TestisPrime(n int) bool {
	if n > 2 {

		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
			fmt.Println(i)
			if n%i == 0 {
				return false
			}
		}
		return true
	}
	return false
}

func diagonalDifference(arr [][]int32) int32 {
	// Write your code here
	var result int32
	var result1 int32
	var result2 int32

	for i, _ := range arr {
		result1 += arr[i][i]
		result2 += arr[i][len(arr)-i-1]
	}
	result = int32(math.Abs(float64(result1 - result2)))
	return result
}

func factorSum(n int) int {

	for n != solve(n) {
		n = solve(n)
	}
	return n
}

func solve(n int) int {
	sum, k := 0, 2
	for n > 1 {
		for n%k == 0 {
			sum += k // 12
			n /= k   //
		}
		k++
	}
	return sum
}
