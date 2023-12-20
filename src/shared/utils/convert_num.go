package utils

func ConvertFloat64ToFloat32Pointer(number float64) *float32 {
    convertedNumber := float32(number)
    return &convertedNumber
}