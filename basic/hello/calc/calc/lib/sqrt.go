package lib

import "math"

/**
求平方根
 */
func Sqrt(val int) int {
	v := math.Sqrt(float64(val))
	return int(v)
}
