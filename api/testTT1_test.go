 package api

// <<<<<<< ft/docker
// import (
// 	"fmt"
// 	"math"
// 	"testing"
// )
// =======
// // import (
// // 	"fmt"
// // 	"math"
// // 	"testing"

// // 	"github.com/stretchr/testify/require"
// // )
// >>>>>>> main

// // func Test(t *testing.T) {
// // <<<<<<< master

// <<<<<<< ft/docker
// 	// plusMinus([]int32{
// 	// 	1, 1, 0, -1, -1})
// 	// require.Equal(t, factorSum(2), 2)
// 	// // fmt.Println(diagonalDifference([][]int32{
// 	// // 	{1, 2, 3},
// 	// // 	{4, 5, 6},
// 	// // 	{9, 8, 9},
// 	// // }))
// 	// arr := []int32{1, 3, 5,7 ,9 }
// 	miniMaxSum([]int32{256741038, 623958417, 467905213, 714532089, 938071625})

// }

// // CanFastAttack can be executed only when the knight is sleeping.
// func CanFastAttack(knightIsAwake bool) bool {
//  if knightIsAwake {
//  return false    
//  }
// return true
// }

// // CanSpy can be executed if at least one of the characters is awake.
// func CanSpy(knightIsAwake, archerIsAwake, prisonerIsAwake bool) bool {
// 	 if knightIsAwake && archerIsAwake && prisonerIsAwake == false{
//          return true
//      }
// return false
// }

// // CanSignalPrisoner can be executed if the prisoner is awake and the archer is sleeping.
// func CanSignalPrisoner(archerIsAwake, prisonerIsAwake bool) bool {
// 	 if archerIsAwake && prisonerIsAwake == false{
//      return true    
//      }
// return false
// }

// // CanFreePrisoner can be executed if the prisoner is awake and the other 2 characters are asleep
// // or if Annalyn's pet dog is with her and the archer is sleeping.
// func CanFreePrisoner(knightIsAwake, archerIsAwake, prisonerIsAwake, petDogIsPresent bool) bool {
// 	if petDogIsPresent{
//     if(archerIsAwake == false){
//     return true    
//     }else{
//     if(prisonerIsAwake && knightIsAwake == false && archerIsAwake == false){
//         return true
//         }
//         }
//     return false
//     }
// 	return false
// }


// func miniMaxSum(arr []int32) {
// 	var numMax int32 = arr[0]
// 	var numMin int32 = arr[0]
// 	var sumMax int
// 	var sumMin int
// 	var arrMin []int32
// 	var arrMax []int32
// 	var equalNumMin int
// 	var equalNumMax int

// 	for _, v := range arr {
// 		if v < numMin {
// 			numMin = v
// 		}

// 		if v > numMax {
// 			numMax = v
// 		}
// 	}

// 	for _, v := range arr {
// 		if v < numMax {
// 			arrMin = append(arrMin, v)
// 		}
// 		if v > numMin {
// 			arrMax = append(arrMax, v)
// 		}
// 		if v == numMin {
// 			equalNumMin++
// 		}
// 		if v == numMax {
// 			equalNumMax++
// 		}
// 	}

// 	for _, v := range arrMax {
// 		sumMax += int(v)
// 	}

// 	for _, v := range arrMin {
// 		sumMin += int(v)
// 	}

// 	if equalNumMax > 1 {
// 		sumMax += int(numMax) * (int(equalNumMax) - 1)
// 	}
// 	if equalNumMin > 1 {
// 		sumMin += int(numMin) * (int(equalNumMax) - 1)
// 	}

// 	fmt.Printf("%d %d", sumMin, sumMax)
// }

// func plusMinus(arr []int32) {
// 	// Write your code here
// 	nagative := 0.0
// 	positive := 0.0
// 	zeros := 0.0
// 	for _, value := range arr {

// 		if value == 0 {
// 			zeros += 1
// 		} else if value > 0 {
// 			positive += 1
// 		} else if value < 0 {
// 			nagative += 1
// 		}
// 	}

// 	fmt.Printf("%f", float64(positive)/float64(len(arr)))
// 	fmt.Println()
// 	fmt.Printf("%f", float64(nagative)/float64(len(arr)))
// 	fmt.Println()

// 	fmt.Printf("%f", float64(zeros)/float64(len(arr)))

// }
// =======
// // 	require.Equal(t, TestisPrime(5), true)
// // 	require.Equal(t, factorSum(2), 2)
// // 	fmt.Println(diagonalDifference([][]int32{
// // 		{1, 2, 3},
// // 		{4, 5, 6},
// // 		{9, 8, 9},
// // 	}))

// // =======
// // 	require.Equal(t, TestisPrime(120), false)
// // >>>>>>> main
// // }

// // func TestisPrime(n int) bool {
// // 	if n > 2 {
// // <<<<<<< master

// // 		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
// // 			fmt.Println(i)
// // =======
// // 		iCheck := true
// // 		fmt.Printf("âj %d", int(math.Sqrt(float64(n))))
// // 		for i := 2; i <= int(math.Sqrt(float64(n))); i++ {
// // >>>>>>> main
// // 			if n%i == 0 {
// // 				return false
// // 			}
// // 		}
// // <<<<<<< master
// // 		return true
// // 	}
// // 	return false
// // }
// >>>>>>> main

// // func diagonalDifference(arr [][]int32) int32 {
// // 	// Write your code here
// // 	var result int32
// // 	var result1 int32
// // 	var result2 int32

// // 	for i, _ := range arr {
// // 		result1 += arr[i][i]
// // 		result2 += arr[i][len(arr)-i-1]
// // 	}
// // 	result = int32(math.Abs(float64(result1 - result2)))
// // 	return result
// // }

// // func factorSum(n int) int {

// // 	for n != solve(n) {
// // 		n = solve(n)
// // 	}
// // 	return n
// // }

// // func solve(n int) int {
// // 	sum, k := 0, 2
// // 	for n > 1 {
// // 		for n%k == 0 {
// // 			sum += k // 12
// // 			n /= k   //
// // 		}
// // 		k++
// // 	}
// // 	return sum
// // }
// // =======
// // 		return iCheck
// // 	}
// // 	return false
// // }
// // >>>>>>> main
