package api

import (
	"fmt"
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func Test(t *testing.T) {

	plusMinus([]int32{
		1, 1, 0, -1, -1})
	require.Equal(t, factorSum(2), 2)
	// fmt.Println(diagonalDifference([][]int32{
	// 	{1, 2, 3},
	// 	{4, 5, 6},
	// 	{9, 8, 9},
	// }))

}

func plusMinus(arr []int32) {
	// Write your code here
	nagative := 0.0
	positive := 0.0
	zeros := 0.0
	for _, value := range arr {

		if value == 0 {
			zeros += 1
		} else if value > 0 {
			positive += 1
		} else if value < 0 {
			nagative += 1
		}
	}

	fmt.Printf("%f", float64(positive)/float64(len(arr)))
	fmt.Println()
	fmt.Printf("%f", float64(nagative)/float64(len(arr)))
	fmt.Println()

	fmt.Printf("%f", float64(zeros)/float64(len(arr)))

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
